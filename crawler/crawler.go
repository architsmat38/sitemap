package crawler

import (
	"github.com/architsmat38/sitemap/pool"
	"github.com/architsmat38/sitemap/sitemap"
)

/**
 * Structure of request to be processed by crawler
 */
type Request struct {
	WebsiteURL string
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
func NewCrawler(size int, queue chan string) *CrawlerObj {
	s := sitemap.NewSitemapManager()
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
 * Process the crawler request
 */
func (r Request) Execute() {

}
