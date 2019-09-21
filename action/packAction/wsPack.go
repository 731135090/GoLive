package packAction

import (
	"GoLive/action/connectAction"
	"GoLive/config"
	"GoLive/uitl/json"
	"GoLive/uitl/timer"
	"github.com/bitly/go-simplejson"
	"github.com/gorilla/websocket"
)

var WsPackChannel = make(chan *WsPack, 1000000)

type Message struct {
	Action  string `json:"action"` //pack type
	Msg     string `json:"msg"`    //message
	MsgType string `json:"type"`   //msg type
	Time    string `json:"time"`   //server time
}
type WsPack struct {
	conn     *connectAction.WsConnect
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
	pack.conn = connectAction.NewWsConn(conn)
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
	case config.WS_PACK_ACTION_INIT:
		p.InitPack()
	case config.WS_PACK_ACTION_PING:
		p.PingPack()
	case config.WS_PACK_ACTION_CLOSE:
		p.ClosePack()
	case config.WS_PACK_ACTION_MSG:
		p.MsgPack(jsonObj)
	}
}

func (p *WsPack) InitPack() {
	if connectAction.WsConnMap[config.WS_CONN_TYPE_CHEAT][p.form] == nil {
		connectAction.WsConnMap[config.WS_CONN_TYPE_CHEAT][p.form] = make(map[*websocket.Conn]bool)
	}
	connectAction.WsConnMap[config.WS_CONN_TYPE_CHEAT][p.form][p.conn.WsConn] = true
}

//ws pack close
func (p *WsPack) ClosePack() {
	message := Message{
		Msg:     "ok",
		Action:  config.WS_PACK_ACTION_CLOSE,
		MsgType: config.WS_MES_TYPE_TEXT,
		Time:    timer.GetNowDate(),
	}
	msg, err := json.Marshal(message)
	if err != nil {
		config.Logger.Error(err.Error())
	}
	p.conn.WsConn.WriteMessage(websocket.TextMessage, msg)

	err = p.conn.WsConn.Close()
	if err != nil {
		config.Logger.Error(err.Error())
	}

	connectAction.CloseWsConn(p.conn.WsConn, config.WS_CONN_TYPE_CHEAT, p.form)
}

// ws pack ping
func (p *WsPack) PingPack() {
	message := Message{
		Msg:     "ok",
		Action:  config.WS_PACK_ACTION_PING,
		MsgType: config.WS_MES_TYPE_TEXT,
		Time:    timer.GetNowDate(),
	}
	msg, err := json.Marshal(message)
	if err != nil {
		config.Logger.Error(err.Error())
	}
	p.conn.WsConn.WriteMessage(websocket.TextMessage, msg)
}

func (p *WsPack) MsgPack(data *simplejson.Json) {
	content := json.GetString(data, "msg")
	message := Message{
		Msg:     content,
		Action:  config.WS_PACK_ACTION_MSG,
		MsgType: config.WS_MES_TYPE_TEXT,
		Time:    timer.GetNowDate(),
	}
	msg, err := json.Marshal(message)
	if err != nil {
		config.Logger.Error(err.Error())
	}
	p.conn.WsConn.WriteMessage(websocket.TextMessage, msg)
}
