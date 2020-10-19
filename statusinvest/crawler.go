package statusinvest

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/shopspring/decimal"
)

//Crawler interface
type Crawler interface {
	Get(string) (StatusInvest, error)
}

type crawler struct {
	Client  *http.Client
	baseURL string
}

// NewCrawler create new Crawler
func NewCrawler() Crawler {
	return &crawler{&http.Client{}, "https://statusinvest.com.br"}
}

func toDecimal(strValue string) decimal.Decimal {
	if strValue == "-" || strValue == "-%" || strValue == "- %" || strValue == "R$ -" {
		return decimal.Zero
	}
	processedValue := strings.ReplaceAll(strValue, "%", "")
	processedValue = strings.ReplaceAll(processedValue, "R$", "")
	processedValue = strings.ReplaceAll(processedValue, ".", "")
	processedValue = strings.ReplaceAll(processedValue, ",", ".")
	processedValue = strings.TrimSpace(processedValue)

	value, err := decimal.NewFromString(processedValue)
	if err != nil {
		log.Println(fmt.Sprintf("[CRITICAL] Erro ao converter valor: %s", strValue))
		return decimal.Zero
	}
	return value
}

// GetActualValue Return Actual Value of stock
func (c *crawler) Get(name string) (data StatusInvest, err error) {
	r, err := c.Client.Get(c.baseURL + "/acoes/" + name)
	if err != nil {
		return
	}
	defer func() {
		r.Body.Close()
	}()

	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		log.Fatal("[ScrapIndice] Can´t read body from baseUrl")
	}

	fmt.Println(name)

	data.Name = name
	doc.Find("#company-section").Each(func(i int, s *goquery.Selection) {
		data.Sector = s.Find("#company-section > div > div.card.bg-main-gd-h.white-text.rounded > div > div:nth-child(1) > div > div > div > a > strong").Text()
		data.SubSector = s.Find("#company-section > div > div.card.bg-main-gd-h.white-text.rounded > div > div.info.pl-md-2.pr-md-2 > div > div > div > a > strong").Text()
		data.Segment = s.Find("#company-section > div > div.card.bg-main-gd-h.white-text.rounded > div > div:nth-child(3) > div > div > div > a > strong").Text()
	})
	doc.Find("div .paper").Each(func(i int, s *goquery.Selection) {
		data.Value = toDecimal(s.Find("#main-2 > div:nth-child(4) > div > div.pb-3.pb-md-5 > div > div.info.special.w-100.w-md-33.w-lg-20 > div > div:nth-child(1) > strong").Text())
		data.MinInLast12Month = toDecimal(s.Find("#main-2 > div:nth-child(4) > div > div.pb-3.pb-md-5 > div > div:nth-child(2) > div > div:nth-child(1) > strong").Text())
		data.MaxInLast12Month = toDecimal(s.Find("#main-2 > div:nth-child(4) > div > div.pb-3.pb-md-5 > div > div:nth-child(3) > div > div:nth-child(1) > strong").Text())
		data.PaymentLast12Month = toDecimal(s.Find("#main-2 > div:nth-child(4) > div > div.pb-3.pb-md-5 > div > div:nth-child(4) > div > div.d-flex.justify-between > div > span.sub-value").Text())
		data.DividendYield = toDecimal(s.Find("#main-2 > div:nth-child(4) > div > div.pb-3.pb-md-5 > div > div:nth-child(4) > div > div:nth-child(1) > strong").Text())
		data.MaxInMonth = toDecimal(s.Find("#main-2 > div:nth-child(4) > div > div.pb-3.pb-md-5 > div > div:nth-child(3) > div > div.d-flex.justify-between > div > span.sub-value").Text())
		data.MinInMonth = toDecimal(s.Find("#main-2 > div:nth-child(4) > div > div.pb-3.pb-md-5 > div > div:nth-child(2) > div > div.d-flex.justify-between > div > span.sub-value").Text())
		data.VariationLast12Month = toDecimal(s.Find("#main-2 > div:nth-child(4) > div > div.pb-3.pb-md-5 > div > div:nth-child(5) > div > div:nth-child(1) > strong").Text())
		data.VariationInMonth = toDecimal(s.Find("#main-2 > div:nth-child(4) > div > div.pb-3.pb-md-5 > div > div:nth-child(5) > div > div.d-flex.justify-between > div > span.sub-value > b").Text())

		data.PL = toDecimal(s.Find("#main-2 > div:nth-child(4) > div > div:nth-child(5) > div > div:nth-child(1) > div > div:nth-child(1) > div > div > strong").Text())
		data.EVEbitda = toDecimal(s.Find("#main-2 > div:nth-child(4) > div > div:nth-child(5) > div > div:nth-child(1) > div > div:nth-child(2) > div > div > strong").Text())
		data.PVP = toDecimal(s.Find("#main-2 > div:nth-child(4) > div > div:nth-child(5) > div > div:nth-child(1) > div > div:nth-child(3) > div > div > strong").Text())
		data.EVEbit = toDecimal(s.Find("#main-2 > div:nth-child(4) > div > div:nth-child(5) > div > div:nth-child(1) > div > div:nth-child(4) > div > div > strong").Text())
		data.PEbitda = toDecimal(s.Find("#main-2 > div:nth-child(4) > div > div:nth-child(5) > div > div:nth-child(1) > div > div:nth-child(5) > div > div > strong").Text())
		data.PEbit = toDecimal(s.Find("#main-2 > div:nth-child(4) > div > div:nth-child(5) > div > div:nth-child(1) > div > div:nth-child(6) > div > div > strong").Text())

		data.VPA = toDecimal(s.Find("#main-2 > div:nth-child(4) > div > div:nth-child(5) > div > div:nth-child(1) > div > div:nth-child(7) > div > div > strong").Text())
		data.PAtivo = toDecimal(s.Find("#main-2 > div:nth-child(4) > div > div:nth-child(5) > div > div:nth-child(1) > div > div:nth-child(8) > div > div > strong").Text())
		data.LPA = toDecimal(s.Find("#main-2 > div:nth-child(4) > div > div:nth-child(5) > div > div:nth-child(1) > div > div:nth-child(9) > div > div > strong").Text())
		data.PSR = toDecimal(s.Find("#main-2 > div:nth-child(4) > div > div:nth-child(5) > div > div:nth-child(1) > div > div:nth-child(10) > div > div > strong").Text())
		data.PWorkingCapital = toDecimal(s.Find("#main-2 > div:nth-child(4) > div > div:nth-child(5) > div > div:nth-child(1) > div > div:nth-child(11) > div > div > strong").Text())
		data.PCurrentAssetsNet = toDecimal(s.Find("#main-2 > div:nth-child(4) > div > div:nth-child(5) > div > div:nth-child(1) > div > div:nth-child(12) > div > div > strong").Text())

		data.NetDebtPatrimony = toDecimal(s.Find("#main-2 > div:nth-child(4) > div > div:nth-child(5) > div > div:nth-child(2) > div > div:nth-child(1) > div > div > strong").Text())
		data.NetDebtEbitda = toDecimal(s.Find("#main-2 > div:nth-child(4) > div > div:nth-child(5) > div > div:nth-child(2) > div > div:nth-child(2) > div > div > strong").Text())
		data.NetDebtEbit = toDecimal(s.Find("#main-2 > div:nth-child(4) > div > div:nth-child(5) > div > div:nth-child(2) > div > div:nth-child(3) > div > div > strong").Text())
		data.PatrimonyAssets = toDecimal(s.Find("#main-2 > div:nth-child(4) > div > div:nth-child(5) > div > div:nth-child(2) > div > div:nth-child(4) > div > div > strong").Text())
		data.LiabilitiesAssets = toDecimal(s.Find("#main-2 > div:nth-child(4) > div > div:nth-child(5) > div > div:nth-child(2) > div > div:nth-child(5) > div > div > strong").Text())
		data.CurrentLiquidity = toDecimal(s.Find("#main-2 > div:nth-child(4) > div > div:nth-child(5) > div > div:nth-child(2) > div > div:nth-child(6) > div > div > strong").Text())

		data.GrossMargin = toDecimal(s.Find("#main-2 > div:nth-child(4) > div > div:nth-child(5) > div > div:nth-child(3) > div > div:nth-child(1) > div > div > strong").Text())
		data.EbitdaMargin = toDecimal(s.Find("#main-2 > div:nth-child(4) > div > div:nth-child(5) > div > div:nth-child(3) > div > div:nth-child(2) > div > div > strong").Text())
		data.ROE = toDecimal(s.Find("#main-2 > div:nth-child(4) > div > div:nth-child(5) > div > div:nth-child(4) > div > div:nth-child(1) > div > div > strong").Text())
		data.ROA = toDecimal(s.Find("#main-2 > div:nth-child(4) > div > div:nth-child(5) > div > div:nth-child(4) > div > div:nth-child(2) > div > div > strong").Text())
		data.CAGRRevenue = toDecimal(s.Find("#main-2 > div:nth-child(4) > div > div:nth-child(5) > div > div:nth-child(5) > div > div:nth-child(1) > div > div > strong").Text())
		data.CAGRProfit = toDecimal(s.Find("#main-2 > div:nth-child(4) > div > div:nth-child(5) > div > div:nth-child(5) > div > div:nth-child(2) > div > div > strong").Text())

		data.EbitMargin = toDecimal(s.Find("#main-2 > div:nth-child(4) > div > div:nth-child(5) > div > div:nth-child(3) > div > div:nth-child(3) > div > div > strong").Text())
		data.NetMargin = toDecimal(s.Find("#main-2 > div:nth-child(4) > div > div:nth-child(5) > div > div:nth-child(3) > div > div:nth-child(4) > div > div > strong").Text())
		data.ROIC = toDecimal(s.Find("#main-2 > div:nth-child(4) > div > div:nth-child(5) > div > div:nth-child(4) > div > div:nth-child(3) > div > div > strong").Text())
		data.TurnAsset = toDecimal(s.Find("#main-2 > div:nth-child(4) > div > div:nth-child(5) > div > div:nth-child(4) > div > div:nth-child(4) > div > div > strong").Text())

	})

	return
}
