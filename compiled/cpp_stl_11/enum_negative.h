#pragma once

// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

#include "kaitai/kaitaistruct.h"
#include <stdint.h>
#include <memory>

#if KAITAI_STRUCT_VERSION < 9000L
#error "Incompatible Kaitai Struct C++/STL API: version 0.9 or later is required"
#endif

class enum_negative_t : public kaitai::kstruct {

public:

    enum constants_t {
        CONSTANTS_NEGATIVE_ONE = -1,
        CONSTANTS_POSITIVE_ONE = 1
    };

    enum_negative_t(kaitai::kstream* p__io, kaitai::kstruct* p__parent = nullptr, enum_negative_t* p__root = nullptr);

private:
    void _read();

public:
    ~enum_negative_t();
    void _clean_up();

private:
    constants_t m_f1;
    constants_t m_f2;
    enum_negative_t* m__root;
    kaitai::kstruct* m__parent;

public:
    constants_t f1() const { return m_f1; }
    constants_t f2() const { return m_f2; }
    enum_negative_t* _root() const { return m__root; }
    kaitai::kstruct* _parent() const { return m__parent; }
};
