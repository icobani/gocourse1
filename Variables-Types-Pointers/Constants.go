/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 20/11/2016    
* Time      : 20:25
* Developer : ibrahimcobani
*
*******/
package Constants

const (
	PI = 3.14
	Language = "Go"
)

//enum gibi
const (
	a = iota
	b
	c
	d
)

func main() {
	println(PI, Language, a, b, c, d)
}
