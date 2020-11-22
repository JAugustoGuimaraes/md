package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"

	"github.com/christiansakai/md/lib/mdreader"
	"github.com/christiansakai/md/lib/mdrender"
	"github.com/christiansakai/md/lib/watcher"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(`Usage:
go run <mdfile>`,
		)
		return
	}

	filename := os.Args[1]

	htmlStr, err := mdreader.ReadMDFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	renderer, err := mdrender.New(filename, string(htmlStr))
	if err != nil {
		log.Fatal(err)
	}

	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	var wsConn *websocket.Conn

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatal(err)
		}

		wsConn = conn
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderer.Render(w)
	})

	watcher.WatchFile(filename, func() {
		htmlBytes, err := mdreader.ReadMDFile(filename)
		if err != nil {
			log.Fatal(err)
		}

		if wsConn != nil {
			err := wsConn.WriteMessage(websocket.TextMessage, htmlBytes)
			if err != nil {
				log.Fatal(err)
			}
		}
	})

	fmt.Println("Visit localhost:8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalln(err)
	}
}
