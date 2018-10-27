package main

import (
	"flag"
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/architsmat38/sitemap/crawler"
	"github.com/architsmat38/sitemap/logger"
)

/**
 * Validates the website url
 */
func validateWebsiteURL(websiteURL string) (bool, string) {
	if len(websiteURL) == 0 {
		return false, "Usage: ./sitemap-generator -u example.com"
	}

	_, err := url.ParseRequestURI(websiteURL)
	if err != nil {
		return false, "Please provide a valid website url"
	}

	return true, ""
}

/**
 * Generate sitemap links of website
 */
func generateSitemapLinks(websiteURL string, wg sync.WaitGroup) {
	crawlQueue := make(chan string, 512)
	crawlerObj := crawler.NewCrawler(50, crawlQueue)
	worker := crawlerObj.GetWorker()

	crawlQueue <- websiteURL

	for true {
		select {
		case url := <-crawlQueue:
			task := &crawler.Request{WebsiteURL: url}
			worker.Exec(task)
		case <-time.Tick(10 * time.Second):
			if worker.GetQueueSize()+len(crawlQueue) == 0 {
				break
			}
		}
	}

	// Print sitemap
	crawlerObj.SitemapObj.Print()

	// Close crawler
	crawlerObj.Close()
	wg.Done()
}

/**
 * Main program
 */
func main() {
	var websiteURL string
	flag.StringVar(&websiteURL, "u", "", "Specify the website which needs to be crawled")
	flag.Parse()

	// Validate the argument
	isValidWebsite, errorVal := validateWebsiteURL(websiteURL)
	if !isValidWebsite {
		logger.Print(errorVal)
		os.Exit(1)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	/**
	 * Generate sitemap links
	 * Doing it in this way, as it will be easier to support sitemap generation
	 * of multiple websites simultaneously (TODO)
	 */
	go generateSitemapLinks(websiteURL, wg)

	wg.Wait()
}
