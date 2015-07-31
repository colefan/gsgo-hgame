package handler

import (
	. "github.com/colefan/gsgo-hgame/init"
	"github.com/colefan/gsgo/netio"
	"github.com/colefan/gsgo/netio/packet"
)

type HandlerInf interface {
	SetModuleId(id uint16)
	GetModuleId() uint16
	HandleMsg(cmdid uint16, pack *packet.Packet, conn netio.ConnInf)
}

type BaseHandler struct {
	ModuleID uint16
}

func (h *BaseHandler) SetModuleId(id uint16) {
	h.ModuleID = id
}

func (h *BaseHandler) GetModuleId() uint16 {
	return h.ModuleID
}

func (h *BaseHandler) UnHandledMsg(cmdid uint16, conn netio.ConnInf) {
	Log.Warn("unknow msg in this module, msgtype :", h.GetModuleId(), ", cmdid : ", cmdid, ", userid : ", conn.GetUID())
}
