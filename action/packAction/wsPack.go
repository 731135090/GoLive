package packAction

import (
	"GoLive/uitl/json"
	"GoLive/uitl/timer"
	"github.com/gorilla/websocket"
)

const (
	WS_PACK_TYPE_CHEAT    = iota //聊天消息
	WS_PACK_TYPE_SYSTEM          //系统消息
	WS_PACK_TYPE_CUSTOMER        //客服消息
)

const (
	WS_PACK_ACTION_PING = "ping"

	WS_MES_TYPE_TEXT = "text"
	WS_MES_TYPE_IMG  = "img"
)

type Message struct {
	Action  string `json:"action"` //pack type
	Msg     string `json:"msg"`    //message
	MsgType string `json:"type"`   //msg type
	Time    string `json:"time"`   //server time
}
type WsPack struct {
	conn     *websocket.Conn
	form     string //pack form
	to       string //pack to
	packType int    //pack type
	Message
}

func NewWsPack(conn *websocket.Conn, packType int) *WsPack {
	pack := new(WsPack)
	pack.conn = conn
	pack.packType = packType

	pack.Time = timer.GetNowDate()
	return pack
}

func (p *WsPack) Parse(jsonByte []byte) {
	jsonObj, err := json.Unmarshal(jsonByte)
	if err != nil {
		return
	}
	action := json.GetString(jsonObj, "action")
	//msg :=json.GetString(jsonObj, "msg")

	if action == WS_PACK_ACTION_PING {
		p.PingPack()
	}
}

func (p *WsPack) PingPack() {
	p.Action = WS_PACK_ACTION_PING
	p.Msg = "pong"
	msg, err := json.Marshal(p.Message)
	if err == nil {
		p.conn.WriteMessage(websocket.TextMessage, msg)
	}
}
