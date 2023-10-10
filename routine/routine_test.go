package routine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld(i int) {
	fmt.Println("Hello world number : ", i)
}

func TestRunRoutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go RunHelloWorld(i)
	}

	time.Sleep(3 * time.Second)
}

func BenchmarkPerformHelloWorld(b *testing.B) {
  for i := 0; i < b.N; i++ {
   go RunHelloWorld(i); 
  }

  time.Sleep(2 * time.Second)
}

func TestDoSomething(t *testing.T) {
  channel := make(chan string)
  defer close(channel)

  go func(c chan string) {
    c <- "maulana"
  }(channel)

  data := <- channel

  fmt.Println(data)
}
