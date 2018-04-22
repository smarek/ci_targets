<?php
// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

namespace Kaitai\Struct\Tests;

class FixedContents extends \Kaitai\Struct\Struct {
    public function __construct(\Kaitai\Struct\Stream $_io, \Kaitai\Struct\Struct $_parent = null, \Kaitai\Struct\Tests\FixedContents $_root = null) {
        parent::__construct($_io, $_parent, $_root);
        $this->_read();
    }

    private function _read() {
        $this->_m_normal = $this->_io->ensureFixedContents("\x50\x41\x43\x4B\x2D\x31");
        $this->_m_highBit8 = $this->_io->ensureFixedContents("\xFF\xFF");
    }
    protected $_m_normal;
    protected $_m_highBit8;
    public function normal() { return $this->_m_normal; }
    public function highBit8() { return $this->_m_highBit8; }
}