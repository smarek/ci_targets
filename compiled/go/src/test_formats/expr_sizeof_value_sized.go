// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package test_formats

import (
	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
	"bytes"
)

type ExprSizeofValueSized struct {
	Block1 *ExprSizeofValueSized_Block
	More uint16
	_io *kaitai.Stream
	_root *ExprSizeofValueSized
	_parent interface{}
	_raw_Block1 []byte
	_f_selfSizeof bool
	selfSizeof int
	_f_sizeofBlock bool
	sizeofBlock int
	_f_sizeofBlockB bool
	sizeofBlockB int
	_f_sizeofBlockA bool
	sizeofBlockA int
	_f_sizeofBlockC bool
	sizeofBlockC int
}

func (this *ExprSizeofValueSized) Read(io *kaitai.Stream, parent interface{}, root *ExprSizeofValueSized) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp1, err := this._io.ReadBytes(int(12))
	if err != nil {
		return err
	}
	tmp1 = tmp1
	this._raw_Block1 = tmp1
	_io__raw_Block1 := kaitai.NewStream(bytes.NewReader(this._raw_Block1))
	tmp2 := new(ExprSizeofValueSized_Block)
	err = tmp2.Read(_io__raw_Block1, this, this._root)
	if err != nil {
		return err
	}
	this.Block1 = tmp2
	tmp3, err := this._io.ReadU2le()
	if err != nil {
		return err
	}
	this.More = uint16(tmp3)
	return err
}
func (this *ExprSizeofValueSized) SelfSizeof() (v int, err error) {
	if (this._f_selfSizeof) {
		return this.selfSizeof, nil
	}
	this.selfSizeof = int(14)
	this._f_selfSizeof = true
	return this.selfSizeof, nil
}
func (this *ExprSizeofValueSized) SizeofBlock() (v int, err error) {
	if (this._f_sizeofBlock) {
		return this.sizeofBlock, nil
	}
	this.sizeofBlock = int(12)
	this._f_sizeofBlock = true
	return this.sizeofBlock, nil
}
func (this *ExprSizeofValueSized) SizeofBlockB() (v int, err error) {
	if (this._f_sizeofBlockB) {
		return this.sizeofBlockB, nil
	}
	this.sizeofBlockB = int(4)
	this._f_sizeofBlockB = true
	return this.sizeofBlockB, nil
}
func (this *ExprSizeofValueSized) SizeofBlockA() (v int, err error) {
	if (this._f_sizeofBlockA) {
		return this.sizeofBlockA, nil
	}
	this.sizeofBlockA = int(1)
	this._f_sizeofBlockA = true
	return this.sizeofBlockA, nil
}
func (this *ExprSizeofValueSized) SizeofBlockC() (v int, err error) {
	if (this._f_sizeofBlockC) {
		return this.sizeofBlockC, nil
	}
	this.sizeofBlockC = int(2)
	this._f_sizeofBlockC = true
	return this.sizeofBlockC, nil
}
type ExprSizeofValueSized_Block struct {
	A uint8
	B uint32
	C []byte
	_io *kaitai.Stream
	_root *ExprSizeofValueSized
	_parent *ExprSizeofValueSized
}

func (this *ExprSizeofValueSized_Block) Read(io *kaitai.Stream, parent *ExprSizeofValueSized, root *ExprSizeofValueSized) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp4, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.A = tmp4
	tmp5, err := this._io.ReadU4le()
	if err != nil {
		return err
	}
	this.B = uint32(tmp5)
	tmp6, err := this._io.ReadBytes(int(2))
	if err != nil {
		return err
	}
	tmp6 = tmp6
	this.C = tmp6
	return err
}