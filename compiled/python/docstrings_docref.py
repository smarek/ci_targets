from pkg_resources import parse_version
from kaitaistruct import __version__ as ks_version, KaitaiStruct, KaitaiStream, BytesIO
# This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild



if parse_version(ks_version) < parse_version('0.7'):
    raise Exception("Incompatible Kaitai Struct Python API: 0.7 or later is required, but you have %s" % (ks_version))

class DocstringsDocref(KaitaiStruct):
    def __init__(self, _io, _parent=None, _root=None):
        self._io = _io
        self._parent = _parent
        self._root = _root if _root else self
        self.one = self._io.read_u1()
        self.two = self._io.read_u1()
        self.three = self._io.read_u1()

    @property
    def foo(self):
        if hasattr(self, '_m_foo'):
            return self._m_foo if hasattr(self, '_m_foo') else None

        self._m_foo = True
        return self._m_foo if hasattr(self, '_m_foo') else None

    @property
    def parse_inst(self):
        if hasattr(self, '_m_parse_inst'):
            return self._m_parse_inst if hasattr(self, '_m_parse_inst') else None

        _pos = self._io.pos()
        self._io.seek(0)
        self._m_parse_inst = self._io.read_u1()
        self._io.seek(_pos)
        return self._m_parse_inst if hasattr(self, '_m_parse_inst') else None


