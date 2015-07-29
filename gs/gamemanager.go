package gs

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/colefan/gsgo-hgame/config"
	. "github.com/colefan/gsgo-hgame/init"
	"github.com/colefan/gsgo-hgame/module/m_login"
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

	Log.Info(">>Init game modules...")

	return true
}

func (this *GameManager) initModules() bool {
	var bVal bool = false
	bVal = GetGameServerInst().RegisterHandler(m_login.GetLoginHandlerInst().GetModuleId(), m_login.GetLoginHandlerInst())
	if !bVal {
		return false
	}

	return true
}

func (this *GameManager) Run() {
	GetGameServerInst().Run()
	this.checkinput()
}

func (this *GameManager) ShutDown() {
	GetGameServerInst().ShutDown()
}

func (this *GameManager) checkinput() {
	fmt.Println("please entry q or Q to quit the progame!")
	reader := bufio.NewReader(os.Stdin)
	for {

		data, _, _ := reader.ReadLine()
		command := string(data)
		if command == "q" || command == "Q" {
			break
		}
		fmt.Println("please entry q or Q to quit the progame!")
		time.Sleep(1 * time.Second)
	}
}
