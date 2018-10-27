package sitemap

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
	allLinks     map[string]bool
	sitemapLinks LinksObj
	sitemapChan  chan LinksObj
}

/**
 * Creates sitemap manager to store the sitemap links
 */
func NewSitemapManager() *SitemapObj {
	// Links object
	l := LinksObj{
		links: make(map[string][]*LinksObj),
	}

	// Sitemap object
	s := SitemapObj{
		allLinks:     make(map[string]bool),
		sitemapLinks: l,
		sitemapChan:  make(chan LinksObj),
	}
	return &s
}

/**
 * Add new links in the sitemap
 */
func (s *SitemapObj) addLinks(fetchedLinks LinksObj) {

}

/**
 * Prints the sitemap
 */
func (s *SitemapObj) Print() {

}
