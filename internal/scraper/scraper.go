package scraper

import "github.com/gocolly/colly/v2"

type UaserialsScrapers interface {
	ScrapMainPage()
	ScrapOneMovie(URLPath string)
}

type UaserialsScraper struct {
	c       *colly.Collector
	rootUrl string
}

func NewUaserialsScraper(c *colly.Collector, rootUrl string) UaserialsScrapers {
	return &UaserialsScraper{
		c:       c,
		rootUrl: rootUrl,
	}
}

type Scraper struct {
	Uaserials UaserialsScrapers
}

func NewScraper(c *colly.Collector, uaserialsRootUrl string) *Scraper {
	return &Scraper{
		Uaserials: NewUaserialsScraper(c, uaserialsRootUrl),
	}
}
