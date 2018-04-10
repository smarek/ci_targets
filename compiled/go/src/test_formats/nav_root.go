// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package test_formats

import "github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
type NavRoot struct {
	Header *NavRoot_HeaderObj
	Index *NavRoot_IndexObj
	_io *kaitai.Stream
	_root *NavRoot
	_parent interface{}
}

func (this *NavRoot) Read(io *kaitai.Stream, parent interface{}, root *NavRoot) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp1 := new(NavRoot_HeaderObj)
	err = tmp1.Read(this._io, this, this._root)
	if err != nil {
		return err
	}
	this.Header = tmp1
	tmp2 := new(NavRoot_IndexObj)
	err = tmp2.Read(this._io, this, this._root)
	if err != nil {
		return err
	}
	this.Index = tmp2
	return err
}
type NavRoot_HeaderObj struct {
	QtyEntries uint32
	FilenameLen uint32
	_io *kaitai.Stream
	_root *NavRoot
	_parent *NavRoot
}

func (this *NavRoot_HeaderObj) Read(io *kaitai.Stream, parent *NavRoot, root *NavRoot) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp3, err := this._io.ReadU4le()
	if err != nil {
		return err
	}
	this.QtyEntries = tmp3
	tmp4, err := this._io.ReadU4le()
	if err != nil {
		return err
	}
	this.FilenameLen = tmp4
	return err
}
type NavRoot_IndexObj struct {
	Magic []byte
	Entries []*NavRoot_Entry
	_io *kaitai.Stream
	_root *NavRoot
	_parent *NavRoot
}

func (this *NavRoot_IndexObj) Read(io *kaitai.Stream, parent *NavRoot, root *NavRoot) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp5, err := this._io.ReadBytes(int(4))
	if err != nil {
		return err
	}
	this.Magic = tmp5
	this.Entries = make([]*NavRoot_Entry, this._root.Header.QtyEntries)
	for i := range this.Entries {
		tmp6 := new(NavRoot_Entry)
		err = tmp6.Read(this._io, this, this._root)
		if err != nil {
			return err
		}
		this.Entries[i] = tmp6
	}
	return err
}
type NavRoot_Entry struct {
	Filename string
	_io *kaitai.Stream
	_root *NavRoot
	_parent *NavRoot_IndexObj
}

func (this *NavRoot_Entry) Read(io *kaitai.Stream, parent *NavRoot_IndexObj, root *NavRoot) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp7, err := this._io.ReadBytes(int(this._root.Header.FilenameLen))
	if err != nil {
		return err
	}
	this.Filename = string(tmp7)
	return err
}