package client

import (
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"

	"testing"
)

func TestGetStockValue(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{"chart":{"result":[{"meta":{"currency":"BRL","symbol":"GRND3.SA","exchangeName":"SAO","instrumentType":"EQUITY","firstTradeDate":1099053900,"regularMarketTime":1571688360,"gmtoffset":-10800,"timezone":"BRST","exchangeTimezoneName":"America/Sao_Paulo","regularMarketPrice":9.27,"chartPreviousClose":9.27,"previousClose":9.27,"scale":3,"priceHint":2,"currentTradingPeriod":{"pre":{"timezone":"BRST","start":1571658300,"end":1571659200,"gmtoffset":-7200},"regular":{"timezone":"BRST","start":1571659200,"end":1571684400,"gmtoffset":-7200},"post":{"timezone":"BRST","start":1571684400,"end":1571688000,"gmtoffset":-7200}},"tradingPeriods":[[{"timezone":"BRST","start":1571659200,"end":1571684400,"gmtoffset":-7200}]],"dataGranularity":"1m","range":"1d","validRanges":["1d","5d","1mo","3mo","6mo","1y","2y","5y","10y","ytd","max"]},"timestamp":[],"indicators":{"quote":[{"open":[],"low":[],"close":[],"volume":[],"high":[]}]}}],"error":null}}`))
	}))
	defer server.Close()

	yahooFinance := yahooFinance{server.Client(), server.URL}

	assert := assert.New(t)
	value, err := yahooFinance.GetActualValue("GRND3.SA")
	assert.NoError(err)
	assert.Equal(9.27, value)
}
