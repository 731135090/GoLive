package connectAction

import (
	"GoLive/config"
	"github.com/gorilla/websocket"
	"sync"
)

var wsConnLock = sync.Mutex{}

type WsConnect struct {
	WsConn   *websocket.Conn
	ConnType uint8
}

//PACK_TYPE : uid : conn :bool
var WsConnMap = make(map[uint8]map[string]map[*websocket.Conn]bool)

func init() {
	WsConnMap[config.WS_CONN_TYPE_USER] = make(map[string]map[*websocket.Conn]bool)
	WsConnMap[config.WS_CONN_TYPE_CUSTOMER] = make(map[string]map[*websocket.Conn]bool)
}

func NewWsConn(conn *websocket.Conn, connType uint8) *WsConnect {
	wsConn := new(WsConnect)
	wsConn.WsConn = conn
	wsConn.ConnType = connType
	return wsConn
}

func CloseWsConn(conn *websocket.Conn, wsConnType uint8, from string) {
	conn.Close()
	if _, ok := WsConnMap[wsConnType][from]; ok {
		wsConnLock.Lock()
		delete(WsConnMap[wsConnType][from], conn)
		if len(WsConnMap[wsConnType][from]) == 0 {
			delete(WsConnMap[wsConnType], from)
		}
		wsConnLock.Unlock()
	}
}
