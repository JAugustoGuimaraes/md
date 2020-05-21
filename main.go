package main

import (
  "log"
  "net/http"

  "github.com/gorilla/websocket"

  "github.com/christiansakai/md/lib/mdreader"
  "github.com/christiansakai/md/lib/mdrender"
  "github.com/christiansakai/md/lib/watcher"
)

func main() {
  filename := "README.md"

  htmlStr, err := mdreader.ReadMDFile(filename)
  if (err != nil) {
    log.Fatal(err)
  }

  renderer, err := mdrender.New()
  if (err != nil) {
    log.Fatal(err)
  }

  renderer.SetTitle(filename)
  renderer.SetContent(htmlStr)

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
    htmlStr, err := mdreader.ReadMDFile(filename)
    if (err != nil) {
      log.Fatal(err)
    }

    renderer.SetContent(htmlStr)

    if wsConn != nil {
      err := wsConn.WriteMessage(websocket.TextMessage, []byte("refresh"))
      if err != nil {
        log.Print(err)
      }

      wsConn.Close()
    }
  })

  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatalln(err.Error())
  }
}
