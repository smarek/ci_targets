# This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

from pkg_resources import parse_version
import kaitaistruct
from kaitaistruct import KaitaiStruct, KaitaiStream, BytesIO


if parse_version(kaitaistruct.__version__) < parse_version('0.9'):
    raise Exception("Incompatible Kaitai Struct Python API: 0.9 or later is required, but you have %s" % (kaitaistruct.__version__))

class CastNested(KaitaiStruct):
    def __init__(self, _io, _parent=None, _root=None):
        self._io = _io
        self._parent = _parent
        self._root = _root if _root else self
        self._read()

    def _read(self):
        self.opcodes = []
        i = 0
        while not self._io.is_eof():
            self.opcodes.append(CastNested.Opcode(self._io, self, self._root))
            i += 1


    class Opcode(KaitaiStruct):
        def __init__(self, _io, _parent=None, _root=None):
            self._io = _io
            self._parent = _parent
            self._root = _root if _root else self
            self._read()

        def _read(self):
            self.code = self._io.read_u1()
            _on = self.code
            if _on == 73:
                self.body = CastNested.Opcode.Intval(self._io, self, self._root)
            elif _on == 83:
                self.body = CastNested.Opcode.Strval(self._io, self, self._root)

        class Intval(KaitaiStruct):
            def __init__(self, _io, _parent=None, _root=None):
                self._io = _io
                self._parent = _parent
                self._root = _root if _root else self
                self._read()

            def _read(self):
                self.value = self._io.read_u1()


        class Strval(KaitaiStruct):
            def __init__(self, _io, _parent=None, _root=None):
                self._io = _io
                self._parent = _parent
                self._root = _root if _root else self
                self._read()

            def _read(self):
                self.value = (self._io.read_bytes_term(0, False, True, True)).decode(u"ASCII")



    @property
    def opcodes_0_str(self):
        if hasattr(self, '_m_opcodes_0_str'):
            return self._m_opcodes_0_str if hasattr(self, '_m_opcodes_0_str') else None

        self._m_opcodes_0_str = self.opcodes[0].body
        return self._m_opcodes_0_str if hasattr(self, '_m_opcodes_0_str') else None

    @property
    def opcodes_0_str_value(self):
        if hasattr(self, '_m_opcodes_0_str_value'):
            return self._m_opcodes_0_str_value if hasattr(self, '_m_opcodes_0_str_value') else None

        self._m_opcodes_0_str_value = self.opcodes[0].body.value
        return self._m_opcodes_0_str_value if hasattr(self, '_m_opcodes_0_str_value') else None

    @property
    def opcodes_1_int(self):
        if hasattr(self, '_m_opcodes_1_int'):
            return self._m_opcodes_1_int if hasattr(self, '_m_opcodes_1_int') else None

        self._m_opcodes_1_int = self.opcodes[1].body
        return self._m_opcodes_1_int if hasattr(self, '_m_opcodes_1_int') else None

    @property
    def opcodes_1_int_value(self):
        if hasattr(self, '_m_opcodes_1_int_value'):
            return self._m_opcodes_1_int_value if hasattr(self, '_m_opcodes_1_int_value') else None

        self._m_opcodes_1_int_value = self.opcodes[1].body.value
        return self._m_opcodes_1_int_value if hasattr(self, '_m_opcodes_1_int_value') else None


