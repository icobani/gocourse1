/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 20/11/2016    
* Time      : 17:49
* Developer : ibrahimcobani
*
*******/
package main

import "fmt"

func main() {
	message := "Hello Worlds"
	var greeting *string = &message
	// message'in adresini de bozdu
	*greeting = "hi"
	fmt.Println(message, greeting, *greeting)
}
