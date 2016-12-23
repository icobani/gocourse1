/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 05/12/2016    
* Time      : 20:47
* Developer : ibrahimcobani
*
*******/
package LoopsSamples

import "fmt"

type Student struct {
	No      string
	Name    string
	SurName string
	Class   string
	Gender  string
}

func (student Student) String() string {
	return fmt.Sprintf("%v numaralı %v %v %v sınıfında okuyor", student.No, student.Name, student.SurName, student.Class)
}

func BasicRangeSample() {
	students := []Student{
		{"5462", "Alisinan", "Cobani", "7-B", "Male"},
		{"5748", "Ata", "Ravalı", "8-B", "Male"},
		{"3934", "Yiğit", "Kaplan", "7-B", "Male"}}

	for _, student := range students {
		if (student.Class == "8-B") {
			fmt.Println(student.No, student.Name, student.Class)
		}
	}

}