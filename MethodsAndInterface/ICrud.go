/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 06/12/2016    
* Time      : 20:49
* Developer : ibrahimcobani
*
*******/
package MethodsAndInterface

type ICrud interface {
	Insert()
	Modify()
	Delete()
	Get()
	FindFirst() bool
}
