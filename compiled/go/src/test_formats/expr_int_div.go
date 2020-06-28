// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package test_formats

import "github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"

type ExprIntDiv struct {
	IntU uint32
	IntS int32
	_io *kaitai.Stream
	_root *ExprIntDiv
	_parent interface{}
	_f_divPosConst bool
	divPosConst int
	_f_divNegConst bool
	divNegConst int
	_f_divPosSeq bool
	divPosSeq int
	_f_divNegSeq bool
	divNegSeq int
}
func NewExprIntDiv() *ExprIntDiv {
	return &ExprIntDiv{
	}
}

func (this *ExprIntDiv) Read(io *kaitai.Stream, parent interface{}, root *ExprIntDiv) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp1, err := this._io.ReadU4le()
	if err != nil {
		return err
	}
	this.IntU = uint32(tmp1)
	tmp2, err := this._io.ReadS4le()
	if err != nil {
		return err
	}
	this.IntS = int32(tmp2)
	return err
}
func (this *ExprIntDiv) DivPosConst() (v int, err error) {
	if (this._f_divPosConst) {
		return this.divPosConst, nil
	}
	this.divPosConst = int((9837 / 13))
	this._f_divPosConst = true
	return this.divPosConst, nil
}
func (this *ExprIntDiv) DivNegConst() (v int, err error) {
	if (this._f_divNegConst) {
		return this.divNegConst, nil
	}
	this.divNegConst = int((-9837 / 13))
	this._f_divNegConst = true
	return this.divNegConst, nil
}
func (this *ExprIntDiv) DivPosSeq() (v int, err error) {
	if (this._f_divPosSeq) {
		return this.divPosSeq, nil
	}
	this.divPosSeq = int((this.IntU / 13))
	this._f_divPosSeq = true
	return this.divPosSeq, nil
}
func (this *ExprIntDiv) DivNegSeq() (v int, err error) {
	if (this._f_divNegSeq) {
		return this.divNegSeq, nil
	}
	this.divNegSeq = int((this.IntS / 13))
	this._f_divNegSeq = true
	return this.divNegSeq, nil
}