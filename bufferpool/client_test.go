package bufferpool

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func runNewConnection(wg *sync.WaitGroup) {
	defer wg.Done()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	log.Printf("Connecting to %s", "ws://localhost:3000/echo")
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:3000/echo", nil)
	if err != nil {
		log.Fatal("Dial", err)
		return
	}

	defer conn.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Fatal("Read", err)
				return
			}

			log.Printf("Receive %s", message)
		}
	}()

	ticker := time.NewTicker((time.Second * 1))
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
    case t := <-ticker.C:
      text := fmt.Sprintf("Hello world in %d", t.Second())
			err := conn.WriteMessage(websocket.TextMessage, []byte(text))
			if err != nil {
				log.Println("Write error : ", err)
				return
			}
		case <-interrupt:
			log.Println("Interrupt")

			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close: ", err)
				return
			}

			select {
			case <-done:
			case <-time.After(time.Second):
			}

			return
		}
	}
}

func TestRunClient(t *testing.T) {
	flag.Parse()
	log.SetFlags(0)

	wg := &sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go runNewConnection(wg)
	}

	wg.Wait()
}
