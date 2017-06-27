# This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

require 'kaitai/struct/struct'

unless Gem::Version.new(Kaitai::Struct::VERSION) >= Gem::Version.new('0.7')
  raise "Incompatible Kaitai Struct Ruby API: 0.7 or later is required, but you have #{Kaitai::Struct::VERSION}"
end

class EnumNegative < Kaitai::Struct::Struct

  CONSTANTS = {
    -1 => :constants_negative_one,
    1 => :constants_positive_one,
  }
  I__CONSTANTS = CONSTANTS.invert
  def initialize(_io, _parent = nil, _root = self)
    super(_io, _parent, _root)
    _read
  end
  def _read
    @f1 = Kaitai::Struct::Stream::resolve_enum(CONSTANTS, @_io.read_s1)
    @f2 = Kaitai::Struct::Stream::resolve_enum(CONSTANTS, @_io.read_s1)
  end
  attr_reader :f1
  attr_reader :f2
end
