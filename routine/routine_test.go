package routine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func RunHelloWorld(i int) {
	fmt.Println("Hello world number : ", i)
}

func RunWorld(data chan string) {
	time.Sleep(1 * time.Second)
	fmt.Println(data)
}

func TestRunRoutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go RunHelloWorld(i)
	}

	time.Sleep(3 * time.Second)
}

func BenchmarkPerformHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go RunHelloWorld(i)
	}

	time.Sleep(2 * time.Second)
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string, 3)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Sending data number " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data ", data)
	}
}

func TestSelectChannel(t *testing.T) {
	channel := make(chan string, 3)
	channel1 := make(chan string, 3)

	defer close(channel)
	defer close(channel1)

	go RunWorld(channel)
	go RunWorld(channel1)

	counter := 0
	for {
		select {
		case data := <-channel:
			fmt.Println("Data channel 1", data)
			counter++
		case data := <-channel1:
			fmt.Println("Data channel 2", data)
			counter++
		default:
			fmt.Println("Waiting for data")
		}

		if counter == 2 {
			break
		}
	}
}
