/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 05/12/2016    
* Time      : 21:21
* Developer : ibrahimcobani
*
*******/
package Maps

import (
	"github.com/icobani/course1/LoopsSamples"
	"fmt"
)

func BasicMap(studentNumber string) {
	var studentMap map[string]LoopsSamples.Student
	studentMap = make(map[string]LoopsSamples.Student)

	studentMap["5462"] = LoopsSamples.Student{No: "5462", Name: "Alisinan", SurName: "Cobani", Class: "7-B", Gender: "Male"}
	studentMap["3549"] = LoopsSamples.Student{No: "3549", Name: "Serhat", SurName: "Deli", Class: "7-B", Gender: "Male"}
	studentMap["2543"] = LoopsSamples.Student{No: "2543", Name: "Selin", SurName: "Deli", Class: "7-A", Gender: "Female"}

	fmt.Println(studentMap[studentNumber])
}
