//登录模块
package m_login

import (
	"github.com/colefan/gsgo-hgame/handler"
	"github.com/colefan/gsgo/netio"
	"github.com/colefan/gsgo/netio/packet"
)

var instLogin *LoginHandler

type LoginHandler struct {
	handler.BaseHandler
}

func GetLoginHandlerInst() *LoginHandler {
	return instLogin
}

func (this *LoginHandler) HandleMsg(cmdid uint16, pack *packet.Packet, conn netio.ConnInf) {

}

func init() {
	instLogin = &LoginHandler{}
}
