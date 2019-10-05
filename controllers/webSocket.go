package controllers

import (
	"GoLive/action/connectAction"
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

type AdminWsController struct {
	BaseController
}

func (c *WebSocketController) Get() {
	access(&c.BaseController, config.WS_CONN_TYPE_USER)
}
func (c *AdminWsController) Get() {
	access(&c.BaseController, config.WS_CONN_TYPE_CUSTOMER)
}
func access(c *BaseController, connType uint8) {
	header := http.Header{}
	header.Add("Server", "tomcat")

	conn, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, header)
	if err != nil {
		config.Logger.Error(err.Error())
		return
	}
	defer conn.Close()
	defer connectAction.CloseWsConn(conn, connType, "")

	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			config.Logger.Error(err.Error())
			continue
		}
		config.Logger.Info(string(data))
		packAction.WsPackChannel <- packAction.NewWsPack(conn, connType, data)
	}
	return
}
