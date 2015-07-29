package gs

import (
	"github.com/colefan/gsgo-hgame/handler"
	. "github.com/colefan/gsgo-hgame/init"
	"github.com/colefan/gsgo/netio"
	"github.com/colefan/gsgo/netio/packet"
)

var server *GameServer

type GameServer struct {
	netio.BaseServer
	handlers map[uint16]handler.HandlerInf
}

func GetGameServerInst() *GameServer {
	if server == nil {
		server = &GameServer{}
		server.handlers = make(map[uint16]handler.HandlerInf)
	}
	return server
}

func (this *GameServer) SessionOpen(conn netio.ConnInf) {

}

func (this *GameServer) SessionClose(conn netio.ConnInf) {

}

func (this *GameServer) HandleMsg(cmdId uint16, pack *packet.Packet, conn netio.ConnInf) {

}

func (this *GameServer) RegisterHandler(moduleId uint16, handler handler.HandlerInf) bool {
	tmp := this.handlers[moduleId]
	if tmp != nil {
		Log.Info("RegisterHandler error, repeated moduleId = ", moduleId)
		return false
	}
	this.handlers[moduleId] = handler
	Log.Info("RegisterHandler success, moduleId = ", moduleId)
	return true
}
