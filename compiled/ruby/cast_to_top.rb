# This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

require 'kaitai/struct/struct'

unless Gem::Version.new(Kaitai::Struct::VERSION) >= Gem::Version.new('0.7')
  raise "Incompatible Kaitai Struct Ruby API: 0.7 or later is required, but you have #{Kaitai::Struct::VERSION}"
end

class CastToTop < Kaitai::Struct::Struct
  def initialize(_io, _parent = nil, _root = self)
    super(_io, _parent, _root)
    _read
  end

  def _read
    @code = @_io.read_u1
    self
  end
  def header
    return @header unless @header.nil?
    _pos = @_io.pos
    @_io.seek(1)
    @header = CastToTop.new(@_io)
    @_io.seek(_pos)
    @header
  end
  def header_casted
    return @header_casted unless @header_casted.nil?
    @header_casted = header
    @header_casted
  end
  attr_reader :code
end