/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 06/12/2016    
* Time      : 19:53
* Developer : ibrahimcobani
*
*******/
package MethodsAndInterface

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

func (student Student) Insert() {
	fmt.Println("Merhaba " + student.Name)
}
func (student Student) Modify() {
	fmt.Println("Merhaba " + student.Name)
}
func (student Student) Delete() {
	fmt.Println("Merhaba " + student.Name)
}
func (student Student) FindFirst() (bool, Student) {
	student.Name = "ibrahim"
	student.SurName = "COBANI"
	student.Gender = "Male"
	student.Class = "7-A"
	return true, student
}