/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 04/12/2016    
* Time      : 09:21
* Developer : ibrahimcobani
*
*******/

package Loops
import "fmt"
func SimpleSample() {
	var times int = 4
	for i := 0; i < times; i++ {
		fmt.Print(i + 1)
		fmt.Println(" times hello")
	}
}