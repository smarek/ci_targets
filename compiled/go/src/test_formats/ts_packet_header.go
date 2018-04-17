// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package test_formats

import "github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"


/**
 * describes the first 4 header bytes of a TS Packet header
 */

type TsPacketHeader_AdaptationFieldControlEnum int
const (
	TsPacketHeader_AdaptationFieldControlEnum__Reserved TsPacketHeader_AdaptationFieldControlEnum = 0
	TsPacketHeader_AdaptationFieldControlEnum__PayloadOnly TsPacketHeader_AdaptationFieldControlEnum = 1
	TsPacketHeader_AdaptationFieldControlEnum__AdaptationFieldOnly TsPacketHeader_AdaptationFieldControlEnum = 2
	TsPacketHeader_AdaptationFieldControlEnum__AdaptationFieldAndPayload TsPacketHeader_AdaptationFieldControlEnum = 3
)
type TsPacketHeader struct {
	SyncByte uint8
	TransportErrorIndicator bool
	PayloadUnitStartIndicator bool
	TransportPriority bool
	Pid uint64
	TransportScramblingControl uint64
	AdaptationFieldControl TsPacketHeader_AdaptationFieldControlEnum
	ContinuityCounter uint64
	TsPacketRemain []byte
	_io *kaitai.Stream
	_root *TsPacketHeader
	_parent interface{}
}

func (this *TsPacketHeader) Read(io *kaitai.Stream, parent interface{}, root *TsPacketHeader) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp1, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.SyncByte = tmp1
	tmp2, err := this._io.ReadBitsInt(1) != 0
	if err != nil {
		return err
	}
	this.TransportErrorIndicator = tmp2
	tmp3, err := this._io.ReadBitsInt(1) != 0
	if err != nil {
		return err
	}
	this.PayloadUnitStartIndicator = tmp3
	tmp4, err := this._io.ReadBitsInt(1) != 0
	if err != nil {
		return err
	}
	this.TransportPriority = tmp4
	tmp5, err := this._io.ReadBitsInt(13)
	if err != nil {
		return err
	}
	this.Pid = tmp5
	tmp6, err := this._io.ReadBitsInt(2)
	if err != nil {
		return err
	}
	this.TransportScramblingControl = tmp6
	tmp7, err := this._io.ReadBitsInt(2)
	if err != nil {
		return err
	}
	this.AdaptationFieldControl = TsPacketHeader_AdaptationFieldControlEnum(tmp7)
	tmp8, err := this._io.ReadBitsInt(4)
	if err != nil {
		return err
	}
	this.ContinuityCounter = tmp8
	this._io.AlignToByte()
	tmp9, err := this._io.ReadBytes(int(184))
	if err != nil {
		return err
	}
	this.TsPacketRemain = tmp9
	return err
}