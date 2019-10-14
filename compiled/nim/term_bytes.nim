import ../../runtime/nim/kaitai



type
  TermBytes* = ref TermBytesObj
  TermBytesObj* = object
    io: KaitaiStream
    root*: TermBytes
    parent*: ref RootObj
    s1*: seq[byte]
    s2*: seq[byte]
    s3*: seq[byte]

# TermBytes
proc read*(_: typedesc[TermBytes], io: KaitaiStream, root: TermBytes, parent: ref RootObj): owned TermBytes =
  result = new(TermBytes)
  let root = if root == nil: cast[TermBytes](result) else: root
  result.io = io
  result.root = root
  result.parent = parent

  let s1 = readBytesTerm(io, 124, false, true, true)
  result.s1 = s1
  let s2 = readBytesTerm(io, 124, false, false, true)
  result.s2 = s2
  let s3 = readBytesTerm(io, 64, true, true, true)
  result.s3 = s3


proc fromFile*(_: typedesc[TermBytes], filename: string): owned TermBytes =
  TermBytes.read(newKaitaiStream(filename), nil, nil)

proc `=destroy`(x: var TermBytesObj) =
  close(x.io)
