package scraper

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
	"github.com/zipqt/goScrape/controllers"
	"github.com/zipqt/goScrape/models"
)

func Start() {
	for j := 1; j <= 25; j++ {
		urlToScrape := fmt.Sprintf("https://www.freepik.com/free-photos-vectors/kawaii-coloring-page/%d", j)
		fmt.Printf("PAGE %d DONE\n", j)
		for i := 1; i <= 50; i++ {
			currentI := i
			c := colly.NewCollector()
			c.UserAgent = "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0"
			xpathQuery := fmt.Sprintf("/html/body/main/div[3]/div/div[2]/section/figure[%d]/div/a/img", currentI)

			c.OnXML(xpathQuery, func(e *colly.XMLElement) {
				src := e.Attr("data-src")
				fmt.Printf("Image #%d Src: %s\n", currentI, src)
				page := models.Page{
					URL: src,
				}
				if err := controllers.Save(&page); err != nil {
					log.Fatal(err)
				}
			})
			c.Visit(urlToScrape)
			c.Wait()
		}
	}
}
