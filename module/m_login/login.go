//登录模块
package m_login

import (
	"github.com/colefan/gsgo-hgame/config"
	"github.com/colefan/gsgo-hgame/handler"
	. "github.com/colefan/gsgo-hgame/init"
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
		this.roleStatusReq(pack, conn)
	case p_login.CMD_C_CRETE_ROLE_REQ:
		this.createRoleReq(pack, conn)
	case p_login.CMD_C_ROLE_LIST_REQ:
		this.roleListReq(pack, conn)
	case p_login.CMD_C_SET_PAY_ROLE_REQ:
		this.setPayRoleReq(pack, conn)
	default:
		this.UnHandledMsg(cmdid, conn)
	}
}

func (this *LoginHandler) loginReq(pack *packet.Packet, conn netio.ConnInf) {
	//TODO
	if conn.GetBindObj() != nil {
		//登录后在发登录请求，这是不对的。应该剔除
		Log.Error("第一次消息请求时，应该没有已经绑定的对象，说明用户已经有绑定的对象了，是个错误的请求")
		return
	}
	if pack.DecodePacket() == false {

	}

	var loginReq p_login.PlayerLoginReq
	loginReq.Packet = pack
	if false == loginReq.DecodePacket() {
		Log.Error("解析包错误，应该断开连接")
		conn.Close()
		return
	}

	var uid uint64
	var account string
	var superpwd string
	uid = loginReq.UserID
	account = loginReq.UserName
	superpwd = loginReq.UserPwd

	if config.GetServerConfig().GetXmlConf().Server.AccessPtMode == 1 {
		//平台接入模式，要查看平台登录信息
	} else {
		//非平台接入模式，只要查看自己的登录信息即可
	}

}

func (this *LoginHandler) loginReadyReq(pack *packet.Packet, conn netio.ConnInf) {
	//TODO
}

func (this *LoginHandler) roleStatusReq(pack *packet.Packet, conn netio.ConnInf) {

}

func (this *LoginHandler) createRoleReq(pack *packet.Packet, conn netio.ConnInf) {

}

func (this *LoginHandler) roleListReq(pack *packet.Packet, conn netio.ConnInf) {

}

func (this *LoginHandler) setPayRoleReq(pack *packet.Packet, conn netio.ConnInf) {

}
func init() {
	instLogin = &LoginHandler{}
}
