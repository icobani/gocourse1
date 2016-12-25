/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 25/12/2016    
* Time      : 14:13
* Developer : ibrahimcobani
*
*******/
package IntegrationGoRoutinesAndChannels

import (
	"fmt"
	"errors"
)

func SimulatingPromisesMain() {
	so := new(SalesOrder)
	so.Amount = 472.39

	SaveSO(so, true).Then(
		// Eğer hata olmadıysa buraya gelir.
		func(obj interface{}) error {
			so := obj.(*SalesOrder)
			fmt.Printf("Sales order saved with ID : %s\n", so.No)
			return nil
		},
		// Eğer hata olduysa buraya gelir.
		func(err error) {
			fmt.Printf("Faild to save Sales Order : %s\n", err.Error())
		},
	)

	fmt.Scanln()
}

// Satış siparişi
type SalesOrder struct {
	No     string
	Amount float64
}

//
//@so : Satış Siparişi
//@shouldFail : Sadece similasyon için true gönderilirse sanki hata olmuş gibi yapıyor.
func SaveSO(so *SalesOrder, shouldFail bool) *Promise {
	//Söz verilen Promicess objesi oluşturuluyor.
	result := new(Promise)
	result.sucsesChannel = make(chan interface{}, 1)
	result.failureChannel = make(chan error, 1)

	go func() {
		if shouldFail {
			// Bir hata alınması durumunda
			result.failureChannel <- errors.New("Satış siparişini kaydederken bir sorun çıktı")
		} else {
			// Bir hata alınmaması durumunda oluşan siparişin yeni numarası geri döndürülüyor.
			so.No = "SS2016-0003"
			result.sucsesChannel <- so
		}
	}()
	return result
}

//Bu struct bir anlamda bir durumda kesinlikle buraların çağrılacağı anlamına geliyor.
type Promise struct {
	sucsesChannel  chan interface{}
	failureChannel chan error
}

func (this *Promise) Then(
success func(interface{}) error,
failler func(error),
) *Promise {

	result := new(Promise)
	result.sucsesChannel = make(chan interface{}, 1)
	result.failureChannel = make(chan error, 1)

	go func() {
		select {
		case obj := <-this.sucsesChannel:
			newErr := success(obj)
			if newErr == nil {
				result.sucsesChannel <- obj
			} else {
				result.failureChannel <- newErr
			}
		case err := <-this.failureChannel:
			failler(err)
			result.failureChannel <- err
		}
	}()
	return result
}