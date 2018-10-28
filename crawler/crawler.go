package crawler

import (
	"net/http"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/architsmat38/sitemap/logger"
	"github.com/architsmat38/sitemap/parser"
	"github.com/architsmat38/sitemap/pool"
	"github.com/architsmat38/sitemap/sitemap"
	"github.com/architsmat38/sitemap/utils"
)

/**
 * Structure of request to be processed by crawler
 */
type Request struct {
	WebsiteHost string
	LinkURL     string
	CrawlObj    *CrawlerObj
	CrawlQueue  chan string
}

/**
 * Structure of crawler object
 */
type CrawlerObj struct {
	poolObj      *pool.Pool
	poolSize     int
	SitemapObj   *sitemap.SitemapObj
	crawledLinks map[string]bool
	crawlChan    chan string
}

/**
 * Creates crawler with worker pool of specified size
 */
func NewCrawler(size int, queue chan string, linkURL string) *CrawlerObj {
	s := sitemap.NewSitemapManager(linkURL)
	c := CrawlerObj{
		poolSize:     size,
		SitemapObj:   s,
		crawledLinks: make(map[string]bool),
		crawlChan:    queue,
	}
	c.init()
	return &c
}

/**
 * Initializes worker pool for the crawler
 */
func (c *CrawlerObj) init() {
	c.poolObj = pool.NewPool(c.poolSize)
}

/**
 * Get worker for crawler
 */
func (c *CrawlerObj) GetWorker() *pool.Pool {
	return c.poolObj
}

/**
 * Close the crawler
 */
func (c *CrawlerObj) Close() {
	c.poolObj.Close()
}

/**
 * Filter out all URLs which have been already crawled and also update crawled URLs
 */
func (c *CrawlerObj) FilterOutAndUpdateCrawledURLs(allURLs []string) []string {
	var filteredURLs []string
	var mu sync.RWMutex
	mu.Lock()
	for _, url := range allURLs {
		if _, ok := c.crawledLinks[url]; !ok {
			filteredURLs = append(filteredURLs, url)
			c.crawledLinks[url] = true
		}
	}
	mu.Unlock()

	return filteredURLs
}

/**
 * Process the crawler request
 */
func (r Request) Execute() {
	// Make HTTP request
	crawlableURL := utils.GetCrawlableURL(r.LinkURL)

	// Create and modify HTTP request before sending
	request, err := http.NewRequest("GET", crawlableURL, nil)
	if err != nil {
		logger.Error(err)
		return
	}
	request.Header.Set("User-Agent", utils.GetProxyUserAgent())
	request.Header.Add("X-Forwarded-For", utils.GetRandomIP())

	// Make request
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	response, err := client.Do(request)
	if err != nil {
		logger.Error(err)
		return
	}
	defer response.Body.Close()

	filteredURLs := []string{}
	if response.StatusCode >= 200 && response.StatusCode <= 299 {
		// Create goquery document from response
		document, err := goquery.NewDocumentFromReader(response.Body)
		if err != nil {
			logger.Error(err)
			return
		}

		// Parse all URLs in the document
		allURLs := parser.ParseAllURLs(document, crawlableURL, r.WebsiteHost)
		filteredURLs = append(filteredURLs, r.CrawlObj.FilterOutAndUpdateCrawledURLs(allURLs)...)

		for _, url := range filteredURLs {
			r.CrawlQueue <- url
		}
	}

	// Record the sitemap
	// In order to record all unique links
	r.CrawlObj.SitemapObj.AddLinks(crawlableURL, filteredURLs)

	// In order to record all links (can be duplicate as well)
	// r.CrawlObj.SitemapObj.AddLinks(crawlableURL, allURLs)
	return
}
