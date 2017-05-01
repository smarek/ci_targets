# This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

require 'kaitai/struct/struct'

unless Gem::Version.new(Kaitai::Struct::VERSION) >= Gem::Version.new('0.7')
  raise "Incompatible Kaitai Struct Ruby API: 0.7 or later is required, but you have #{Kaitai::Struct::VERSION}"
end

class NestedTypes < Kaitai::Struct::Struct
  def initialize(_io, _parent = nil, _root = self)
    super(_io, _parent, _root)
    @one = SubtypeA.new(@_io, self, @_root)
    @two = SubtypeB.new(@_io, self, @_root)
  end
  class SubtypeA < Kaitai::Struct::Struct
    def initialize(_io, _parent = nil, _root = self)
      super(_io, _parent, _root)
      @typed_at_root = SubtypeB.new(@_io, self, @_root)
      @typed_here = SubtypeC.new(@_io, self, @_root)
    end
    class SubtypeC < Kaitai::Struct::Struct
      def initialize(_io, _parent = nil, _root = self)
        super(_io, _parent, _root)
        @value_c = @_io.read_s1
      end
      attr_reader :value_c
    end
    attr_reader :typed_at_root
    attr_reader :typed_here
  end
  class SubtypeB < Kaitai::Struct::Struct
    def initialize(_io, _parent = nil, _root = self)
      super(_io, _parent, _root)
      @value_b = @_io.read_s1
    end
    attr_reader :value_b
  end
  attr_reader :one
  attr_reader :two
end
