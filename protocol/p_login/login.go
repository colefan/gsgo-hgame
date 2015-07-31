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
	CMD_C_LOGIN_READY_REQ = 0x0103
	CMD_C_LOGIN_READY_RESP = 0x0104
	CMD_C_ROLE_STATUS_REQ = 0x0105
	CMD_C_ROLE_STATUS_RESP = 0x0106
	CMD_C_CRETE_ROLE_REQ = 0x0107
	CMD_C_CRETE_ROLE_RESP = 0x0108
	CMD_C_ROLE_LIST_REQ = 0x0109
	CMD_C_ROLE_LIST_RESP = 0x0120
	CMD_C_SET_PAY_ROLE_REQ = 0x0121
	CMD_C_SET_PAY_ROLE_RESP = 0x0122
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
type PlayerLoginReadyReq struct {
	*packet.Packet
	ValidCode	string //预留校验码
}

func (this *PlayerLoginReadyReq) DecodePacket() bool {
	if this.IsDecoded() {
		return true
	}
	packet.DecoderReadValue(this.Packet, &this.ValidCode)
	this.PackDecoded = true
	return true
}

func (this *PlayerLoginReadyReq) EncodePacket(nLen int) *iobuffer.OutBuffer {
	buf := iobuffer.NewOutBuffer(nLen)
	buf = this.Packet.Header.Encode(buf)
	buf.PutRawValue(this.ValidCode)
	nPackLen := buf.GetLen() - packet.PACKET_PROXY_HEADER_LEN
	buf.SetUint16(uint16(nPackLen), 0)
	return buf
}
type PlayerLoginReadyResp struct {
	*packet.Packet
	ErrCode	uint16 //错误码
}

func (this *PlayerLoginReadyResp) DecodePacket() bool {
	if this.IsDecoded() {
		return true
	}
	packet.DecoderReadValue(this.Packet, &this.ErrCode)
	this.PackDecoded = true
	return true
}

func (this *PlayerLoginReadyResp) EncodePacket(nLen int) *iobuffer.OutBuffer {
	buf := iobuffer.NewOutBuffer(nLen)
	buf = this.Packet.Header.Encode(buf)
	buf.PutRawValue(this.ErrCode)
	nPackLen := buf.GetLen() - packet.PACKET_PROXY_HEADER_LEN
	buf.SetUint16(uint16(nPackLen), 0)
	return buf
}
type RoleStatusReq struct {
	*packet.Packet
}

func (this *RoleStatusReq) DecodePacket() bool {
	if this.IsDecoded() {
		return true
	}
	this.PackDecoded = true
	return true
}

func (this *RoleStatusReq) EncodePacket(nLen int) *iobuffer.OutBuffer {
	buf := iobuffer.NewOutBuffer(nLen)
	buf = this.Packet.Header.Encode(buf)
	nPackLen := buf.GetLen() - packet.PACKET_PROXY_HEADER_LEN
	buf.SetUint16(uint16(nPackLen), 0)
	return buf
}
type RoleStatusResp struct {
	*packet.Packet
	RoleStatus	uint16 //0:代表没有角色；1：代表有角色；2代表有多个角色
}

func (this *RoleStatusResp) DecodePacket() bool {
	if this.IsDecoded() {
		return true
	}
	packet.DecoderReadValue(this.Packet, &this.RoleStatus)
	this.PackDecoded = true
	return true
}

func (this *RoleStatusResp) EncodePacket(nLen int) *iobuffer.OutBuffer {
	buf := iobuffer.NewOutBuffer(nLen)
	buf = this.Packet.Header.Encode(buf)
	buf.PutRawValue(this.RoleStatus)
	nPackLen := buf.GetLen() - packet.PACKET_PROXY_HEADER_LEN
	buf.SetUint16(uint16(nPackLen), 0)
	return buf
}
type CreateRoleReq struct {
	*packet.Packet
	RoleName	string //角色名字
	RoleJob	uint8 //角色职业
}

func (this *CreateRoleReq) DecodePacket() bool {
	if this.IsDecoded() {
		return true
	}
	packet.DecoderReadValue(this.Packet, &this.RoleName)
	packet.DecoderReadValue(this.Packet, &this.RoleJob)
	this.PackDecoded = true
	return true
}

func (this *CreateRoleReq) EncodePacket(nLen int) *iobuffer.OutBuffer {
	buf := iobuffer.NewOutBuffer(nLen)
	buf = this.Packet.Header.Encode(buf)
	buf.PutRawValue(this.RoleName)
	buf.PutRawValue(this.RoleJob)
	nPackLen := buf.GetLen() - packet.PACKET_PROXY_HEADER_LEN
	buf.SetUint16(uint16(nPackLen), 0)
	return buf
}
type CreateRoleResp struct {
	*packet.Packet
	ErrCode	uint16 //错误码
}

func (this *CreateRoleResp) DecodePacket() bool {
	if this.IsDecoded() {
		return true
	}
	packet.DecoderReadValue(this.Packet, &this.ErrCode)
	this.PackDecoded = true
	return true
}

func (this *CreateRoleResp) EncodePacket(nLen int) *iobuffer.OutBuffer {
	buf := iobuffer.NewOutBuffer(nLen)
	buf = this.Packet.Header.Encode(buf)
	buf.PutRawValue(this.ErrCode)
	nPackLen := buf.GetLen() - packet.PACKET_PROXY_HEADER_LEN
	buf.SetUint16(uint16(nPackLen), 0)
	return buf
}
type LoginRoleEntity struct {
	UserID	uint64 //账号ID
	RoleID	uint64 //角色ID
	OldServerId	string //原服ID
	RoleName	string //角色名称
	Level	uint16 //等级
	RoleJob	uint8 //职业
	RoleIcon	uint32 //头像
	IsPayRole	uint8 //是否充值角色（0：否；1：是）
	LastLoginTime	uint32 //最后登录时间
}

func (this *LoginRoleEntity) DecodeEntity(p *packet.Packet) bool {
	packet.DecoderReadValue(p, &this.UserID)
	packet.DecoderReadValue(p, &this.RoleID)
	packet.DecoderReadValue(p, &this.OldServerId)
	packet.DecoderReadValue(p, &this.RoleName)
	packet.DecoderReadValue(p, &this.Level)
	packet.DecoderReadValue(p, &this.RoleJob)
	packet.DecoderReadValue(p, &this.RoleIcon)
	packet.DecoderReadValue(p, &this.IsPayRole)
	packet.DecoderReadValue(p, &this.LastLoginTime)
	return true
}

func (this *LoginRoleEntity) EncodeEntity(buf *iobuffer.OutBuffer) *iobuffer.OutBuffer {
	buf.PutRawValue(this.UserID)
	buf.PutRawValue(this.RoleID)
	buf.PutRawValue(this.OldServerId)
	buf.PutRawValue(this.RoleName)
	buf.PutRawValue(this.Level)
	buf.PutRawValue(this.RoleJob)
	buf.PutRawValue(this.RoleIcon)
	buf.PutRawValue(this.IsPayRole)
	buf.PutRawValue(this.LastLoginTime)
	return buf
}
type RoleListReq struct {
	*packet.Packet
	UserName	string //账号
}

func (this *RoleListReq) DecodePacket() bool {
	if this.IsDecoded() {
		return true
	}
	packet.DecoderReadValue(this.Packet, &this.UserName)
	this.PackDecoded = true
	return true
}

func (this *RoleListReq) EncodePacket(nLen int) *iobuffer.OutBuffer {
	buf := iobuffer.NewOutBuffer(nLen)
	buf = this.Packet.Header.Encode(buf)
	buf.PutRawValue(this.UserName)
	nPackLen := buf.GetLen() - packet.PACKET_PROXY_HEADER_LEN
	buf.SetUint16(uint16(nPackLen), 0)
	return buf
}
type RoleListResp struct {
	*packet.Packet
	RoleList	[]LoginRoleEntity //角色列表
}

func (this *RoleListResp) DecodePacket() bool {
	if this.IsDecoded() {
		return true
	}
	arrLen:=packet.DecoderReadArrayLength(this.Packet)
	for i :=0; i < arrLen; i++ {
		e := &LoginRoleEntity{}
		packet.DecoderReadEntity(this.Packet, e)
		this.RoleList = append(this.RoleList, *e)
	}
	this.PackDecoded = true
	return true
}

func (this *RoleListResp) EncodePacket(nLen int) *iobuffer.OutBuffer {
	buf := iobuffer.NewOutBuffer(nLen)
	buf = this.Packet.Header.Encode(buf)
	if len(this.RoleList) > 0 {
		buf.PutRawValue(uint16(len(this.RoleList)))
		for _,tmp := range this.RoleList {
			buf = tmp.EncodeEntity(buf)
		}
	}
	nPackLen := buf.GetLen() - packet.PACKET_PROXY_HEADER_LEN
	buf.SetUint16(uint16(nPackLen), 0)
	return buf
}
type SetPayRoleReq struct {
	*packet.Packet
	UserID	uint64 //账号ID
	RoleID	uint64 //角色ID
	RoleName	string //角色名
}

func (this *SetPayRoleReq) DecodePacket() bool {
	if this.IsDecoded() {
		return true
	}
	packet.DecoderReadValue(this.Packet, &this.UserID)
	packet.DecoderReadValue(this.Packet, &this.RoleID)
	packet.DecoderReadValue(this.Packet, &this.RoleName)
	this.PackDecoded = true
	return true
}

func (this *SetPayRoleReq) EncodePacket(nLen int) *iobuffer.OutBuffer {
	buf := iobuffer.NewOutBuffer(nLen)
	buf = this.Packet.Header.Encode(buf)
	buf.PutRawValue(this.UserID)
	buf.PutRawValue(this.RoleID)
	buf.PutRawValue(this.RoleName)
	nPackLen := buf.GetLen() - packet.PACKET_PROXY_HEADER_LEN
	buf.SetUint16(uint16(nPackLen), 0)
	return buf
}
type SetPayRoleResp struct {
	*packet.Packet
	ErrCode	uint16 //错误码
}

func (this *SetPayRoleResp) DecodePacket() bool {
	if this.IsDecoded() {
		return true
	}
	packet.DecoderReadValue(this.Packet, &this.ErrCode)
	this.PackDecoded = true
	return true
}

func (this *SetPayRoleResp) EncodePacket(nLen int) *iobuffer.OutBuffer {
	buf := iobuffer.NewOutBuffer(nLen)
	buf = this.Packet.Header.Encode(buf)
	buf.PutRawValue(this.ErrCode)
	nPackLen := buf.GetLen() - packet.PACKET_PROXY_HEADER_LEN
	buf.SetUint16(uint16(nPackLen), 0)
	return buf
}
