/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 25/12/2016    
* Time      : 09:36
* Developer : ibrahimcobani
*
*******/
package IntegrationGoRoutinesAndChannels

import (
	"fmt"
	"runtime"
	"os"
	"time"
)

func IntegrationGoRoutinesAndChannelsMain() {
	runtime.GOMAXPROCS(4)

	f, _ := os.Create("./log.txt")
	f.Close()

	logCh := make(chan string, 50)

	go func() {
		for {
			msg, ok := <-logCh

			if ok {
				if f, err := os.OpenFile("./log.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend); err != nil {
					panic(err)
				} else {
					logTime := time.Now().Format(time.RFC3339)
					//fmt.Print("**** =>" + logTime + " - " + msg)
					if _, err := f.WriteString(logTime + " - " + msg); err != nil {
						fmt.Println(err.Error())
					}
					f.Close()
				}
			} else {
				fmt.Print("Channel closed! \n")
				break
			}
		}
	}()

	mutex := make(chan bool, 1)
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			mutex <- true
			go func() {
				msg := fmt.Sprintf("%d + %d = %d\n", i, j, i + j)
				logCh <- msg
				<-mutex
			}()
		}
	}
	fmt.Scanln()
}