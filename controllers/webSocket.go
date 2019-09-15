package controllers

import (
	"GoLive/action/packAction"
	"GoLive/config"
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
)

type WebSocketController struct {
	BaseController
}

func (c *WebSocketController) Get() {
	header := http.Header{}
	header.Add("Server", "tomcat")

	conn, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, header)
	if err != nil {
		config.Logger.Error(err.Error())
		return
	}
	defer conn.Close()

	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			config.Logger.Error(err.Error())
			continue
		}
		config.Logger.Info(string(data))
		packAction.WsPackChannel <- packAction.NewWsPack(conn, packAction.WS_PACK_TYPE_CHEAT, data)
	}
	return
}
