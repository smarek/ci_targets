# This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

from pkg_resources import parse_version
import kaitaistruct
from kaitaistruct import KaitaiStruct, KaitaiStream, BytesIO


if parse_version(kaitaistruct.__version__) < parse_version('0.9'):
    raise Exception("Incompatible Kaitai Struct Python API: 0.9 or later is required, but you have %s" % (kaitaistruct.__version__))

class ProcessCoerceBytes(KaitaiStruct):
    def __init__(self, _io, _parent=None, _root=None):
        self._io = _io
        self._parent = _parent
        self._root = _root if _root else self
        self._read()

    def _read(self):
        self.records = [None] * (2)
        for i in range(2):
            self.records[i] = ProcessCoerceBytes.Record(self._io, self, self._root)


    class Record(KaitaiStruct):
        def __init__(self, _io, _parent=None, _root=None):
            self._io = _io
            self._parent = _parent
            self._root = _root if _root else self
            self._read()

        def _read(self):
            self.flag = self._io.read_u1()
            if self.flag == 0:
                self.buf_unproc = self._io.read_bytes(4)

            if self.flag != 0:
                self._raw_buf_proc = self._io.read_bytes(4)
                self.buf_proc = KaitaiStream.process_xor_one(self._raw_buf_proc, 170)


        @property
        def buf(self):
            if hasattr(self, '_m_buf'):
                return self._m_buf if hasattr(self, '_m_buf') else None

            self._m_buf = (self.buf_unproc if self.flag == 0 else self.buf_proc)
            return self._m_buf if hasattr(self, '_m_buf') else None



