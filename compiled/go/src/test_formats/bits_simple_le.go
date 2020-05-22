// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package test_formats

import "github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"

type BitsSimpleLe struct {
	Byte1 uint64
	Byte2 uint64
	BitsA bool
	BitsB uint64
	BitsC uint64
	LargeBits1 uint64
	Spacer uint64
	LargeBits2 uint64
	NormalS2 int16
	Byte8910 uint64
	Byte11To14 uint64
	Byte15To19 uint64
	Byte20To27 uint64
	_io *kaitai.Stream
	_root *BitsSimpleLe
	_parent interface{}
	_f_testIfB1 bool
	testIfB1 int8
}
func NewBitsSimpleLe() *BitsSimpleLe {
	return &BitsSimpleLe{
	}
}

func (this *BitsSimpleLe) Read(io *kaitai.Stream, parent interface{}, root *BitsSimpleLe) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp1, err := this._io.ReadBitsIntLe(8)
	if err != nil {
		return err
	}
	this.Byte1 = tmp1
	tmp2, err := this._io.ReadBitsIntLe(8)
	if err != nil {
		return err
	}
	this.Byte2 = tmp2
	tmp3, err := this._io.ReadBitsIntLe(1)
	if err != nil {
		return err
	}
	this.BitsA = tmp3 != 0
	tmp4, err := this._io.ReadBitsIntLe(3)
	if err != nil {
		return err
	}
	this.BitsB = tmp4
	tmp5, err := this._io.ReadBitsIntLe(4)
	if err != nil {
		return err
	}
	this.BitsC = tmp5
	tmp6, err := this._io.ReadBitsIntLe(10)
	if err != nil {
		return err
	}
	this.LargeBits1 = tmp6
	tmp7, err := this._io.ReadBitsIntLe(3)
	if err != nil {
		return err
	}
	this.Spacer = tmp7
	tmp8, err := this._io.ReadBitsIntLe(11)
	if err != nil {
		return err
	}
	this.LargeBits2 = tmp8
	this._io.AlignToByte()
	tmp9, err := this._io.ReadS2be()
	if err != nil {
		return err
	}
	this.NormalS2 = int16(tmp9)
	tmp10, err := this._io.ReadBitsIntLe(24)
	if err != nil {
		return err
	}
	this.Byte8910 = tmp10
	tmp11, err := this._io.ReadBitsIntLe(32)
	if err != nil {
		return err
	}
	this.Byte11To14 = tmp11
	tmp12, err := this._io.ReadBitsIntLe(40)
	if err != nil {
		return err
	}
	this.Byte15To19 = tmp12
	tmp13, err := this._io.ReadBitsIntLe(64)
	if err != nil {
		return err
	}
	this.Byte20To27 = tmp13
	return err
}
func (this *BitsSimpleLe) TestIfB1() (v int8, err error) {
	if (this._f_testIfB1) {
		return this.testIfB1, nil
	}
	if (this.BitsA == true) {
		this.testIfB1 = int8(123)
	}
	this._f_testIfB1 = true
	return this.testIfB1, nil
}