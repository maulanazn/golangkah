package bufferpool

import (
	"flag"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  256,
	WriteBufferSize: 256,
	WriteBufferPool: &sync.Pool{},
}

func process(c *websocket.Conn) {
	defer c.Close()
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Print("read:", err)
			break
		}

		log.Printf("Receive server: %s", message)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade: ", err)
		return
	}

	go process(c)
}

func TestRunServer(t *testing.T) {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", handler)

	server := http.Server{
		Addr:              "localhost:3000",
		ReadHeaderTimeout: 10 * time.Minute,
	}

	log.Fatal(server.ListenAndServe())
}
