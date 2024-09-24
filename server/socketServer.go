package server

import (
	"github.com/e2site/sharks-go-lib/log"
	"github.com/e2site/sharks-go-lib/telegram"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const pingPeriod = 54 * time.Second

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

func ping(ws *websocket.Conn, telegramId string, done chan struct{}) {
	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if err := ws.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(pingPeriod)); err != nil {
				delete(clients, telegramId)
				ws.Close()
			}
		case <-done:
			return
		}
	}
}

func CreateSocketServer(handlerReader func(message string), httpHeaderName string, tokenTg string) {
	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		queryParam := request.URL.Query()
		authStr := queryParam.Get("auth")

		if authStr == "" {
			http.Error(writer, "Bad token", 400)
			return
		}
		tgAuth := telegram.NewTelegramAuth(authStr, tokenTg)
		if !tgAuth.CheckAuth() {
			http.Error(writer, "Bad token", 400)
			return
		}

		telegramId := strconv.Itoa(tgAuth.ID)

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

		stdoutDone := make(chan struct{})
		defer func() {
			close(stdoutDone)
		}()

		go ping(ws, telegramId, stdoutDone)

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
