import kaitai_struct_nim_runtime

type
  InstanceStdArray* = ref InstanceStdArrayObj
  InstanceStdArrayObj* = object
    ofs*: uint32
    entrySize*: uint32
    qtyEntries*: uint32
    io*: KaitaiStream
    root*: InstanceStdArray
    parent*: ref RootObj

### InstanceStdArray ###
proc read*(_: typedesc[InstanceStdArray], io: KaitaiStream, root: InstanceStdArray, parent: ref RootObj): InstanceStdArray =
  result = new(InstanceStdArray)
  let root = if root == nil: cast[InstanceStdArray](result) else: root
  result.io = io
  result.root = root
  result.parent = parent
  let ofs = io.readU4le()
  result.ofs = ofs
  let entrySize = io.readU4le()
  result.entrySize = entrySize
  let qtyEntries = io.readU4le()
  result.qtyEntries = qtyEntries

proc fromFile*(_: typedesc[InstanceStdArray], filename: string): InstanceStdArray =
  InstanceStdArray.read(newKaitaiFileStream(filename), nil, nil)

proc `=destroy`(x: var InstanceStdArrayObj) =
  close(x.io)

