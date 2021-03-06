// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package test_formats

import "github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"

type ZlibWithHeader78 struct {
	Data []byte
	_io *kaitai.Stream
	_root *ZlibWithHeader78
	_parent interface{}
	_raw_Data []byte
}
func NewZlibWithHeader78() *ZlibWithHeader78 {
	return &ZlibWithHeader78{
	}
}

func (this *ZlibWithHeader78) Read(io *kaitai.Stream, parent interface{}, root *ZlibWithHeader78) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp1, err := this._io.ReadBytesFull()
	if err != nil {
		return err
	}
	tmp1 = tmp1
	this._raw_Data = tmp1
	tmp2, err := kaitai.ProcessZlib(this._raw_Data)
	if err != nil {
		return err
	}
	this.Data = tmp2
	return err
}
