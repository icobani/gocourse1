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
	var times int = 1000
	var i int = 2
	for {
		i++
		if i >= times {
			break
		}

		var noPrimeNumber bool  = false

		for t := 2; t < i; t++ {
			if (i % t == 0) {
				noPrimeNumber = true
				break
			}
		}

		if noPrimeNumber {
			continue
		}


		fmt.Printf(" %v is prime number\n", i)

	}
}
