/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 29/11/2016    
* Time      : 23:00
* Developer : ibrahimcobani
*
*******/
package variadicFunction

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " | ",len(nums), " => eleman, Toplam : ")
	total := 0
	for _, num := range nums {
		total += num
	}
	println(total)
}



func main() {
	sum(5,4,3,2,5)
}

