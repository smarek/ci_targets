// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

(function (root, factory) {
  if (typeof define === 'function' && define.amd) {
    define(['kaitai-struct/KaitaiStream'], factory);
  } else if (typeof module === 'object' && module.exports) {
    module.exports = factory(require('kaitai-struct/KaitaiStream'));
  } else {
    root.ValidFailRangeStr = factory(root.KaitaiStream);
  }
}(this, function (KaitaiStream) {
var ValidFailRangeStr = (function() {
  function ValidFailRangeStr(_io, _parent, _root) {
    this._io = _io;
    this._parent = _parent;
    this._root = _root || this;

    this._read();
  }
  ValidFailRangeStr.prototype._read = function() {
    this.foo = KaitaiStream.bytesToStr(this._io.readBytes(2), "ASCII");
    if (!(this.foo >= "H@")) {
      throw new KaitaiStream.ValidationLessThanError("H@", this.foo, this._io, "/seq/0");
    }
    if (!(this.foo <= "O~")) {
      throw new KaitaiStream.ValidationGreaterThanError("O~", this.foo, this._io, "/seq/0");
    }
  }

  return ValidFailRangeStr;
})();
return ValidFailRangeStr;
}));
