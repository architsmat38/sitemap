package parser

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/architsmat38/sitemap/utils"
)

/**
 * Parses all URLs present in the document, and also filters out irrelevant URLs (i.e. with different host)
 */
func ParseAllURLs(document *goquery.Document, linkURL string, websiteHost string) []string {
	allURLs := document.Find("a").FilterFunction(func(key int, element *goquery.Selection) bool {
		// See if the href attribute exists on the element
		href, exists := element.Attr("href")
		validURL := utils.CreateValidURL(href, websiteHost)
		if exists {
			// Checks if host is same for the crawled website
			host, err := utils.GetHost(validURL)
			if err != nil {
				return false
			}
			return websiteHost == host
		}
		return false
	}).Map(func(key int, element *goquery.Selection) string {
		href, _ := element.Attr("href")
		validURL := utils.CreateValidURL(href, websiteHost)
		return utils.GetCrawlableURL(validURL)
	})

	return allURLs
}
