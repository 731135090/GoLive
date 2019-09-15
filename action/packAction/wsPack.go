package packAction

import (
	"GoLive/uitl/json"
	"GoLive/uitl/timer"
	"github.com/gorilla/websocket"
)

type WsPack struct {
	conn    *websocket.Conn
	form    string	//pack form
	to      string  //pack to
	action  string  //pack type
	mesType string  //message type
	message string  //message
	time    string  //server time
}

func NewWsPack(conn *websocket.Conn) *WsPack {
	pack := new(WsPack)
	pack.conn = conn
	pack.time = timer.GetNowDate()
	return pack
}

func (p *WsPack) Parse(jsonStr string) {
	json.Unmarshal([]byte(jsonStr))
}
