// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

#include "bits_enum.h"



bits_enum_t::bits_enum_t(kaitai::kstream *p_io, kaitai::kstruct* p_parent, bits_enum_t *p_root) : kaitai::kstruct(p_io) {
    m__parent = p_parent;
    m__root = this;
    _read();
}

void bits_enum_t::_read() {
    m_one = static_cast<bits_enum_t::animal_t>(m__io->read_bits_int(4));
    m_two = static_cast<bits_enum_t::animal_t>(m__io->read_bits_int(8));
    m_three = static_cast<bits_enum_t::animal_t>(m__io->read_bits_int(1));
}

bits_enum_t::~bits_enum_t() {
}
