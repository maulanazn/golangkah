package routine

import (
	"fmt"
	"testing"
	"time"
)

func SendData(channel chan<- string) {
	channel <- "maulana"
}

func ReceiveData(channel <-chan string) {
	data := <-channel

	fmt.Println(data)
}

func TestInOut(t *testing.T) {
	channel := make(chan string)

	defer close(channel)

	go SendData(channel)
	go ReceiveData(channel)

	time.Sleep(1 * time.Second)
}
