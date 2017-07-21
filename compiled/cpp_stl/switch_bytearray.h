#ifndef SWITCH_BYTEARRAY_H_
#define SWITCH_BYTEARRAY_H_

// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

#include <kaitai/kaitaistruct.h>
#include <kaitai/kaitaistream.h>

#include <stdint.h>
#include <vector>

#if KAITAI_STRUCT_VERSION < 7000L
#error "Incompatible Kaitai Struct C++/STL API: version 0.7 or later is required"
#endif

class switch_bytearray_t : public kaitai::kstruct {

public:
    class opcode_t;

    switch_bytearray_t(kaitai::kstream* p_io, kaitai::kstruct* p_parent = 0, switch_bytearray_t* p_root = 0);
    void _read();
    ~switch_bytearray_t();

    class opcode_t : public kaitai::kstruct {

    public:
        class intval_t;
        class strval_t;

        opcode_t(kaitai::kstream* p_io, switch_bytearray_t* p_parent = 0, switch_bytearray_t* p_root = 0);
        void _read();
        ~opcode_t();

        class intval_t : public kaitai::kstruct {

        public:

            intval_t(kaitai::kstream* p_io, switch_bytearray_t::opcode_t* p_parent = 0, switch_bytearray_t* p_root = 0);
            void _read();
            ~intval_t();

        private:
            uint8_t m_value;
            switch_bytearray_t* m__root;
            switch_bytearray_t::opcode_t* m__parent;

        public:
            uint8_t value() const { return m_value; }
            switch_bytearray_t* _root() const { return m__root; }
            switch_bytearray_t::opcode_t* _parent() const { return m__parent; }
        };

        class strval_t : public kaitai::kstruct {

        public:

            strval_t(kaitai::kstream* p_io, switch_bytearray_t::opcode_t* p_parent = 0, switch_bytearray_t* p_root = 0);
            void _read();
            ~strval_t();

        private:
            std::string m_value;
            switch_bytearray_t* m__root;
            switch_bytearray_t::opcode_t* m__parent;

        public:
            std::string value() const { return m_value; }
            switch_bytearray_t* _root() const { return m__root; }
            switch_bytearray_t::opcode_t* _parent() const { return m__parent; }
        };

    private:
        std::string m_code;
        kaitai::kstruct* m_body;
        bool n_body;

    public:
        bool _is_null_body() { body(); return n_body; };

    private:
        switch_bytearray_t* m__root;
        switch_bytearray_t* m__parent;

    public:
        std::string code() const { return m_code; }
        kaitai::kstruct* body() const { return m_body; }
        switch_bytearray_t* _root() const { return m__root; }
        switch_bytearray_t* _parent() const { return m__parent; }
    };

private:
    std::vector<opcode_t*>* m_opcodes;
    switch_bytearray_t* m__root;
    kaitai::kstruct* m__parent;

public:
    std::vector<opcode_t*>* opcodes() const { return m_opcodes; }
    switch_bytearray_t* _root() const { return m__root; }
    kaitai::kstruct* _parent() const { return m__parent; }
};

#endif  // SWITCH_BYTEARRAY_H_