package context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	maincontext := context.Background()
	contexta := context.WithValue(maincontext, "name", "maulana")
	contextb := context.WithValue(contexta, "sister", "naima")
	contextc := context.WithValue(contextb, "teacher", "suprap")
	contextd := context.WithValue(contexta, "age", "20")

	fmt.Println(contextc.Value("age"))
	fmt.Println(contextd)
}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)

		counter := 1

		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter += 1
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return destination
}

func TestCancelContext(t *testing.T) {
	fmt.Println("GoRoutine", runtime.NumGoroutine())

	parent := context.Background()
	// ctx, cancel := context.WithCancel(parent)
	// timeout, cancel := context.WithTimeout(parent, 5*time.Second)
	deadline, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	destination := CreateCounter(deadline)

	for n := range destination {
		fmt.Println(n)

		// if n == 10 {
		// 	break
		// }
	}

	// TODO: Adding with waitgroup
	time.Sleep(1 * time.Second)

	fmt.Println("GoRoutine", runtime.NumGoroutine())
}
