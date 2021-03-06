# This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

from pkg_resources import parse_version
import kaitaistruct
from kaitaistruct import KaitaiStruct, KaitaiStream, BytesIO


if parse_version(kaitaistruct.__version__) < parse_version('0.9'):
    raise Exception("Incompatible Kaitai Struct Python API: 0.9 or later is required, but you have %s" % (kaitaistruct.__version__))

class EnumToIClassBorder2(KaitaiStruct):
    def __init__(self, parent, _io, _parent=None, _root=None):
        self._io = _io
        self._parent = _parent
        self._root = _root if _root else self
        self.parent = parent
        self._read()

    def _read(self):
        pass

    @property
    def is_dog(self):
        if hasattr(self, '_m_is_dog'):
            return self._m_is_dog if hasattr(self, '_m_is_dog') else None

        self._m_is_dog = self.parent.some_dog.value == 4
        return self._m_is_dog if hasattr(self, '_m_is_dog') else None


