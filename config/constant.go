package config

//定义常量

const (
	RUN_LEVEL_DEBUG   = "debug"
	RUN_LEVEL_TEST    = "test"
	RUN_LEVEL_RELEASE = "release"
)

const (
	WS_CONN_TYPE_CHEAT    = iota //聊天链接
	WS_CONN_TYPE_SYSTEM          //系统链接
	WS_CONN_TYPE_CUSTOMER        //客服链接
)

const (
	WS_PACK_ACTION_INIT  = "init"  //初始化聊天信息
	WS_PACK_ACTION_PING  = "ping"  //心跳包
	WS_PACK_ACTION_CLOSE = "close" //关闭包
	WS_PACK_ACTION_MSG   = "msg"   //信息包
	WS_PACK_ACTION_ICON  = "icon"
)

const (
	WS_MES_TYPE_TEXT  = "text"  //信息类型：文字
	WS_MES_TYPE_IMG   = "img"   //信息类型：图片
	WS_MES_TYPE_VOICE = "voice" //信息类型：声音
)
