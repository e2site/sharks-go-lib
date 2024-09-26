package server

import (
	"github.com/btcsuite/btcutil/base58"
	"github.com/e2site/sharks-go-lib/log"
	"github.com/e2site/sharks-go-lib/telegram"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const pingPeriod = 54 * time.Second

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type StatusConnection struct {
	Connection map[string]*websocket.Conn
}

func (s *StatusConnection) NewClient(ws *websocket.Conn) (string, bool) {
	if s.Connection == nil {
		s.Connection = make(map[string]*websocket.Conn)
	}
	if len(s.Connection) > 9 {
		return "", false
	}

	index := time.Now().String()
	s.Connection[index] = ws
	return index, true
}
func (s *StatusConnection) CountClient() int {
	return len(s.Connection)
}

func (s *StatusConnection) DelClient(ind string) {
	_, exist := s.Connection[ind]
	if exist {
		delete(s.Connection, ind)
	}
}
func (s *StatusConnection) SendMessage(message string) {
	for ind, ws := range s.Connection {
		err := ws.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			ws.Close()
			s.DelClient(ind)
		}
	}
}

var clients = make(map[string]*StatusConnection)
var mu sync.Mutex

func SendWsMessage(telegramId int64, message string) {
	mu.Lock()
	defer mu.Unlock()
	keyTg := strconv.FormatInt(telegramId, 10)
	if sc, ok := clients[keyTg]; ok {
		sc.SendMessage(message)
		if sc.CountClient() == 0 {
			delete(clients, keyTg)
		}
	}
}

func ping(ws *websocket.Conn, telegramId string, idx string, done chan struct{}) {
	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if err := ws.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(pingPeriod)); err != nil {
				_, exist := clients[telegramId]
				if exist {
					clients[telegramId].DelClient(idx)
				}
				if clients[telegramId].CountClient() == 0 {
					delete(clients, telegramId)
				}
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
		decode := base58.Decode(authStr)

		tgAuth := telegram.NewTelegramAuth(string(decode), tokenTg)
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

		mu.Lock()
		if !exist {
			clients[telegramId] = &StatusConnection{}
		}
		idx, status := clients[telegramId].NewClient(ws)
		mu.Unlock()

		if !status {
			http.Error(writer, "Max limit connection", 402)
			return
		}

		stdoutDone := make(chan struct{})
		defer func() {
			close(stdoutDone)
		}()

		go ping(ws, telegramId, idx, stdoutDone)

		for {
			_, msg, err := ws.ReadMessage()
			if err != nil {
				break
			}
			handlerReader(string(msg))
		}

		mu.Lock()
		clients[telegramId].DelClient(idx)
		if clients[telegramId].CountClient() == 0 {
			delete(clients, telegramId)
		}
		mu.Unlock()
	})
	err := http.ListenAndServe(":8080", nil)
	log.CheckEndLogFatal(err)
}
