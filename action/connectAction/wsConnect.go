package connectAction

import (
	"GoLive/config"
	"github.com/gorilla/websocket"
	"sync"
)

var wsConnLock = sync.Mutex{}

type WsConnect struct {
	WsConn *websocket.Conn
}

//PACK_TYPE : uid : conn :bool
var WsConnMap = make(map[int]map[string]map[*websocket.Conn]bool)

func init() {
	WsConnMap[config.WS_CONN_TYPE_USER] = make(map[string]map[*websocket.Conn]bool)
	WsConnMap[config.WS_CONN_TYPE_CUSTOMER] = make(map[string]map[*websocket.Conn]bool)
}

func NewWsConn(conn *websocket.Conn) *WsConnect {
	wsConn := new(WsConnect)
	wsConn.WsConn = conn
	return wsConn
}

func CloseWsConn(conn *websocket.Conn, wsConnType int, from string) {
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
