// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package test_formats

import "github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"

type Imported1 struct {
	One uint8
	Two *Imported2
	_io *kaitai.Stream
	_root *Imported1
	_parent interface{}
}
func NewImported1() *Imported1 {
	return &Imported1{
	}
}

func (this *Imported1) Read(io *kaitai.Stream, parent interface{}, root *Imported1) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp1, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.One = tmp1
	tmp2 := NewImported2()
	err = tmp2.Read(this._io, this, nil)
	if err != nil {
		return err
	}
	this.Two = tmp2
	return err
}
