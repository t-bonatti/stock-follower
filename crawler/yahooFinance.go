package crawler

import (
	"encoding/json"
	"net/http"
)

//YahooFinance interface
type YahooFinance interface {
	GetActualValue(string) (float64, error)
}

type yahooFinance struct {
	Client  *http.Client
	baseURL string
}

type yahooFinanceResponse struct {
	Chart struct {
		Result []struct {
			Meta struct {
				Currency             string  `json:"currency"`
				Symbol               string  `json:"symbol"`
				ExchangeName         string  `json:"exchangeName"`
				InstrumentType       string  `json:"instrumentType"`
				FirstTradeDate       int     `json:"firstTradeDate"`
				RegularMarketTime    int     `json:"regularMarketTime"`
				Gmtoffset            int     `json:"gmtoffset"`
				Timezone             string  `json:"timezone"`
				ExchangeTimezoneName string  `json:"exchangeTimezoneName"`
				RegularMarketPrice   float64 `json:"regularMarketPrice"`
				ChartPreviousClose   float64 `json:"chartPreviousClose"`
				PreviousClose        float64 `json:"previousClose"`
				Scale                int     `json:"scale"`
				PriceHint            int     `json:"priceHint"`
			} `json:"meta"`
		} `json:"result"`
	} `json:"chart"`
}

// NewYahooFinance create new YahooFinance
func NewYahooFinance() YahooFinance {
	return &yahooFinance{&http.Client{}, "https://query1.finance.yahoo.com"}
}

// GetActualValue Return Actual Value of stock
func (y *yahooFinance) GetActualValue(name string) (value float64, err error) {
	r, err := y.Client.Get(y.baseURL + "/v8/finance/chart/" + name)
	if err != nil {
		return
	}
	defer func() {
		r.Body.Close()
	}()

	var yahooFinanceResponse yahooFinanceResponse
	err = json.NewDecoder(r.Body).Decode(&yahooFinanceResponse)
	if err != nil {
		return
	}

	return yahooFinanceResponse.Chart.Result[0].Meta.RegularMarketPrice, nil
}
