package p_login

import (
	"github.com/colefan/gsgo/netio/iobuffer"
	"github.com/colefan/gsgo/netio/packet"
)

const (
	MSG_TYPE_LOGIN = 0x01	//登录模块
)

const (
	CMD_C_LOGIN_REQ = 0x0101
	CMD_C_LOGIN_RESP = 0x0102
)

type PlayerLoginReq struct {
	*packet.Packet
	UserName	string //账户
	UserPwd	uint8 //密码
	UserID	uint64 //用户ID
}

func (this *PlayerLoginReq) DecodePacket() bool {
	if this.IsDecoded() {
		return true
	}
	packet.DecoderReadValue(this.Packet, &this.UserName)
	packet.DecoderReadValue(this.Packet, &this.UserPwd)
	packet.DecoderReadValue(this.Packet, &this.UserID)
	this.PackDecoded = true
	return true
}

func (this *PlayerLoginReq) EncodePacket(nLen int) *iobuffer.OutBuffer {
	buf := iobuffer.NewOutBuffer(nLen)
	buf = this.Packet.Header.Encode(buf)
	buf.PutRawValue(this.UserName)
	buf.PutRawValue(this.UserPwd)
	buf.PutRawValue(this.UserID)
	nPackLen := buf.GetLen() - packet.PACKET_PROXY_HEADER_LEN
	buf.SetUint16(uint16(nPackLen), 0)
	return buf
}
type PlayerLoginResp struct {
	*packet.Packet
	ErrCode	uint16 //
}

func (this *PlayerLoginResp) DecodePacket() bool {
	if this.IsDecoded() {
		return true
	}
	packet.DecoderReadValue(this.Packet, &this.ErrCode)
	this.PackDecoded = true
	return true
}

func (this *PlayerLoginResp) EncodePacket(nLen int) *iobuffer.OutBuffer {
	buf := iobuffer.NewOutBuffer(nLen)
	buf = this.Packet.Header.Encode(buf)
	buf.PutRawValue(this.ErrCode)
	nPackLen := buf.GetLen() - packet.PACKET_PROXY_HEADER_LEN
	buf.SetUint16(uint16(nPackLen), 0)
	return buf
}
