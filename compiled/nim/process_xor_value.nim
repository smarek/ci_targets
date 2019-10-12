import ../../runtime/nim/kaitai



type
  ProcessXorValue* = ref ProcessXorValueObj
  ProcessXorValueObj* = object
    io: KaitaiStream
    root*: ProcessXorValue
    parent*: ref RootObj
    key*: uint8
    buf*: seq[byte]

# ProcessXorValue
proc read*(_: typedesc[ProcessXorValue], io: KaitaiStream, root: ProcessXorValue, parent: ref RootObj): owned ProcessXorValue =
  result = new(ProcessXorValue)
  let root = if root == nil: cast[ProcessXorValue](result) else: root
  result.io = io
  result.root = root
  result.parent = parent

  result.key = readU1(io)
  result.buf = readBytesFull(io)


proc fromFile*(_: typedesc[ProcessXorValue], filename: string): owned ProcessXorValue =
  ProcessXorValue.read(newKaitaiStream(filename), nil, nil)

proc `=destroy`(x: var ProcessXorValueObj) =
  close(x.io)

