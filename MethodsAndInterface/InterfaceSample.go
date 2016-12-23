/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 06/12/2016    
* Time      : 20:34
* Developer : ibrahimcobani
*
*******/
package MethodsAndInterface

import "fmt"

type Consultant struct {
	No      string
	Name    string
	SurName string
	Class   string
	Gender  string
}

func (consultant Consultant) String() string {
	return fmt.Sprintf("%v numaralı %v %v %v sınıfında okuyor",consultant.No,consultant.Name,consultant.SurName,consultant.Class)
}

func (consultant Consultant) Insert() {
	fmt.Println("Merhaba " + consultant.Name)
}

func (consultant Consultant) Modify() {
	fmt.Println("Merhaba " + consultant.Name)
}

func (consultant Consultant) Delete() {
	fmt.Println("Merhaba " + consultant.Name)
}

func (consultant Consultant) Get() {
	fmt.Println("Merhaba " + consultant.Name)
}

func (consultant *Consultant) FindFirst() bool {
	consultant.Name = "ibrahim"
	consultant.SurName = "COBANI"
	consultant.Gender = "Male"
	consultant.Class = "7-A"
	return true
}