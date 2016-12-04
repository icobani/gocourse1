/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 04/12/2016    
* Time      : 12:43
* Developer : ibrahimcobani
*
*******/
package LoopsSamples

import "fmt"

// Sonsuz döngü
func InfiniteLoop() {
	var times int = 3
	var i int = 0
	for {
		i++
		if i >= times {
			break
		}
		fmt.Println("Heloo")

	}
}
