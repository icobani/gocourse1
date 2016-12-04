/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 29/11/2016    
* Time      : 23:32
* Developer : ibrahimcobani
*
*******/
package passingFunction

import "fmt"

func sum(do func(int), nums ...int) {
	fmt.Print(nums, " | ",len(nums), " => eleman, Toplam : ")
	total := 0
	for _, num := range nums {
		total += num
	}
	do(total)
}

func PrinL(val int){
	fmt.Println(val)
	fmt.Println("******")
}

func main() {
	sum(PrinL,5,4,3,2,5)
}
