package liveupdate

import (
	"fmt"
	"live-server/logger"
	"live-server/watcher"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Handler(path string, logEnabled bool) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			logger.Log(true, err.Error())
			return
		}
		defer conn.Close()

		dispose := watcher.Subscribe(path, func(event watcher.Event) {
			logger.Log(logEnabled, event.String())
			conn.WriteMessage(1, []byte("FILES_CHANGED"))
		})
		defer dispose()

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				logger.Log(true, err.Error())
				break
			}
			logger.Log(logEnabled, fmt.Sprintf("Client message: %s", message))
		}
	}
}
