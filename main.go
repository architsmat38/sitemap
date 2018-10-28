package main

import (
	"flag"
	"os"
	"sync"
	"time"

	"github.com/architsmat38/sitemap/crawler"
	"github.com/architsmat38/sitemap/logger"
	"github.com/architsmat38/sitemap/utils"
)

/**
 * Validates the website url
 */
func validateWebsiteURL(websiteURL string) (bool, string) {
	if len(websiteURL) == 0 {
		return false, "Usage: ./sitemap-generator -u http://example.com"
	}

	if !utils.IsValidURL(websiteURL) {
		return false, "Please provide a valid website url"
	}

	return true, ""
}

/**
 * Generate sitemap links of website
 */
func generateSitemapLinks(websiteURL string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Fetch host for website URL
	host, err := utils.GetHost(websiteURL)
	if err != nil {
		logger.Error(err)
		return
	}

	crawlQueue := make(chan string, 512)
	crawlerObj := crawler.NewCrawler(20, crawlQueue, websiteURL)
	worker := crawlerObj.GetWorker()

	crawlerObj.FilterOutAndUpdateCrawledURLs([]string{websiteURL})
	crawlQueue <- websiteURL

	var totalCrawled int
	for {
		select {
		case url := <-crawlQueue:
			logger.Debug("Enqueue url: ", url)
			task := &crawler.Request{WebsiteHost: host, LinkURL: url, CrawlObj: crawlerObj, CrawlQueue: crawlQueue}
			worker.Exec(task)
			totalCrawled++

			// Adding delay to avoid getting blocked
			if totalCrawled%100 == 0 {
				time.Sleep(time.Millisecond * 250)
			}

		case <-time.Tick(10 * time.Second):
			if worker.GetQueueSize()+len(crawlQueue) == 0 {
				logger.Debug("Completed processing sitemap")
				// Print sitemap
				crawlerObj.SitemapObj.Print()

				// Close crawler
				crawlerObj.Close()
				return
			}
		}
	}
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
	go generateSitemapLinks(websiteURL, &wg)

	wg.Wait()
}
