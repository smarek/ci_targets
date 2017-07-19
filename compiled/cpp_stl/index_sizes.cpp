// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

#include "index_sizes.h"



index_sizes_t::index_sizes_t(kaitai::kstream *p_io, kaitai::kstruct* p_parent, index_sizes_t *p_root) : kaitai::kstruct(p_io) {
    m__parent = p_parent;
    m__root = this;
    _read();
}

void index_sizes_t::_read() {
    m_qty = m__io->read_u4le();
    int l_sizes = qty();
    m_sizes = new std::vector<uint32_t>();
    m_sizes->reserve(l_sizes);
    for (int i = 0; i < l_sizes; i++) {
        m_sizes->push_back(m__io->read_u4le());
    }
    int l_bufs = qty();
    m_bufs = new std::vector<std::string>();
    m_bufs->reserve(l_bufs);
    for (int i = 0; i < l_bufs; i++) {
        m_bufs->push_back(kaitai::kstream::bytes_to_str(m__io->read_bytes(sizes()->at(i)), std::string("ASCII")));
    }
}

index_sizes_t::~index_sizes_t() {
    delete m_sizes;
    delete m_bufs;
}
