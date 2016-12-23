/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 20/11/2016    
* Time      : 13:42
* Developer : ibrahimcobani
*
*******/
package main

import (
	"github.com/icobani/GOTCMBCurrencyHelper"
	"time"
)

func main() {

	CurrencyDate,_ := time.Parse("02-01-2006","07-08-2016")

	dt := new(GOTCMBCurrencyHelper.Tarih_Date)
	dt.GetToday(CurrencyDate)
}