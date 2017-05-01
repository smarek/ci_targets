# This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

require 'kaitai/struct/struct'

unless Gem::Version.new(Kaitai::Struct::VERSION) >= Gem::Version.new('0.7')
  raise "Incompatible Kaitai Struct Ruby API: 0.7 or later is required, but you have #{Kaitai::Struct::VERSION}"
end

class Debug0 < Kaitai::Struct::Struct
  attr_reader :_debug
  SEQ_FIELDS = ["one", "array_of_ints", "_unnamed2"]
  def initialize(_io, _parent = nil, _root = self)
    super(_io, _parent, _root)
    @_debug = {}
  end

  def _read
    (@_debug['one'] ||= {})[:start] = @_io.pos
    @one = @_io.read_u1
    (@_debug['one'] ||= {})[:end] = @_io.pos
    (@_debug['array_of_ints'] ||= {})[:start] = @_io.pos
    @array_of_ints = Array.new(3)
    (3).times { |i|
      (@_debug['array_of_ints'][:arr] ||= [])[i] = {:start => @_io.pos}
      @array_of_ints[i] = @_io.read_u1
      @_debug['array_of_ints'][:arr][i][:end] = @_io.pos
    }
    (@_debug['array_of_ints'] ||= {})[:end] = @_io.pos
    (@_debug['_unnamed2'] ||= {})[:start] = @_io.pos
    @_unnamed2 = @_io.read_u1
    (@_debug['_unnamed2'] ||= {})[:end] = @_io.pos

    self
  end
  attr_reader :one
  attr_reader :array_of_ints
  attr_reader :_unnamed2
end
