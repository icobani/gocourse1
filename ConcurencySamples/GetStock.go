/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 23/12/2016    
* Time      : 21:57
* Developer : ibrahimcobani
*
*******/
package ConcurencySamples

import (
	"net/http"
	"io/ioutil"
	"encoding/xml"
	"fmt"
	"time"
	"runtime"
)

func GetStockValue() {
	runtime.GOMAXPROCS(4)

	start := time.Now()


	stockSymbol := []string{
		"googl",
		"msft",
		"aapl",
		"bbry",
		"hpq",
		"vz",
		"t",
		"tmus",
		"s",
	}

	numcomplete := 0
	for _, symbol := range stockSymbol {
		go func(symbol string) {
			resp, _ := http.Get("http://dev.markitondemand.com/MODApis/Api/v2/quote?symbol=" + symbol)
			defer resp.Body.Close()
			data, _ := ioutil.ReadAll(resp.Body)

			quote := new(QuoteResponce)
			xml.Unmarshal(data, &quote)

			fmt.Printf("%s: %.2f\n", quote.Name, quote.LastPrice)
			numcomplete++

		}(symbol)
	}
	for numcomplete < len(stockSymbol){
		time.Sleep(10 * time.Millisecond)
	}


	elapsed := time.Since(start)
	fmt.Printf("Execution Time : %s", elapsed)
}

type QuoteResponce struct {
	Status           string
	Name             string
	LastPrice        float32
	Change           float32
	ChangePercent    float32
	TimeStamp        string
	MSDate           string
	MarketCap        int
	Volume           int
	ChangeYDT        float32
	ChangePercentYTD float32
	High             float32
	Low              float32
	Open             float32
}
