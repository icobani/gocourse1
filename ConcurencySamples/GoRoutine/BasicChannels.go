/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 24/12/2016    
* Time      : 10:07
* Developer : ibrahimcobani
*
*******/
package ConcurencySamples

import (
	"fmt"
)

func BasicChannelMain() {
	ch := make(chan string, 1)
	ch <- "Hello World"

	fmt.Println(<-ch)
}
