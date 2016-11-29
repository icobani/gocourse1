/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 20/11/2016    
* Time      : 13:44
* Developer : ibrahimcobani
*
*******/
package main

import "fmt"
func main() {
	//	var yerine := kullanılabiliyor.
	message := "Hello Worlds"

	// tip başına * o değişkenin pointer'in ifade ediyor.
	// değişkenin pointer değerini almak için & işareti kullanılıyor.
	var greeting *string = &message


	// değerin kendisi, değerin pointer adresi, pointer adresin değeri
	fmt.Println(message, greeting, *greeting)

}
