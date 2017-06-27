package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// flagを使いオプションでアドレスを指定できるようにする
var addr = flag.String("addr", "localhost:8080", "http service address")
var upgrader = websocket.Upgrader{}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("upgrade:", err)
		return
	}
	defer c.Close()

	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			log.Fatal("read:", err)
			break
		}
		log.Printf("recv: %s", msg)
		if err := c.WriteMessage(mt, msg); err != nil {
			log.Fatal("write: %s", err)
			break
		}
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)

	http.HandleFunc("/echo", echo)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
