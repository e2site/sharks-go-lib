package server

import (
	"github.com/e2site/sharks-go-lib/log"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
	"sync"
)

var upgrader = websocket.Upgrader{}

var clients = make(map[string]*websocket.Conn)
var mu sync.Mutex

func SendWsMessage(telegramId int64, message string) {
	mu.Lock()
	defer mu.Unlock()
	keyTg := strconv.FormatInt(telegramId, 10)
	if ws, ok := clients[keyTg]; ok {
		err := ws.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			delete(clients, keyTg)
			ws.Close()
		}
	}
}

func CreateSocketServer(handlerReader func(message string), httpHeaderName string) {
	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		telegramId := request.Header.Get(httpHeaderName)
		if telegramId == "" {
			http.Error(writer, "Bad token", 400)
			return
		}

		ws, err := upgrader.Upgrade(writer, request, nil)
		defer ws.Close()
		if err != nil {
			log.Log(err)
			return
		}

		_, exist := clients[telegramId]

		if exist {
			http.Error(writer, "Connection exist", 421)
			return
		}

		mu.Lock()
		clients[telegramId] = ws
		mu.Unlock()

		for {
			_, msg, err := ws.ReadMessage()
			if err != nil {
				break
			}
			handlerReader(string(msg))
		}
		mu.Lock()
		delete(clients, telegramId)
		mu.Unlock()
	})
	err := http.ListenAndServe(":8080", nil)
	log.CheckEndLogFatal(err)
}
