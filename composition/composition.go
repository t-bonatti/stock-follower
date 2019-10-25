package composition

import (
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type composition struct {
	Name       string
	Code       string
	Percentage float64
}

// GetComposition get monetus composition
func GetComposition() []composition {
	// Request the HTML page.
	res, err := http.Get("https://monetus.com.br/comparador/fundos-de-investimento/monetus-fia")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var items []composition

	doc.Find(".sc-fjdhpX .details div.sc-iQtOjA").Each(func(i int, s *goquery.Selection) {
		name := s.Find("span:nth-child(1)").First().Text()
		reg := regexp.MustCompile(`\((.*)\)`)
		stock := reg.FindStringSubmatch(name)[1]

		percentageString := s.Find("span:nth-child(2)").First().Text()
		percentageString = strings.Replace(percentageString, "%", "", 1)
		percentageString = strings.Replace(percentageString, ",", ".", 1)
		percentage, _ := strconv.ParseFloat(percentageString, 64)

		items = append(items, composition{name, stock, percentage})
	})

	return items
}
