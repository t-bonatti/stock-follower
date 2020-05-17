package client

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/shopspring/decimal"
)

//StatusInvest interface
type StatusInvest interface {
	Get(string) (StatusInvestData, error)
}

type statusInvest struct {
	Client  *http.Client
	baseURL string
}

//StatusInvestData data
type StatusInvestData struct {
	Name                 string          `json:"Nome"`
	Sector               string          `json:"Sector"`
	SubSector            string          `json:"SubSector"`
	Segment              string          `json:"Segment"`
	Value                decimal.Decimal `json:"Value"`
	VariationLast12Month decimal.Decimal `json:"VariationLast12Month"`
	MinLastLast12Month   decimal.Decimal `json:"MinLastLast12Month"`
	MaxInLast12Month     decimal.Decimal `json:"MaxInLast12Month"`
	PaymentLast12Month   decimal.Decimal `json:"PaymentLast12Month"`
	DividendYield        decimal.Decimal `json:"DividendYield"`
	VariationInMonth     decimal.Decimal `json:"VariationInMonth"`
	MinInMonth           decimal.Decimal `json:"MinInMonth"`
	MaxInMonth           decimal.Decimal `json:"MaxInMonth"`
	AppreciationInMonth  decimal.Decimal `json:"AppreciationInMonth"`
	PVP                  decimal.Decimal `json:"PVP"`
	PL                   decimal.Decimal `json:"PL"`
	PEbitda              decimal.Decimal `json:"PEbitda"`
	PEbit                decimal.Decimal `json:"PEbit"`
	PAtivo               decimal.Decimal `json:"PAtivo"`
	EVEbitda             decimal.Decimal `json:"EVEbitda"`
	EVEbit               decimal.Decimal `json:"EVEbit"`
	PSR                  decimal.Decimal `json:"PSR"`
	PWorkingCapital      decimal.Decimal `json:"PWorkingCapital"`
	PCurrentAssetsNet    decimal.Decimal `json:"PCurrentAssetsNet"`
	GrossMargin          decimal.Decimal `json:"GrossMargin"`
	EbitdaMargin         decimal.Decimal `json:"EbitdaMargin"`
	EbitMargin           decimal.Decimal `json:"EbitMargin"`
	NetMargin            decimal.Decimal `json:"NetMargin"`
	TurnAsset            decimal.Decimal `json:"TurnAsset"`
	ROE                  decimal.Decimal `json:"ROE"`
	ROA                  decimal.Decimal `json:"ROA"`
	ROIC                 decimal.Decimal `json:"ROIC"`
	LPA                  decimal.Decimal `json:"LPA"`
	VPA                  decimal.Decimal `json:"VPA"`
	NetDebtPatrimony     decimal.Decimal `json:"NetDebtPatrimony"`
	NetDebtEbitda        decimal.Decimal `json:"NetDebtEbitda"`
	NetDebtEbit          decimal.Decimal `json:"NetDebtEbit"`
	PatrimonyAssets      decimal.Decimal `json:"PatrimonyAssets"`
	LiabilitiesAssets    decimal.Decimal `json:"LiabilitiesAssets"`
	CurrentLiquidity     decimal.Decimal `json:"CurrentLiquidity"`
	CAGRRevenue          decimal.Decimal `json:"CAGRRevenue"`
	CAGRProfit           decimal.Decimal `json:"CAGRProfit"`
}

// NewStatusInvest create new StatusInvest
func NewStatusInvest() StatusInvest {
	return &statusInvest{&http.Client{}, "https://statusinvest.com.br"}
}

func toDecimal(strValue string) decimal.Decimal {
	if strValue == "-" || strValue == "-%" {
		return decimal.Zero
	}
	processedValue := strings.ReplaceAll(strValue, "%", "")
	processedValue = strings.ReplaceAll(processedValue, "R$", "")
	processedValue = strings.ReplaceAll(processedValue, ".", "")
	processedValue = strings.ReplaceAll(processedValue, ",", ".")
	processedValue = strings.TrimSpace(processedValue)

	value, err := decimal.NewFromString(processedValue)
	if err != nil {
		log.Println("[CRITICAL] Erro ao converter valor")
		return decimal.Zero
	}
	return value
}

// GetActualValue Return Actual Value of stock
func (s *statusInvest) Get(name string) (data StatusInvestData, err error) {
	r, err := s.Client.Get(s.baseURL + "/acoes/" + name)
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

	data.Name = name
	doc.Find("#company-section").Each(func(i int, s *goquery.Selection) {
		data.Sector = s.Find("#company-section > div > div.card.bg-main-gd-h.white-text.rounded > div > div:nth-child(1) > div > div > div > a > strong").Text()
		data.SubSector = s.Find("#company-section > div > div.card.bg-main-gd-h.white-text.rounded > div > div.info.pl-md-2.pr-md-2 > div > div > div > a > strong").Text()
		data.Segment = s.Find("#company-section > div > div.card.bg-main-gd-h.white-text.rounded > div > div:nth-child(3) > div > div > div > a > strong").Text()
	})
	doc.Find("div .paper").Each(func(i int, s *goquery.Selection) {
		data.Value = toDecimal(s.Find("#main-2 > div:nth-child(3) > div > div.pb-3.pb-md-5 > div > div.info.special.w-100.w-md-33.w-lg-20 > div > div:nth-child(1) > strong").Text())
		data.VariationLast12Month = toDecimal(s.Find("#main-2 > div:nth-child(3) > div > div.pb-3.pb-md-5 > div > div:nth-child(5) > div > div:nth-child(1) > strong").Text())
		data.MinLastLast12Month = toDecimal(s.Find("#main-2 > div:nth-child(3) > div > div.pb-3.pb-md-5 > div > div:nth-child(2) > div > div:nth-child(1) > strong").Text())
		data.MaxInLast12Month = toDecimal(s.Find("#main-2 > div:nth-child(3) > div > div.pb-3.pb-md-5 > div > div:nth-child(3) > div > div:nth-child(1) > strong").Text())
		data.PaymentLast12Month = toDecimal(s.Find("#main-2 > div:nth-child(3) > div > div.pb-3.pb-md-5 > div > div:nth-child(4) > div > div.d-flex.justify-between > div > span.sub-value").Text())
		data.DividendYield = toDecimal(s.Find("#main-2 > div:nth-child(3) > div > div.pb-3.pb-md-5 > div > div:nth-child(4) > div > div:nth-child(1) > strong").Text())
		data.MaxInMonth = toDecimal(s.Find("#main-2 > div:nth-child(3) > div > div.pb-3.pb-md-5 > div > div:nth-child(3) > div > div.d-flex.justify-between > div > span.sub-value").Text())
		data.MinInMonth = toDecimal(s.Find("#main-2 > div:nth-child(3) > div > div.pb-3.pb-md-5 > div > div:nth-child(2) > div > div.d-flex.justify-between > div > span.sub-value").Text())
		data.AppreciationInMonth = toDecimal(s.Find("#main-2 > div:nth-child(3) > div > div.pb-3.pb-md-5 > div > div:nth-child(5) > div > div.d-flex.justify-between > div > span.sub-value > b").Text())
		data.VariationInMonth = toDecimal(s.Find("#main-2 > div:nth-child(3) > div > div.pb-3.pb-md-5 > div > div:nth-child(5) > div > div:nth-child(1) > strong").Text())
		data.PVP = toDecimal(s.Find("#main-2 > div:nth-child(3) > div > div:nth-child(5) > div > div:nth-child(1) > div > div > strong").Text())
		data.PL = toDecimal(s.Find("#main-2 > div:nth-child(3) > div > div:nth-child(5) > div > div:nth-child(2) > div > div > strong").Text())
		data.PEbitda = toDecimal(s.Find("#main-2 > div:nth-child(3) > div > div:nth-child(5) > div > div:nth-child(3) > div > div > strong").Text())
		data.PEbit = toDecimal(s.Find("#main-2 > div:nth-child(3) > div > div:nth-child(5) > div > div:nth-child(4) > div > div > strong").Text())
		data.PAtivo = toDecimal(s.Find("#main-2 > div:nth-child(3) > div > div:nth-child(5) > div > div:nth-child(5) > div > div > strong").Text())
		data.EVEbitda = toDecimal(s.Find("#main-2 > div:nth-child(3) > div > div:nth-child(5) > div > div:nth-child(6) > div > div > strong").Text())
		data.EVEbit = toDecimal(s.Find("#main-2 > div:nth-child(3) > div > div:nth-child(5) > div > div:nth-child(7) > div > div > strong").Text())
		data.PSR = toDecimal(s.Find("#main-2 > div:nth-child(3) > div > div:nth-child(5) > div > div:nth-child(8) > div > div > strong").Text())
		data.PWorkingCapital = toDecimal(s.Find("#main-2 > div:nth-child(3) > div > div:nth-child(5) > div > div:nth-child(9) > div > div > strong").Text())
		data.PCurrentAssetsNet = toDecimal(s.Find("#main-2 > div:nth-child(3) > div > div:nth-child(5) > div > div:nth-child(10) > div > div > strong").Text())
		data.GrossMargin = toDecimal(s.Find("#main-2 > div:nth-child(3) > div > div:nth-child(5) > div > div:nth-child(11) > div > div > strong").Text())
		data.EbitdaMargin = toDecimal(s.Find("#main-2 > div:nth-child(3) > div > div:nth-child(5) > div > div:nth-child(12) > div > div > strong").Text())
		data.EbitMargin = toDecimal(s.Find("#main-2 > div:nth-child(3) > div > div:nth-child(5) > div > div:nth-child(13) > div > div > strong").Text())
		data.NetMargin = toDecimal(s.Find("#main-2 > div:nth-child(3) > div > div:nth-child(5) > div > div:nth-child(14) > div > div > strong").Text())
		data.TurnAsset = toDecimal(s.Find("#main-2 > div:nth-child(3) > div > div:nth-child(5) > div > div:nth-child(15) > div > div > strong").Text())
		data.ROE = toDecimal(s.Find("#main-2 > div:nth-child(3) > div > div:nth-child(5) > div > div:nth-child(16) > div > div > strong").Text())
		data.ROA = toDecimal(s.Find("#main-2 > div:nth-child(3) > div > div:nth-child(5) > div > div:nth-child(17) > div > div > strong").Text())
		data.ROIC = toDecimal(s.Find("#main-2 > div:nth-child(3) > div > div:nth-child(5) > div > div:nth-child(18) > div > div > strong").Text())
		data.LPA = toDecimal(s.Find("#main-2 > div:nth-child(3) > div > div:nth-child(5) > div > div:nth-child(19) > div > div > strong").Text())
		data.VPA = toDecimal(s.Find("#main-2 > div:nth-child(3) > div > div:nth-child(5) > div > div:nth-child(20) > div > div > strong").Text())
		data.NetDebtPatrimony = toDecimal(s.Find("#main-2 > div:nth-child(3) > div > div:nth-child(5) > div > div:nth-child(21) > div > div > strong").Text())
		data.NetDebtEbitda = toDecimal(s.Find("#main-2 > div:nth-child(3) > div > div:nth-child(5) > div > div:nth-child(22) > div > div > strong").Text())
		data.NetDebtEbit = toDecimal(s.Find("#main-2 > div:nth-child(3) > div > div:nth-child(5) > div > div:nth-child(23) > div > div > strong").Text())
		data.PatrimonyAssets = toDecimal(s.Find("#main-2 > div:nth-child(3) > div > div:nth-child(5) > div > div:nth-child(24) > div > div > strong").Text())
		data.LiabilitiesAssets = toDecimal(s.Find("#main-2 > div:nth-child(3) > div > div:nth-child(5) > div > div:nth-child(25) > div > div > strong").Text())
		data.CurrentLiquidity = toDecimal(s.Find("#main-2 > div:nth-child(3) > div > div:nth-child(5) > div > div:nth-child(26) > div > div > strong").Text())
		data.CAGRRevenue = toDecimal(s.Find("#main-2 > div:nth-child(3) > div > div:nth-child(5) > div > div:nth-child(27) > div > div > strong").Text())
		data.CAGRProfit = toDecimal(s.Find("#main-2 > div:nth-child(3) > div > div:nth-child(5) > div > div:nth-child(28) > div > div > strong").Text())
	})
	return
}
