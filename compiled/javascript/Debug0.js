// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

(function (root, factory) {
  if (typeof define === 'function' && define.amd) {
    define(['kaitai-struct/KaitaiStream'], factory);
  } else if (typeof module === 'object' && module.exports) {
    module.exports = factory(require('kaitai-struct/KaitaiStream'));
  } else {
    root.Debug0 = factory(root.KaitaiStream);
  }
}(this, function (KaitaiStream) {
var Debug0 = (function() {
  function Debug0(_io, _parent, _root) {
    this._io = _io;
    this._parent = _parent;
    this._root = _root || this;
    this._debug = {};

  }
  Debug0.prototype._read = function() {
    this._debug.one = { start: this._io.pos, ioOffset: this._io._byteOffset };
    this.one = this._io.readU1();
    this._debug.one.end = this._io.pos;
    this._debug.arrayOfInts = { start: this._io.pos, ioOffset: this._io._byteOffset };
    this.arrayOfInts = new Array(3);
    this._debug.arrayOfInts.arr = new Array(3);
    for (var i = 0; i < 3; i++) {
      this._debug.arrayOfInts.arr[i] = { start: this._io.pos, ioOffset: this._io._byteOffset };
      this.arrayOfInts[i] = this._io.readU1();
      this._debug.arrayOfInts.arr[i].end = this._io.pos;
    }
    this._debug.arrayOfInts.end = this._io.pos;
    this._debug._unnamed2 = { start: this._io.pos, ioOffset: this._io._byteOffset };
    this._unnamed2 = this._io.readU1();
    this._debug._unnamed2.end = this._io.pos;
  }

  return Debug0;
})();
return Debug0;
}));