#pragma once

// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

#include "kaitai/kaitaistruct.h"
#include <stdint.h>
#include <memory>
#include <vector>

#if KAITAI_STRUCT_VERSION < 9000L
#error "Incompatible Kaitai Struct C++/STL API: version 0.9 or later is required"
#endif

class repeat_eos_struct_t : public kaitai::kstruct {

public:
    class chunk_t;

    repeat_eos_struct_t(kaitai::kstream* p__io, kaitai::kstruct* p__parent = nullptr, repeat_eos_struct_t* p__root = nullptr);

private:
    void _read();

public:
    ~repeat_eos_struct_t();
    void _clean_up();

    class chunk_t : public kaitai::kstruct {

    public:

        chunk_t(kaitai::kstream* p__io, repeat_eos_struct_t* p__parent = nullptr, repeat_eos_struct_t* p__root = nullptr);

    private:
        void _read();

    public:
        ~chunk_t();
        void _clean_up();

    private:
        uint32_t m_offset;
        uint32_t m_len;
        repeat_eos_struct_t* m__root;
        repeat_eos_struct_t* m__parent;

    public:
        uint32_t offset() const { return m_offset; }
        uint32_t len() const { return m_len; }
        repeat_eos_struct_t* _root() const { return m__root; }
        repeat_eos_struct_t* _parent() const { return m__parent; }
    };

private:
    std::unique_ptr<std::vector<std::unique_ptr<chunk_t>>> m_chunks;
    repeat_eos_struct_t* m__root;
    kaitai::kstruct* m__parent;

public:
    std::vector<std::unique_ptr<chunk_t>>* chunks() const { return m_chunks.get(); }
    repeat_eos_struct_t* _root() const { return m__root; }
    kaitai::kstruct* _parent() const { return m__parent; }
};
