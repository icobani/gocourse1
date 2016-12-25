/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 25/12/2016    
* Time      : 13:50
* Developer : ibrahimcobani
*
*******/
package IntegrationGoRoutinesAndChannels

import (
	"fmt"
)

func SimulatingCallbackMain() {

	po := new(PurchaseOrder)
	po.Amount = 450.99

	ch := make(chan *PurchaseOrder)
	go SavePurchaseOrder(po, ch)

	newPo := <- ch
	fmt.Printf("PO : %v", newPo)
}

type PurchaseOrder struct {
	No     string
	Amount float64
}

func SavePurchaseOrder(po *PurchaseOrder, callback chan *PurchaseOrder) {
	po.No = "SS2016-0001"
	callback <- po
}