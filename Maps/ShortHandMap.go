/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 05/12/2016    
* Time      : 21:38
* Developer : ibrahimcobani
*
*******/
package Maps

import (
	"github.com/icobani/course1/LoopsSamples"
	"fmt"
)

func ShortHandMap() {
	students := map[string]LoopsSamples.Student{
		"123": LoopsSamples.Student{No:"123", Name:"ibrahim", SurName:"ÇOBANİ", Class:"4-A", Gender:"Male"},
		"333": LoopsSamples.Student{No:"333", Name:"Seyhan", SurName:"ÇOBANİ", Class:"4-A", Gender:"Female"},
	}

	// Update Sample
	students["333"] = LoopsSamples.Student{No:"333", Name:"Seyhan", SurName:"ÇOBANİ", Class:"4-D", Gender:"Female"}
	// Insert Sample
	students["335"] = LoopsSamples.Student{No:"335", Name:"Alisinan", SurName:"ÇOBANİ", Class:"4-D", Gender:"Female"}
	students["888"] = LoopsSamples.Student{No:"888", Name:"For Delete", SurName:"ÇOBANİ", Class:"4-D", Gender:"Female"}

	fmt.Println(" ")
	fmt.Println("---------------------------")
	fmt.Println("Shorthand Map")
	fmt.Println("---------------------------")
	fmt.Println(students["333"])
	fmt.Println(students["335"])
	fmt.Println(students["888"])
	delete(students, "888")

	if value, exist := students["888"]; exist {
		fmt.Println(value)
	} else {
		fmt.Println("Not found")
	}
	fmt.Println(" ")
}
