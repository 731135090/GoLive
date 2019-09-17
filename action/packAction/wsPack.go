package packAction

import (
	"GoLive/config"
	"GoLive/uitl/json"
	"GoLive/uitl/timer"
	"github.com/bitly/go-simplejson"
	"github.com/gorilla/websocket"
)

const (
	WS_PACK_TYPE_CHEAT    = iota //聊天消息
	WS_PACK_TYPE_SYSTEM          //系统消息
	WS_PACK_TYPE_CUSTOMER        //客服消息
)

const (
	WS_PACK_ACTION_PING  = "ping"
	WS_PACK_ACTION_LOGIN = "login"
	WS_PACK_ACTION_SEND  = "send"
	WS_PACK_ACTION_CLOSE = "close"
	WS_PACK_ACTION_MSG   = "msg"
	WS_PACK_ACTION_IMG   = "img"
	WS_PACK_ACTION_ICON  = "icon"

	WS_MES_TYPE_TEXT = "text"
	WS_MES_TYPE_IMG  = "img"
)

var WsPackChannel = make(chan *WsPack, 1000000)

type Message struct {
	Action  string `json:"action"`  //pack type
	Msg     string `json:"msg"`     //message
	MsgType string `json:"type"`    //msg type
	Time    string `json:"time"`    //server time
}
type WsPack struct {
	conn     *websocket.Conn
	form     string //pack form
	to       string //pack to
	packType int    //pack type
	data     []byte
}

func init() {
	config.Gwg.Add(1)
	go func() {
		defer config.Gwg.Done()
		for {
			wsPack, ok := <-WsPackChannel
			if !ok {
				break
			}
			go wsPack.Parse()
		}
	}()
}

// new ws pack
func NewWsPack(conn *websocket.Conn, packType int, data []byte) *WsPack {
	pack := new(WsPack)
	pack.conn = conn
	pack.packType = packType
	pack.data = data
	return pack
}

// parse ws pack
func (p *WsPack) Parse() {
	jsonObj, err := json.Unmarshal(p.data)
	if err != nil {
		return
	}
	action := json.GetString(jsonObj, "action")
	//msg := json.GetString(jsonObj, "msg")

	switch action {
	case WS_PACK_ACTION_PING:
		p.PingPack()
	case WS_PACK_ACTION_CLOSE:
		p.ClosePack()
	case WS_PACK_ACTION_MSG:
		p.MsgPack(jsonObj)
	}
}

//ws pack close
func (p *WsPack) ClosePack() {
	message := Message{
		Msg:     "ok",
		Action:  WS_PACK_ACTION_CLOSE,
		MsgType: WS_MES_TYPE_TEXT,
		Time:    timer.GetNowDate(),
	}
	msg, err := json.Marshal(message)
	if err != nil {
		config.Logger.Error(err.Error())
	}
	p.conn.WriteMessage(websocket.TextMessage, msg)

	err = p.conn.Close()
	if err != nil {
		config.Logger.Error(err.Error())
	}
}

// ws pack ping
func (p *WsPack) PingPack() {
	message := Message{
		Msg:     "ok",
		Action:  WS_PACK_ACTION_PING,
		MsgType: WS_MES_TYPE_TEXT,
		Time:    timer.GetNowDate(),
	}
	msg, err := json.Marshal(message)
	if err != nil {
		config.Logger.Error(err.Error())
	}
	p.conn.WriteMessage(websocket.TextMessage, msg)
}

func (p *WsPack) MsgPack(data *simplejson.Json) {
	content := json.GetString(data, "msg")
	message := Message{
		Msg:     content,
		Action:  WS_PACK_ACTION_MSG,
		MsgType: WS_MES_TYPE_TEXT,
		Time:    timer.GetNowDate(),
	}
	msg, err := json.Marshal(message)
	if err != nil {
		config.Logger.Error(err.Error())
	}
	p.conn.WriteMessage(websocket.TextMessage, msg)
}
