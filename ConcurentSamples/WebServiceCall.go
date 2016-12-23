/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 11/12/2016    
* Time      : 19:49
* Developer : ibrahimcobani
*
*******/
package main

import (
	"net/http"
	"io/ioutil"
	"encoding/xml"
	"fmt"
)

func main() {
	dt := new(Tarih_Date)
	dt.GetToday("2016", "12", "09")
}

func (c *Tarih_Date) GetToday(year string, month string, day string) {
	resp, _ := http.Get("http://www.tcmb.gov.tr/kurlar/" + year + month + "/" + day + month + year +
	".xml")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	tarih := new(Tarih_Date)

	xml.Unmarshal(body, &tarih)
	c = &Tarih_Date{}
	fmt.Println(tarih.Date, tarih.Bulten_No)
	fmt.Println("------------------")
	for _, curr := range tarih.Currency {
		fmt.Println(curr.Kod, curr.Isim, curr.ForexBuying)
	}

}

type Tarih_Date struct {
	XMLName   xml.Name `xml:"Tarih_Date"`
	Tarih     string `xml:"Tarih,attr"`
	Date      string `xml:"Date,attr"`
	Bulten_No string `xml:"Bulten_No,attr"`
	Currency  []Currency `xml:"Currency"`
}

type Currency struct {
	Kod             string `xml:"Kod,attr"`
	CrossOrder      int `xml:"CrossOrder,attr"`
	CurrencyCode    string `xml:"CurrencyCode,attr"`
	Unit            int `xml:"Unit"`
	Isim            string `xml:"Isim"`
	CurrencyName    string `xml:"CurrencyName"`
	ForexBuying     string `xml:"ForexBuying"`
	BanknoteBuying  string `xml:"BanknoteBuying"`
	BanknoteSelling string `xml:"BanknoteSelling"`
	CrossRateUSD    string `xml:"CrossRateUSD"`
	CrossRateOther  string `xml:"CrossRateOther"`
}