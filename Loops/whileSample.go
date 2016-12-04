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
package Loops
import "fmt"

func whileLoopSample() {
	var times int = 3
	var i int = 0
	for i < times {
		fmt.Print(i)
		fmt.Println(" times hello")
	}
}
