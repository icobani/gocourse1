/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 11/12/2016    
* Time      : 19:04
* Developer : ibrahim COBANI
*
*******/
package main

import "time"

func main() {
	godur, _ := time.ParseDuration("1ms")
	go func() {
		for i := 0; i < 100; i++ {
			println(i + 1, "Hello")
			time.Sleep(godur)
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			println(i + 1, "Go")
			time.Sleep(godur)
		}
	}()

	dur, _ := time.ParseDuration("120ms")
	time.Sleep(dur)
}
