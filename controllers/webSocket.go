package controllers

import (
	"GoLive/config"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{}
)

type WebSocketController struct {
	BaseController
}

func (c *WebSocketController) Get() {
	conn, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if err != nil {
		config.Logger.Error(err.Error())
		return
	}
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			config.Logger.Error(err.Error())
			continue
		}
		config.Logger.Info(string(msg))
		err = conn.WriteMessage(websocket.TextMessage, []byte("hello client"))
		if err != nil{
			config.Logger.Error(err.Error())
		}
	}
	return
}
