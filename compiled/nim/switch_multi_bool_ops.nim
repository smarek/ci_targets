import kaitai_struct_nim_runtime
import options

template defineEnum(typ) =
  type typ* = distinct int64
  proc `==`*(x, y: typ): bool {.borrow.}

type
  SwitchMultiBoolOps* = ref object of KaitaiStruct
    opcodes*: seq[SwitchMultiBoolOps_Opcode]
    parent*: KaitaiStruct
  SwitchMultiBoolOps_Opcode* = ref object of KaitaiStruct
    code*: uint8
    body*: uint64
    parent*: SwitchMultiBoolOps

proc read*(_: typedesc[SwitchMultiBoolOps], io: KaitaiStream, root: KaitaiStruct, parent: KaitaiStruct): SwitchMultiBoolOps
proc read*(_: typedesc[SwitchMultiBoolOps_Opcode], io: KaitaiStream, root: KaitaiStruct, parent: SwitchMultiBoolOps): SwitchMultiBoolOps_Opcode


proc read*(_: typedesc[SwitchMultiBoolOps], io: KaitaiStream, root: KaitaiStruct, parent: KaitaiStruct): SwitchMultiBoolOps =
  template this: untyped = result
  this = new(SwitchMultiBoolOps)
  let root = if root == nil: cast[SwitchMultiBoolOps](this) else: cast[SwitchMultiBoolOps](root)
  this.io = io
  this.root = root
  this.parent = parent

  block:
    var i: int
    while not this.io.isEof:
      let it = SwitchMultiBoolOps_Opcode.read(this.io, this.root, this)
      this.opcodes.add(it)
      inc i

proc fromFile*(_: typedesc[SwitchMultiBoolOps], filename: string): SwitchMultiBoolOps =
  SwitchMultiBoolOps.read(newKaitaiFileStream(filename), nil, nil)

proc read*(_: typedesc[SwitchMultiBoolOps_Opcode], io: KaitaiStream, root: KaitaiStruct, parent: SwitchMultiBoolOps): SwitchMultiBoolOps_Opcode =
  template this: untyped = result
  this = new(SwitchMultiBoolOps_Opcode)
  let root = if root == nil: cast[SwitchMultiBoolOps](this) else: cast[SwitchMultiBoolOps](root)
  this.io = io
  this.root = root
  this.parent = parent

  let codeExpr = this.io.readU1()
  this.code = codeExpr
  case ord((if  ((this.code > 0) and (this.code <= 8) and ((if this.code != 10: true else: false))) : this.code else: 0))
  of 1:
    let bodyExpr = uint64(this.io.readU1())
    this.body = bodyExpr
  of 2:
    let bodyExpr = uint64(this.io.readU2le())
    this.body = bodyExpr
  of 4:
    let bodyExpr = uint64(this.io.readU4le())
    this.body = bodyExpr
  of 8:
    let bodyExpr = this.io.readU8le()
    this.body = bodyExpr
  else: discard

proc fromFile*(_: typedesc[SwitchMultiBoolOps_Opcode], filename: string): SwitchMultiBoolOps_Opcode =
  SwitchMultiBoolOps_Opcode.read(newKaitaiFileStream(filename), nil, nil)

