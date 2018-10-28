package sitemap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSitemapManager(t *testing.T) {
	linkURL := "http://example.com"
	s := NewSitemapManager(linkURL)

	_, ok := s.sitemapLinks.links[linkURL]
	assert.True(t, ok)

	linksObj := []*LinksObj{}
	assert.EqualValues(t, s.sitemapLinks.links[linkURL], linksObj)
}

func TestAddLinks(t *testing.T) {
	linkURL := "http://example.com"
	s := NewSitemapManager(linkURL)

	hrefLinks := []string{"http://example.com/abc", "http://example.com/xyz"}
	c := make(chan bool, 5)
	s.AddLinks(linkURL, hrefLinks, &c)

	_, ok := s.sitemapLinks.links[linkURL]
	assert.True(t, ok)

	allLinksObj := []*LinksObj{}
	for _, val := range hrefLinks {
		linkObj := make(map[string][]*LinksObj)
		linkObj[val] = []*LinksObj{}
		allLinksObj = append(allLinksObj, &LinksObj{links: linkObj})
	}
	assert.EqualValues(t, s.sitemapLinks.links[linkURL], allLinksObj)
}
