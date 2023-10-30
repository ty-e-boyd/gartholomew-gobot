package scrapers

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

type XurInfo struct {
	Location string
}

func ScrapeXur() string {
	c := colly.NewCollector()
	r := "Unable to pull location."

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"

	// our first (and only right now) OnHtml handler that will run on any successfully hit html files
	c.OnHTML(".et_pb_countdown_timer_container", func(e *colly.HTMLElement) {
		//url := e.ChildAttr(".crayons-story__hidden-navigation-link", "href")
		//name := e.ChildText(".crayons-story__hidden-navigation-link")
		location := e.ChildText(".title")

		fmt.Printf("found location: %v\n\n", location)

		if location != "" {
			xurInfo := XurInfo{
				Location: location,
			}

			r = xurInfo.Location
		} else {
			fmt.Println("could not pull location")
		}
	})

	// going and attempting the scrape
	err := c.Visit("https://whereisxur.com/")
	if err != nil {
		log.Fatal("error on visit --")
	}

	// some results of our scraping
	return r
}
