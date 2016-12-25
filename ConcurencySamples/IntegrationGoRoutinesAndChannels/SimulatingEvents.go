/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 25/12/2016    
* Time      : 10:45
* Developer : ibrahimcobani
*
*******/
package IntegrationGoRoutinesAndChannels

import (
	"fmt"
)

func SimulatingEventsMain() {
	btn := MakeButton()

	handlerOne := make(chan string)
	handlerTwo := make(chan string)

	btn.AddEventListener("click", handlerOne)
	btn.AddEventListener("click", handlerTwo)

	go func() {
		for {
			msg := <- handlerOne
			fmt.Println("Hadler One :" + msg)
		}
	}()

	go func() {
		for {
			msg := <- handlerTwo
			fmt.Println("Hadler Two :" + msg)
		}
	}()

	btn.TriggerEvent("click","Button clicked")
	btn.RemoveEventListener("click", handlerTwo)
	btn.TriggerEvent("click", "Button clicked again amanın düüüt güüüüm")

	fmt.Scanln()
}

type Button struct {
	eventListener map[string][]chan string
}

//Constrator Function for Button
func MakeButton() *Button {
	result := new(Button)
	result.eventListener = make(map[string][]chan string)
	return result
}

//Add event listener to Button
func (this *Button)AddEventListener(event string, responceChannel chan string) {
	//array içinde arama yapılıyor eğer bulunursa append edilecek bulunamazsa eklenecek.
	if _, present := this.eventListener[event]; present {
		this.eventListener[event] = append(this.eventListener[event], responceChannel)
	} else {
		this.eventListener[event] = []chan string{responceChannel}
	}
}

func (this *Button)RemoveEventListener(event string, listenerChannel chan string) {

	if _, present := this.eventListener[event]; present {
		for idx, _ := range this.eventListener[event] {
			if this.eventListener[event][idx] == listenerChannel {
				this.eventListener[event] = append(this.eventListener[event][:idx],
					this.eventListener[event][idx + 1:]...)
				break
			}
		}
	}
}


func (this *Button)TriggerEvent(event string, responce string){
	if _, presend := this.eventListener[event]; presend{
		for _, handler := range this.eventListener[event]{
			go func(handler chan string) {
				handler <- responce
			}(handler)
		}
	}
}