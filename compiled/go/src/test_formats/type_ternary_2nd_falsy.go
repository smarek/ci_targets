// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package test_formats

import "github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"

type TypeTernary2ndFalsy struct {
	IntTruthy uint8
	Ut *TypeTernary2ndFalsy_Foo
	IntArray []uint8
	IntArrayEmpty []uint8
	_io *kaitai.Stream
	_root *TypeTernary2ndFalsy
	_parent interface{}
	_f_nullUt bool
	nullUt *TypeTernary2ndFalsy_Foo
	_f_vFloatZero bool
	vFloatZero float64
	_f_t bool
	t bool
	_f_vIntNegZero bool
	vIntNegZero int
	_f_vIntZero bool
	vIntZero int8
	_f_vFalse bool
	vFalse bool
	_f_vStrEmpty bool
	vStrEmpty string
	_f_vIntArrayEmpty bool
	vIntArrayEmpty []uint8
	_f_vNullUt bool
	vNullUt *TypeTernary2ndFalsy_Foo
	_f_vFloatNegZero bool
	vFloatNegZero float64
	_f_vStrWZero bool
	vStrWZero string
}
func NewTypeTernary2ndFalsy() *TypeTernary2ndFalsy {
	return &TypeTernary2ndFalsy{
	}
}

func (this *TypeTernary2ndFalsy) Read(io *kaitai.Stream, parent interface{}, root *TypeTernary2ndFalsy) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp1, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.IntTruthy = tmp1
	tmp2 := NewTypeTernary2ndFalsy_Foo()
	err = tmp2.Read(this._io, this, this._root)
	if err != nil {
		return err
	}
	this.Ut = tmp2
	this.IntArray = make([]uint8, 2)
	for i := range this.IntArray {
		tmp3, err := this._io.ReadU1()
		if err != nil {
			return err
		}
		this.IntArray[i] = tmp3
	}
	this.IntArrayEmpty = make([]uint8, 0)
	for i := range this.IntArrayEmpty {
		tmp4, err := this._io.ReadU1()
		if err != nil {
			return err
		}
		this.IntArrayEmpty[i] = tmp4
	}
	return err
}
func (this *TypeTernary2ndFalsy) NullUt() (v *TypeTernary2ndFalsy_Foo, err error) {
	if (this._f_nullUt) {
		return this.nullUt, nil
	}
	if (false) {
		this.nullUt = this.Ut
	}
	this._f_nullUt = true
	return this.nullUt, nil
}
func (this *TypeTernary2ndFalsy) VFloatZero() (v float64, err error) {
	if (this._f_vFloatZero) {
		return this.vFloatZero, nil
	}
	var tmp5 float64;
	tmp6, err := this.T()
	if err != nil {
		return 0, err
	}
	if (tmp6) {
		tmp5 = 0.0
	} else {
		tmp5 = 3.14
	}
	this.vFloatZero = float64(tmp5)
	this._f_vFloatZero = true
	return this.vFloatZero, nil
}
func (this *TypeTernary2ndFalsy) T() (v bool, err error) {
	if (this._f_t) {
		return this.t, nil
	}
	this.t = bool(true)
	this._f_t = true
	return this.t, nil
}
func (this *TypeTernary2ndFalsy) VIntNegZero() (v int, err error) {
	if (this._f_vIntNegZero) {
		return this.vIntNegZero, nil
	}
	var tmp7 int;
	tmp8, err := this.T()
	if err != nil {
		return 0, err
	}
	if (tmp8) {
		tmp7 = -0
	} else {
		tmp7 = -20
	}
	this.vIntNegZero = int(tmp7)
	this._f_vIntNegZero = true
	return this.vIntNegZero, nil
}
func (this *TypeTernary2ndFalsy) VIntZero() (v int8, err error) {
	if (this._f_vIntZero) {
		return this.vIntZero, nil
	}
	var tmp9 int8;
	tmp10, err := this.T()
	if err != nil {
		return 0, err
	}
	if (tmp10) {
		tmp9 = 0
	} else {
		tmp9 = 10
	}
	this.vIntZero = int8(tmp9)
	this._f_vIntZero = true
	return this.vIntZero, nil
}
func (this *TypeTernary2ndFalsy) VFalse() (v bool, err error) {
	if (this._f_vFalse) {
		return this.vFalse, nil
	}
	var tmp11 bool;
	tmp12, err := this.T()
	if err != nil {
		return false, err
	}
	if (tmp12) {
		tmp11 = false
	} else {
		tmp11 = true
	}
	this.vFalse = bool(tmp11)
	this._f_vFalse = true
	return this.vFalse, nil
}
func (this *TypeTernary2ndFalsy) VStrEmpty() (v string, err error) {
	if (this._f_vStrEmpty) {
		return this.vStrEmpty, nil
	}
	var tmp13 string;
	tmp14, err := this.T()
	if err != nil {
		return "", err
	}
	if (tmp14) {
		tmp13 = ""
	} else {
		tmp13 = "kaitai"
	}
	this.vStrEmpty = string(tmp13)
	this._f_vStrEmpty = true
	return this.vStrEmpty, nil
}
func (this *TypeTernary2ndFalsy) VIntArrayEmpty() (v []uint8, err error) {
	if (this._f_vIntArrayEmpty) {
		return this.vIntArrayEmpty, nil
	}
	var tmp15 []uint8;
	tmp16, err := this.T()
	if err != nil {
		return nil, err
	}
	if (tmp16) {
		tmp15 = this.IntArrayEmpty
	} else {
		tmp15 = this.IntArray
	}
	this.vIntArrayEmpty = []uint8(tmp15)
	this._f_vIntArrayEmpty = true
	return this.vIntArrayEmpty, nil
}
func (this *TypeTernary2ndFalsy) VNullUt() (v *TypeTernary2ndFalsy_Foo, err error) {
	if (this._f_vNullUt) {
		return this.vNullUt, nil
	}
	var tmp17 *TypeTernary2ndFalsy_Foo;
	tmp18, err := this.T()
	if err != nil {
		return nil, err
	}
	if (tmp18) {
		tmp19, err := this.NullUt()
		if err != nil {
			return nil, err
		}
		tmp17 = tmp19
	} else {
		tmp17 = this.Ut
	}
	this.vNullUt = tmp17
	this._f_vNullUt = true
	return this.vNullUt, nil
}
func (this *TypeTernary2ndFalsy) VFloatNegZero() (v float64, err error) {
	if (this._f_vFloatNegZero) {
		return this.vFloatNegZero, nil
	}
	var tmp20 float64;
	tmp21, err := this.T()
	if err != nil {
		return 0, err
	}
	if (tmp21) {
		tmp20 = -0.0
	} else {
		tmp20 = -2.72
	}
	this.vFloatNegZero = float64(tmp20)
	this._f_vFloatNegZero = true
	return this.vFloatNegZero, nil
}
func (this *TypeTernary2ndFalsy) VStrWZero() (v string, err error) {
	if (this._f_vStrWZero) {
		return this.vStrWZero, nil
	}
	var tmp22 string;
	tmp23, err := this.T()
	if err != nil {
		return "", err
	}
	if (tmp23) {
		tmp22 = "0"
	} else {
		tmp22 = "30"
	}
	this.vStrWZero = string(tmp22)
	this._f_vStrWZero = true
	return this.vStrWZero, nil
}
type TypeTernary2ndFalsy_Foo struct {
	M uint8
	_io *kaitai.Stream
	_root *TypeTernary2ndFalsy
	_parent *TypeTernary2ndFalsy
}
func NewTypeTernary2ndFalsy_Foo() *TypeTernary2ndFalsy_Foo {
	return &TypeTernary2ndFalsy_Foo{
	}
}

func (this *TypeTernary2ndFalsy_Foo) Read(io *kaitai.Stream, parent *TypeTernary2ndFalsy, root *TypeTernary2ndFalsy) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp24, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.M = tmp24
	return err
}