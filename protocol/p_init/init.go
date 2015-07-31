package p_init

import (
	"github.com/colefan/gsgo/netio/iobuffer"
	"github.com/colefan/gsgo/netio/packet"
)

const (
	MSG_TYPE_INIT = 0x02	//初始化模块
)

const (
	CMD_INIT_TIMESTAMP = 0x0201
)

type InitTimeStampNt struct {
	*packet.Packet
	TimeStamp	uint32 //服务器同步时间搓
	SyncCode	string //扩展的同步码
}

func (this *InitTimeStampNt) DecodePacket() bool {
	if this.IsDecoded() {
		return true
	}
	packet.DecoderReadValue(this.Packet, &this.TimeStamp)
	packet.DecoderReadValue(this.Packet, &this.SyncCode)
	this.PackDecoded = true
	return true
}

func (this *InitTimeStampNt) EncodePacket(nLen int) *iobuffer.OutBuffer {
	buf := iobuffer.NewOutBuffer(nLen)
	buf = this.Packet.Header.Encode(buf)
	buf.PutRawValue(this.TimeStamp)
	buf.PutRawValue(this.SyncCode)
	nPackLen := buf.GetLen() - packet.PACKET_PROXY_HEADER_LEN
	buf.SetUint16(uint16(nPackLen), 0)
	return buf
}
