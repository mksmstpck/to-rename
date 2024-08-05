package main

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/mksmstpck/to-rename/services/gateway/views"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetReportCaller(true)
	formatter := &logrus.TextFormatter{
		TimestampFormat:        "02-01-2006 15:04:05",
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf("%s:%d", formatFilePath(f.Function), f.Line)
		},
	}
	logrus.SetFormatter(formatter)
}

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}

func main() {
	//	c := colly.NewCollector(colly.AllowedDomains("uaserials.pro"))

	//	scraper.NewScraper(c, "https://uaserials.pro").Uaserials.ScrapOneMovie("/8599-barva-purpurova-2023.html")

	views.NewHandler().HandleAll()
}
