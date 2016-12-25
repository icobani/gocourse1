/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 24/12/2016    
* Time      : 10:18
* Developer : ibrahimcobani
*
*******/
package ConcurencySamples

import (
	"fmt"
	"strings"
)

func BufferedChannelsMain() {
	phrase := "Kabak kafalı keriz kralın kabak kafalı keriz kızının kara kıllı kaniş köpeği\n"
	words := strings.Split(phrase, " ")

	ch := make(chan string, len(words))

	for _, word := range words {
		ch <- word
	}
	close(ch)
	for msg := range ch {
		fmt.Print(msg + " ")
	}
}