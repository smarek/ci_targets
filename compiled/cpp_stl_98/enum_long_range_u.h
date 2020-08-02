#ifndef ENUM_LONG_RANGE_U_H_
#define ENUM_LONG_RANGE_U_H_

// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

#include "kaitai/kaitaistruct.h"
#include <stdint.h>

#if KAITAI_STRUCT_VERSION < 9000L
#error "Incompatible Kaitai Struct C++/STL API: version 0.9 or later is required"
#endif

class enum_long_range_u_t : public kaitai::kstruct {

public:

    enum constants_t {
        CONSTANTS_ZERO = 0,
        CONSTANTS_INT_MAX = 4294967295,
        CONSTANTS_INT_OVER_MAX = 4294967296,
        CONSTANTS_LONG_MAX = 9223372036854775807
    };

    enum_long_range_u_t(kaitai::kstream* p__io, kaitai::kstruct* p__parent = 0, enum_long_range_u_t* p__root = 0);

private:
    void _read();

public:
    ~enum_long_range_u_t();
    void _clean_up();

private:
    constants_t m_f1;
    constants_t m_f2;
    constants_t m_f3;
    constants_t m_f4;
    enum_long_range_u_t* m__root;
    kaitai::kstruct* m__parent;

public:
    constants_t f1() const { return m_f1; }
    constants_t f2() const { return m_f2; }
    constants_t f3() const { return m_f3; }
    constants_t f4() const { return m_f4; }
    enum_long_range_u_t* _root() const { return m__root; }
    kaitai::kstruct* _parent() const { return m__parent; }
};

#endif  // ENUM_LONG_RANGE_U_H_
