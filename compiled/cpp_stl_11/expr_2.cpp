// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

#include <memory>
#include "expr_2.h"

expr_2_t::expr_2_t(kaitai::kstream* p__io, kaitai::kstruct* p__parent, expr_2_t* p__root) : kaitai::kstruct(p__io) {
    m__parent = p__parent;
    m__root = this;
    m_str1 = nullptr;
    m_str2 = nullptr;
    f_str1_len_mod = false;
    f_str1_len = false;
    f_str1_tuple5 = false;
    f_str2_tuple5 = false;
    f_str1_avg = false;
    f_str1_byte1 = false;
    f_str1_char5 = false;
    _read();
}

void expr_2_t::_read() {
    m_str1 = std::unique_ptr<mod_str_t>(new mod_str_t(m__io, this, m__root));
    m_str2 = std::unique_ptr<mod_str_t>(new mod_str_t(m__io, this, m__root));
}

expr_2_t::~expr_2_t() {
    _clean_up();
}

void expr_2_t::_clean_up() {
}

expr_2_t::mod_str_t::mod_str_t(kaitai::kstream* p__io, expr_2_t* p__parent, expr_2_t* p__root) : kaitai::kstruct(p__io) {
    m__parent = p__parent;
    m__root = p__root;
    m_rest = nullptr;
    m__io__raw_rest = nullptr;
    m_tuple5 = nullptr;
    f_len_mod = false;
    f_char5 = false;
    f_tuple5 = false;
    _read();
}

void expr_2_t::mod_str_t::_read() {
    m_len_orig = m__io->read_u2le();
    m_str = kaitai::kstream::bytes_to_str(m__io->read_bytes(len_mod()), std::string("UTF-8"));
    m__raw_rest = m__io->read_bytes(3);
    m__io__raw_rest = std::unique_ptr<kaitai::kstream>(new kaitai::kstream(m__raw_rest));
    m_rest = std::unique_ptr<tuple_t>(new tuple_t(m__io__raw_rest.get(), this, m__root));
}

expr_2_t::mod_str_t::~mod_str_t() {
    _clean_up();
}

void expr_2_t::mod_str_t::_clean_up() {
    if (f_char5) {
    }
    if (f_tuple5) {
    }
}

int32_t expr_2_t::mod_str_t::len_mod() {
    if (f_len_mod)
        return m_len_mod;
    m_len_mod = (len_orig() - 3);
    f_len_mod = true;
    return m_len_mod;
}

std::string expr_2_t::mod_str_t::char5() {
    if (f_char5)
        return m_char5;
    std::streampos _pos = m__io->pos();
    m__io->seek(5);
    m_char5 = kaitai::kstream::bytes_to_str(m__io->read_bytes(1), std::string("ASCII"));
    m__io->seek(_pos);
    f_char5 = true;
    return m_char5;
}

expr_2_t::tuple_t* expr_2_t::mod_str_t::tuple5() {
    if (f_tuple5)
        return m_tuple5.get();
    std::streampos _pos = m__io->pos();
    m__io->seek(5);
    m_tuple5 = std::unique_ptr<tuple_t>(new tuple_t(m__io, this, m__root));
    m__io->seek(_pos);
    f_tuple5 = true;
    return m_tuple5.get();
}

expr_2_t::tuple_t::tuple_t(kaitai::kstream* p__io, expr_2_t::mod_str_t* p__parent, expr_2_t* p__root) : kaitai::kstruct(p__io) {
    m__parent = p__parent;
    m__root = p__root;
    f_avg = false;
    _read();
}

void expr_2_t::tuple_t::_read() {
    m_byte0 = m__io->read_u1();
    m_byte1 = m__io->read_u1();
    m_byte2 = m__io->read_u1();
}

expr_2_t::tuple_t::~tuple_t() {
    _clean_up();
}

void expr_2_t::tuple_t::_clean_up() {
}

int32_t expr_2_t::tuple_t::avg() {
    if (f_avg)
        return m_avg;
    m_avg = ((byte1() + byte2()) / 2);
    f_avg = true;
    return m_avg;
}

int32_t expr_2_t::str1_len_mod() {
    if (f_str1_len_mod)
        return m_str1_len_mod;
    m_str1_len_mod = str1()->len_mod();
    f_str1_len_mod = true;
    return m_str1_len_mod;
}

int32_t expr_2_t::str1_len() {
    if (f_str1_len)
        return m_str1_len;
    m_str1_len = str1()->str().length();
    f_str1_len = true;
    return m_str1_len;
}

expr_2_t::tuple_t* expr_2_t::str1_tuple5() {
    if (f_str1_tuple5)
        return m_str1_tuple5;
    m_str1_tuple5 = str1()->tuple5();
    f_str1_tuple5 = true;
    return m_str1_tuple5;
}

expr_2_t::tuple_t* expr_2_t::str2_tuple5() {
    if (f_str2_tuple5)
        return m_str2_tuple5;
    m_str2_tuple5 = str2()->tuple5();
    f_str2_tuple5 = true;
    return m_str2_tuple5;
}

int32_t expr_2_t::str1_avg() {
    if (f_str1_avg)
        return m_str1_avg;
    m_str1_avg = str1()->rest()->avg();
    f_str1_avg = true;
    return m_str1_avg;
}

uint8_t expr_2_t::str1_byte1() {
    if (f_str1_byte1)
        return m_str1_byte1;
    m_str1_byte1 = str1()->rest()->byte1();
    f_str1_byte1 = true;
    return m_str1_byte1;
}

std::string expr_2_t::str1_char5() {
    if (f_str1_char5)
        return m_str1_char5;
    m_str1_char5 = str1()->char5();
    f_str1_char5 = true;
    return m_str1_char5;
}
