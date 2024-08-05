package scraper

import "github.com/gocolly/colly/v2"

type Uaserials interface {
	ScrapMainPage()
	ScrapOneMovie(URLPath string)
}

func NewUaserialsScraper(c *colly.Collector, rootUrl string) Uaserials {
	return &Uaserialer{
		c:       c,
		rootUrl: rootUrl,
	}
}

type Scraper struct {
	Uaserials Uaserials
}

func NewScraper(c *colly.Collector, uaserialsRootUrl string) *Scraper {
	return &Scraper{
		Uaserials: NewUaserialsScraper(c, uaserialsRootUrl),
	}
}
