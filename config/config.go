package config

import (
	"GoLive/uitl/dir"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	_ "github.com/astaxie/beego/config/yaml"
	"github.com/astaxie/beego/logs"
	"sync"
)

var Logger *logs.BeeLogger

var Gwg = sync.WaitGroup{}

var (
	AppName  string
	RunLevel string
	HttpPort string
	ListenIp string
	RootDir  string
)

func init() {
	initLogger()
	initConfig()
	loadConfigFile()
}

func initConfig() {
	RootDir = dir.GetPwd()
	if RunLevel != RUN_LEVEL_RELEASE {
		Logger.Info("RootPath: %s\n", RootDir)
	}

	//set static path
	beego.SetStaticPath("/static", "static")
}
func initLogger() {
	Logger = logs.NewLogger()
	Logger.SetLogger(logs.AdapterConsole)
	Logger.EnableFuncCallDepth(true)
}
func loadConfigFile() {
	conf, err := config.NewConfig("yaml", RootDir+"/conf/app.yaml")
	if err != nil {
		Logger.Error(err.Error())
		panic(err)
	}
	AppName = conf.DefaultString("app_name", "")
	RunLevel = conf.DefaultString("run_level", "dev")
	ListenIp = conf.DefaultString("listen_ip", "127.0.0.1")
	HttpPort = conf.DefaultString("http_port", "8080")
}
