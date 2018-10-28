package sitemap

import (
	"fmt"
	"sync"
)

/**
 * Structure of links object which will be stored in sitemap
 */
type LinksObj struct {
	links map[string][]*LinksObj
}

/**
 * Structure of sitemap object
 */
type SitemapObj struct {
	sitemapLinks LinksObj
	sitemapChan  chan LinksObj
}

/**
 * Creates sitemap manager to store the sitemap links
 */
func NewSitemapManager(linkURL string) *SitemapObj {
	// Links object
	l := LinksObj{
		links: make(map[string][]*LinksObj),
	}
	linkObj := make(map[string][]*LinksObj)
	linkObj[linkURL] = []*LinksObj{}
	l.links = linkObj

	// Sitemap object
	s := SitemapObj{
		sitemapLinks: l,
		sitemapChan:  make(chan LinksObj),
	}
	return &s
}

/**
 * Utility function to add links. It iterates over the sitemap and add links at accurate position.
 */
func addLinks(sitemapLinksObj *LinksObj, linkURL string, hrefLinks []string) bool {
	allLinksObj, ok := sitemapLinksObj.links[linkURL]
	if ok {
		for _, linkURL := range hrefLinks {
			var linkExist bool
			for _, linksObj := range allLinksObj {
				if _, ok := linksObj.links[linkURL]; ok {
					linkExist = true
					break
				}
			}

			if !linkExist {
				linkObj := make(map[string][]*LinksObj)
				linkObj[linkURL] = []*LinksObj{}
				allLinksObj = append(allLinksObj, &LinksObj{links: linkObj})
			}
		}
		sitemapLinksObj.links[linkURL] = allLinksObj
		return true
	} else {
		var isAdded bool
		for _, allLinksObj := range sitemapLinksObj.links {
			for _, linksObj := range allLinksObj {
				isAdded = addLinks(linksObj, linkURL, hrefLinks)
				if isAdded {
					break
				}
			}

			if isAdded {
				break
			}
		}
	}
	return false
}

/**
 * Add new links in the sitemap
 */
func (s *SitemapObj) AddLinks(linkURL string, hrefLinks []string) {
	var mu sync.RWMutex
	mu.Lock()
	addLinks(&s.sitemapLinks, linkURL, hrefLinks)
	mu.Unlock()
}

/**
 * Utility function to print sitemap
 */
func print(sitemapLinksObj LinksObj, prefix string) {
	for linkURL, allLinksObj := range sitemapLinksObj.links {
		fmt.Println(prefix, linkURL)
		for _, linksObj := range allLinksObj {
			print(*linksObj, "|	"+prefix)
		}
	}
}

/**
 * Prints the sitemap
 */
func (s *SitemapObj) Print() {
	fmt.Println("\n###################################### SITEMAP ######################################\n")
	print(s.sitemapLinks, "|--")
}
