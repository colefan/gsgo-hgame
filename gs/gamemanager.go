package gs

import (
	"github.com/colefan/gsgo-hgame/config"
	. "github.com/colefan/gsgo-hgame/init"
)

type GameManager struct {
}

func NewGameManager() *GameManager {
	m := &GameManager{}
	return m
}

func (this *GameManager) Init(buildtime string) bool {
	var bVal bool = false
	var err error

	//config loading
	bVal = config.GetServerConfig().InitConfig()
	if bVal {
		Log.Info(">>ServerConfig.InitConfig Success!")
	} else {
		Log.Info(">>ServerConfig.InitConfig failed please check serverconfig.xml")
		return false
	}

	//sql init

	//
	server := GetGameServerInst()
	Log.Info(">>GameServer Initing...")
	err = server.InitServer(server, server, "gameserver", config.GetServerConfig().GetXmlConf().Server.Ip, config.GetServerConfig().GetXmlConf().Server.Port)
	if err != nil {
		Log.Info(">>GameServer Init error,", err)
		return false
	}
	Log.Info(">>GameServer Init success!")

	return true
}

func (this *GameManager) Run() {

}

func (this *GameManager) ShutDown() {

}
