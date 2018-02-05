-- This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild
--
-- This file is compatible with Lua 5.3

local class = require("class")
require("kaitaistruct")
local str_decode = require("string_decode")

StrEncodingsDefault = class.class(KaitaiStruct)

function StrEncodingsDefault:_init(io, parent, root)
  KaitaiStruct._init(self, io)
  self._parent = parent
  self._root = root or self
  self:_read()
end

function StrEncodingsDefault:_read()
  self.len_of_1 = self._io:read_u2le()
  self.str1 = str_decode.decode(self._io:read_bytes(self.len_of_1), "UTF-8")
  self.rest = StrEncodingsDefault.Subtype(self._io, self, self._root)
end


StrEncodingsDefault.Subtype = class.class(KaitaiStruct)

function StrEncodingsDefault.Subtype:_init(io, parent, root)
  KaitaiStruct._init(self, io)
  self._parent = parent
  self._root = root or self
  self:_read()
end

function StrEncodingsDefault.Subtype:_read()
  self.len_of_2 = self._io:read_u2le()
  self.str2 = str_decode.decode(self._io:read_bytes(self.len_of_2), "UTF-8")
  self.len_of_3 = self._io:read_u2le()
  self.str3 = str_decode.decode(self._io:read_bytes(self.len_of_3), "SJIS")
  self.len_of_4 = self._io:read_u2le()
  self.str4 = str_decode.decode(self._io:read_bytes(self.len_of_4), "CP437")
end


