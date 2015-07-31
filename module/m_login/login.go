//登录模块
package m_login

import (
	"github.com/colefan/gsgo-hgame/handler"
	"github.com/colefan/gsgo-hgame/protocol/p_login"
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
	switch cmdid {
	case p_login.CMD_C_LOGIN_REQ:
		this.loginReq(pack, conn)
	case p_login.CMD_C_LOGIN_READY_REQ:
		this.loginReadyReq(pack, conn)
	case p_login.CMD_C_ROLE_STATUS_REQ:
	case p_login.CMD_C_CRETE_ROLE_REQ:
	case p_login.CMD_C_ROLE_LIST_REQ:
	case p_login.CMD_C_SET_PAY_ROLE_REQ:
	default:
		this.UnHandledMsg(cmdid, conn)
	}
}

func (this *LoginHandler) loginReq(pack *packet.Packet, conn netio.ConnInf) {
	//TODO
}

func (this *LoginHandler) loginReadyReq(pack *packet.Packet, conn netio.ConnInf) {
	//TODO
}

func (this *LoginHandler) roleStatusReq(pack *packet.Packet, conn netio.ConnInf) {

}
func init() {
	instLogin = &LoginHandler{}
}
