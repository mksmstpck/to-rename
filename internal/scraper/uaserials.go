package scraper

import (
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
	models "github.com/mksmstpck/to-rename/internal/scrapModels"
	"github.com/sirupsen/logrus"
)

func (s *UaserialsScraper) ScrapMainPage() {
	var titles []string

	s.c.OnHTML("div.th-title.truncate", func(h *colly.HTMLElement) {
		titles = append(titles, h.Text)
	})

	s.c.OnRequest(func(r *colly.Request) {
		logrus.Printf("Visiting - %s\n", r.URL)
	})

	s.c.OnError(func(r *colly.Response, err error) {
		logrus.Print(err)
	})

	s.c.Visit("https://uaserials.pro/")

	logrus.Println(titles)
}

func (s *UaserialsScraper) ScrapOneMovie(URLPath string) {
	var m models.MovieFull

	// parses name in ukrainian
	s.c.OnHTML("span.oname_ua", func(h *colly.HTMLElement) {
		m.Name = h.Text
	})

	// parses name in english
	s.c.OnHTML("div.oname", func(h *colly.HTMLElement) {
		m.NameEng = h.Text
	})

	// parses poster link
	s.c.OnHTML("img", func(h *colly.HTMLElement) {
		m.PosterLink = h.Attr("src")
	})

	// parses list wtih info
	s.c.OnHTML("ul.short-list.fx-1", func(h *colly.HTMLElement) {
		h.ForEach("li", func(i int, h *colly.HTMLElement) {
			// parses phrase
			if strings.Contains(h.Text, "Гасло") {
				// deletes text from <span> and leaves only from <li>
				m.Phrase = strings.Replace(h.Text, h.ChildText("span"), "", -1)
			}
			// parses year
			if strings.Contains(h.Text, "Рік") {
				// deletes text from <span> and leaves only from <li>
				yearStr := strings.Replace(h.Text, h.ChildText("span"), "", -1)
				// deletes " " from string for formating
				yearStr = strings.Replace(yearStr, " ", "", -1)
				year, err := strconv.Atoi(yearStr)
				if err != nil {
					logrus.Println(err)
				}
				m.Year = year
			}
			// parses Genres
			if strings.Contains(h.Text, "Жанр") {
				h.ForEach("a", func(i int, h *colly.HTMLElement) {
					m.Genres = append(m.Genres, h.Text)
				})
			}
			// parses Countries
			if strings.Contains(h.Text, "Країна") {
				h.ForEach("a", func(i int, h *colly.HTMLElement) {
					m.Countries = append(m.Countries, h.Text)
				})
			}
			// parses translations
			if strings.Contains(h.Text, "Переклад") {
				h.ForEach("span", func(i int, h *colly.HTMLElement) {
					if i == 1 {
						m.Translation = h.Text
					}
				})
			}
			// parses directors
			if strings.Contains(h.Text, "Режисер") {
				plainStr := strings.Replace(h.Text, h.ChildText("span"), "", -1)
				m.Director = strings.Split(plainStr, ", ")
			}
			// parses actors
			if strings.Contains(h.Text, "Актори") {
				plainStr := strings.Replace(h.Text, h.ChildText("span"), "", -1)
				m.Actors = strings.Split(plainStr, ", ")
			}
		})
	})

	// parses descriptions
	s.c.OnHTML("div.ftext.full-text.cleasrfix", func(h *colly.HTMLElement) {
		m.Description = h.Text
	})

	// parses min age
	s.c.OnHTML("div.short-rate-in.short-agerating", func(h *colly.HTMLElement) {
		ageStr := strings.Replace(h.Text, "+", "", -1)
		age, err := strconv.Atoi(ageStr)
		if err != nil {
			logrus.Error(err)
		}

		m.Age = age
	})

	// parses rating
	s.c.OnHTML("a.short-rate-in.short-rate-imdb.cursor-help", func(h *colly.HTMLElement) {
		rating, err := strconv.ParseFloat(h.Text, 8)
		if err != nil {
			logrus.Error(err)
		}

		m.Rating = rating
	})

	s.c.OnError(func(r *colly.Response, err error) {
		logrus.Print(err)
	})

	s.c.Visit(s.rootUrl + URLPath)

	logrus.Println(m)
}
