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
	sender   string //pack form
	receiver string //pack to
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
func NewWsPack(conn *websocket.Conn, connType uint8, data []byte) *WsPack {
	pack := new(WsPack)
	pack.conn = connectAction.NewWsConn(conn, connType)
	pack.data = data
	return pack
}

// parse ws pack
func (p *WsPack) Parse() {
	jsonObj, err := json.Unmarshal(p.data)
	if err != nil {
		config.Logger.Error(err.Error())
		return
	}
	action := json.GetString(jsonObj, "action")
	if action == config.WS_PACK_ACTION_PING {
		p.PingPack()
		return
	}
	p.sender = json.GetString(jsonObj, "sender")
	if p.sender == "" {
		config.Logger.Warn("sender is empty")
		return
	}
	p.receiver = json.GetString(jsonObj, "receiver")
	if p.receiver == "" {
		config.Logger.Warn("receiver is empty")
		return
	}
	p.InitPack()
	switch action {
	case config.WS_PACK_ACTION_CLOSE:
		p.ClosePack()
	case config.WS_PACK_ACTION_MSG:
		p.MsgPack(jsonObj)
	}
}

func (p *WsPack) InitPack() {
	if connectAction.WsConnMap[p.conn.ConnType][p.sender] == nil {
		connectAction.WsConnMap[p.conn.ConnType][p.sender] = make(map[*websocket.Conn]bool)
	}

	if _, ok := connectAction.WsConnMap[p.conn.ConnType][p.sender][p.conn.WsConn]; !ok {
		connectAction.WsConnMap[p.conn.ConnType][p.sender][p.conn.WsConn] = true
	}
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

	connectAction.CloseWsConn(p.conn.WsConn, config.WS_CONN_TYPE_USER, p.sender)
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
	msgType := json.GetString(data, "type")
	if _, ok := config.AllowMessType[msgType]; !ok {
		config.Logger.Warn("msg type is error")
		return
	}
	message := Message{
		Msg:     content,
		Action:  config.WS_PACK_ACTION_MSG,
		MsgType: msgType,
		Time:    timer.GetNowDate(),
	}
	msg, err := json.Marshal(message)
	if err != nil {
		config.Logger.Error(err.Error())
		return
	}
	receiverConnMap, ok := connectAction.WsConnMap[getReceiverConnType(p.conn.ConnType)][p.receiver]
	if ok {
		for conn, _ := range receiverConnMap {
			conn.WriteMessage(websocket.TextMessage, msg)
		}
	} else {
		// todo 客服不在线处理逻辑
	}
}

func getReceiverConnType(senderConnType uint8) uint8 {
	switch senderConnType {
	case config.WS_CONN_TYPE_CUSTOMER:
		return config.WS_CONN_TYPE_USER
	case config.WS_CONN_TYPE_USER:
		return config.WS_CONN_TYPE_CUSTOMER
	}
	return config.WS_CONN_TYPE_USER
}
