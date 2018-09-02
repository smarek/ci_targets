// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package test_formats

import "github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"


type DebugEnumName_TestEnum1 int
const (
	DebugEnumName_TestEnum1__EnumValue80 DebugEnumName_TestEnum1 = 80
)

type DebugEnumName_TestEnum2 int
const (
	DebugEnumName_TestEnum2__EnumValue65 DebugEnumName_TestEnum2 = 65
)
type DebugEnumName struct {
	One DebugEnumName_TestEnum1
	ArrayOfInts []DebugEnumName_TestEnum2
	TestType *DebugEnumName_TestSubtype
	_io *kaitai.Stream
	_root *DebugEnumName
	_parent interface{}
}

func (this *DebugEnumName) Read(io *kaitai.Stream, parent interface{}, root *DebugEnumName) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp1, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.One = DebugEnumName_TestEnum1(tmp1)
	this.ArrayOfInts = make([]DebugEnumName_TestEnum2, 1)
	for i := range this.ArrayOfInts {
		tmp2, err := this._io.ReadU1()
		if err != nil {
			return err
		}
		this.ArrayOfInts[i] = DebugEnumName_TestEnum2(tmp2)
	}
	tmp3 := new(DebugEnumName_TestSubtype)
	err = tmp3.Read(this._io, this, this._root)
	if err != nil {
		return err
	}
	this.TestType = tmp3
	return err
}

type DebugEnumName_TestSubtype_InnerEnum1 int
const (
	DebugEnumName_TestSubtype_InnerEnum1__EnumValue67 DebugEnumName_TestSubtype_InnerEnum1 = 67
)

type DebugEnumName_TestSubtype_InnerEnum2 int
const (
	DebugEnumName_TestSubtype_InnerEnum2__EnumValue11 DebugEnumName_TestSubtype_InnerEnum2 = 11
)
type DebugEnumName_TestSubtype struct {
	Field1 DebugEnumName_TestSubtype_InnerEnum1
	Field2 uint8
	_io *kaitai.Stream
	_root *DebugEnumName
	_parent *DebugEnumName
	_f_instanceField bool
	instanceField DebugEnumName_TestSubtype_InnerEnum2
}

func (this *DebugEnumName_TestSubtype) Read(io *kaitai.Stream, parent *DebugEnumName, root *DebugEnumName) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp4, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.Field1 = DebugEnumName_TestSubtype_InnerEnum1(tmp4)
	tmp5, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.Field2 = tmp5
	return err
}
func (this *DebugEnumName_TestSubtype) InstanceField() (v DebugEnumName_TestSubtype_InnerEnum2, err error) {
	if (this._f_instanceField) {
		return this.instanceField, nil
	}
	this.instanceField = DebugEnumName_TestSubtype_InnerEnum2(DebugEnumName_TestSubtype_InnerEnum2((this.Field2 & 15)))
	this._f_instanceField = true
	return this.instanceField, nil
}