/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 04/12/2016    
* Time      : 09:56
* Developer : ibrahimcobani
*
*******/
package LoopsSamples

import "fmt"

func WhileLoopSample() {
	fmt.Println("----------")
	fmt.Println("While Loop Sample")
	fmt.Println("----------")
	var times int = 3
	var i int = 0
	for i < times {
		i++
		fmt.Print(i)
		if i % 2 == 1 {
			fmt.Println(" is single")
		} else {
			fmt.Println(" is double")
		}

	}
	fmt.Println("----------")
}
