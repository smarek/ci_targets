// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package test_formats

import "github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"

type ParamsPassArrayInt struct {
	Ints []uint16
	PassInts *ParamsPassArrayInt_WantsInts
	PassIntsCalc *ParamsPassArrayInt_WantsInts
	_io *kaitai.Stream
	_root *ParamsPassArrayInt
	_parent interface{}
	_f_intsCalc bool
	intsCalc []int
}
func NewParamsPassArrayInt() *ParamsPassArrayInt {
	return &ParamsPassArrayInt{
	}
}

func (this *ParamsPassArrayInt) Read(io *kaitai.Stream, parent interface{}, root *ParamsPassArrayInt) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	this.Ints = make([]uint16, 3)
	for i := range this.Ints {
		tmp1, err := this._io.ReadU2le()
		if err != nil {
			return err
		}
		this.Ints[i] = tmp1
	}
	tmp2 := NewParamsPassArrayInt_WantsInts(this.Ints)
	err = tmp2.Read(this._io, this, this._root)
	if err != nil {
		return err
	}
	this.PassInts = tmp2
	tmp4, err := this.IntsCalc()
	if err != nil {
		return err
	}
	tmp3 := NewParamsPassArrayInt_WantsInts(tmp4)
	err = tmp3.Read(this._io, this, this._root)
	if err != nil {
		return err
	}
	this.PassIntsCalc = tmp3
	return err
}
func (this *ParamsPassArrayInt) IntsCalc() (v []int, err error) {
	if (this._f_intsCalc) {
		return this.intsCalc, nil
	}
	this.intsCalc = []int([]int{27643, 7})
	this._f_intsCalc = true
	return this.intsCalc, nil
}
type ParamsPassArrayInt_WantsInts struct {
	Nums []uint16
	_io *kaitai.Stream
	_root *ParamsPassArrayInt
	_parent *ParamsPassArrayInt
}
func NewParamsPassArrayInt_WantsInts(nums []uint16) *ParamsPassArrayInt_WantsInts {
	return &ParamsPassArrayInt_WantsInts{
		Nums: nums,
	}
}

func (this *ParamsPassArrayInt_WantsInts) Read(io *kaitai.Stream, parent *ParamsPassArrayInt, root *ParamsPassArrayInt) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	return err
}
