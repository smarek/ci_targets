#pragma once

// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

#include "kaitai/kaitaistruct.h"
#include <stdint.h>
#include <memory>
#include <vector>

#if KAITAI_STRUCT_VERSION < 9000L
#error "Incompatible Kaitai Struct C++/STL API: version 0.9 or later is required"
#endif

class switch_manual_int_size_else_t : public kaitai::kstruct {

public:
    class chunk_t;

    switch_manual_int_size_else_t(kaitai::kstream* p__io, kaitai::kstruct* p__parent = nullptr, switch_manual_int_size_else_t* p__root = nullptr);

private:
    void _read();

public:
    ~switch_manual_int_size_else_t();

    class chunk_t : public kaitai::kstruct {

    public:
        class chunk_meta_t;
        class chunk_dir_t;
        class dummy_t;

        chunk_t(kaitai::kstream* p__io, switch_manual_int_size_else_t* p__parent = nullptr, switch_manual_int_size_else_t* p__root = nullptr);

    private:
        void _read();

    public:
        ~chunk_t();

        class chunk_meta_t : public kaitai::kstruct {

        public:

            chunk_meta_t(kaitai::kstream* p__io, switch_manual_int_size_else_t::chunk_t* p__parent = nullptr, switch_manual_int_size_else_t* p__root = nullptr);

        private:
            void _read();

        public:
            ~chunk_meta_t();

        private:
            std::string m_title;
            std::string m_author;
            switch_manual_int_size_else_t* m__root;
            switch_manual_int_size_else_t::chunk_t* m__parent;

        public:
            std::string title() const { return m_title; }
            std::string author() const { return m_author; }
            switch_manual_int_size_else_t* _root() const { return m__root; }
            switch_manual_int_size_else_t::chunk_t* _parent() const { return m__parent; }
        };

        class chunk_dir_t : public kaitai::kstruct {

        public:

            chunk_dir_t(kaitai::kstream* p__io, switch_manual_int_size_else_t::chunk_t* p__parent = nullptr, switch_manual_int_size_else_t* p__root = nullptr);

        private:
            void _read();

        public:
            ~chunk_dir_t();

        private:
            std::unique_ptr<std::vector<std::string>> m_entries;
            switch_manual_int_size_else_t* m__root;
            switch_manual_int_size_else_t::chunk_t* m__parent;

        public:
            std::unique_ptr<std::vector<std::string>> entries() const { return m_entries.get(); }
            switch_manual_int_size_else_t* _root() const { return m__root; }
            switch_manual_int_size_else_t::chunk_t* _parent() const { return m__parent; }
        };

        class dummy_t : public kaitai::kstruct {

        public:

            dummy_t(kaitai::kstream* p__io, switch_manual_int_size_else_t::chunk_t* p__parent = nullptr, switch_manual_int_size_else_t* p__root = nullptr);

        private:
            void _read();

        public:
            ~dummy_t();

        private:
            std::string m_rest;
            switch_manual_int_size_else_t* m__root;
            switch_manual_int_size_else_t::chunk_t* m__parent;

        public:
            std::string rest() const { return m_rest; }
            switch_manual_int_size_else_t* _root() const { return m__root; }
            switch_manual_int_size_else_t::chunk_t* _parent() const { return m__parent; }
        };

    private:
        uint8_t m_code;
        uint32_t m_size;
        kaitai::kstruct* m_body;
        switch_manual_int_size_else_t* m__root;
        switch_manual_int_size_else_t* m__parent;
        std::string m__raw_body;
        kaitai::kstream* m__io__raw_body;

    public:
        uint8_t code() const { return m_code; }
        uint32_t size() const { return m_size; }
        kaitai::kstruct* body() const { return m_body; }
        switch_manual_int_size_else_t* _root() const { return m__root; }
        switch_manual_int_size_else_t* _parent() const { return m__parent; }
        std::string _raw_body() const { return m__raw_body; }
        kaitai::kstream* _io__raw_body() const { return m__io__raw_body; }
    };

private:
    std::unique_ptr<std::vector<std::unique_ptr<chunk_t>>> m_chunks;
    switch_manual_int_size_else_t* m__root;
    kaitai::kstruct* m__parent;

public:
    std::unique_ptr<std::vector<std::unique_ptr<chunk_t>>> chunks() const { return m_chunks.get(); }
    switch_manual_int_size_else_t* _root() const { return m__root; }
    kaitai::kstruct* _parent() const { return m__parent; }
};
