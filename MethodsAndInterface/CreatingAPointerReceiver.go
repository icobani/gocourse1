/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 06/12/2016    
* Time      : 20:16
* Developer : ibrahimcobani
*
*******/
package MethodsAndInterface

import "fmt"

type Programmer struct {
	No      string
	Name    string
	SurName string
	Class   string
	Gender  string
}

func (programmer Programmer) String() string {
	return fmt.Sprintf("%v numaralı %v %v %v sınıfında okuyor", programmer.No, programmer.Name, programmer.SurName, programmer.Class)
}

func (programmer Programmer) Insert() {
	fmt.Println("Merhaba " + programmer.Name)
}
func (programmer Programmer) Modify() {
	fmt.Println("Merhaba " + programmer.Name)
}
func (programmer Programmer) Delete() {
	fmt.Println("Merhaba " + programmer.Name)
}
func (programmer *Programmer) FindFirst() bool {
	programmer.Name = "ibrahim"
	programmer.SurName = "COBANI"
	programmer.Gender = "Male"
	programmer.Class = "7-A"
	return true
}
