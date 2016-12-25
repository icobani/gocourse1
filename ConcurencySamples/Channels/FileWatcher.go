/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 24/12/2016    
* Time      : 09:07
* Developer : ibrahimcobani
*
*******/
package ConcurencySamples

import (
	"fmt"
	"os"
	"time"
	"io/ioutil"
	"strings"
	"encoding/csv"
	"strconv"
	"runtime"
)

const watchedfolder = "./invoices"

func FileWatcher_main() {
	runtime.GOMAXPROCS(4)
	for {
		time.Sleep(10000 * time.Millisecond)
		d, _ := os.Open(watchedfolder)
		files, _ := d.Readdir(-1)
		for _, fl := range files {
			filepath := watchedfolder + "/" + fl.Name()
			f, _ := os.Open(filepath)
			data, _ := ioutil.ReadAll(f)
			f.Close()
			os.Remove(filepath)

			go func(data string) {
				reader := csv.NewReader(strings.NewReader(data))
				records, _ := reader.ReadAll()
				for _, r := range records {
					invoice := new(SalesLine)
					unixTime, _ := strconv.ParseInt(r[0], 10, 64)
					invoice.PostingDate = time.Unix(unixTime, 0)
					invoice.DocumentNo = r[1]
					invoice.LineNo, _ = strconv.Atoi(r[2])
					invoice.ItemNo = r[3]
					invoice.UnitPrice, _ = strconv.ParseFloat(r[4], 64)
					invoice.Quantity, _ = strconv.Atoi(r[5])
					invoice.LineAmount = invoice.UnitPrice * float64(invoice.Quantity)
					fmt.Printf("%s Received Invoice '%v' for $%.2f and Submited\n",invoice.PostingDate, invoice.DocumentNo, invoice.LineAmount)
				}
			}(string(data))
		}
	}
}

type SalesLine struct {
	PostingDate time.Time
	DocumentNo  string
	LineNo      int
	ItemNo      string
	UnitPrice   float64
	Quantity    int
	LineAmount  float64
}