package scraper

type MovieFull struct {
	Name        string
	NameEng     string
	PosterLink  string
	Phrase      string
	Year        int
	Genres      []string
	Countries   []string
	Translation string
	Director    []string
	Actors      []string
	Description string
	Age         int
	Rating      float64
	URL         string
}
