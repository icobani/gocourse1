/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 24/12/2016    
* Time      : 10:43
* Developer : ibrahimcobani
*
*******/
package ConcurencySamples

import "fmt"

func SwitchingBwtweenChannelsMain() {
	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)
/*
	msg := Message{
		To: []string{"ibrahim@imaconsult.com"},
		From: "destek@imaconsult.com",
		Content :"Bu sırrı koru, kendine iyi bak",
	}

	failedMessage := FailedMessage{
		ErrorMessage:"Bu mesaj ici bey tarafından yakalandı.",
		OriginalMessage:Message{},
	}

	msgCh <- msg
	errCh <- failedMessage
*/
	select {
	case receivedMsg := <-msgCh:
		fmt.Println(receivedMsg)
	case receivedErr := <-errCh:
		fmt.Println(receivedErr)
	default:
		fmt.Println("No message received")
	}

}

type Message struct {
	To      []string
	From    string
	Content string
}

type FailedMessage struct {
	ErrorMessage    string
	OriginalMessage Message
}