package scraper

import (
	"github.com/gocolly/colly/v2"
	models "github.com/mksmstpck/to-rename/services/scraper/scraper/models"
	"github.com/mksmstpck/to-rename/services/scraper/translate"
)

type Imdber struct {
	c       *colly.Collector
	t       *translate.Translate
	rootUrl string
}

func ParseByName(name string) (*models.MovieFull, error) {
	return nil, nil
}
