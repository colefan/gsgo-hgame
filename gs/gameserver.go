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
	//打开一个物理链接，为这个物理链接分配一些对象吧

}

func (this *GameServer) SessionClose(conn netio.ConnInf) {

}

func (this *GameServer) HandleMsg(cmdId uint16, pack *packet.Packet, conn netio.ConnInf) {
	msgType := cmdId >> 8
	h := this.handlers[msgType]
	if h == nil {
		Log.Warn("msg has no handler ,cmdid = ", cmdId, ", msgtype = ", msgType)

		//TODO 是否需要断掉
		conn.Close()
	} else {
		h.HandleMsg(cmdId, pack, conn)
	}

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
