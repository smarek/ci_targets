import kaitai_struct_nim_runtime
import options

template defineEnum(typ) =
  type typ* = distinct int64
  proc `==`*(x, y: typ): bool {.borrow.}

type
  CastNested* = ref object of KaitaiStruct
    opcodes*: seq[CastNested_Opcode]
    parent*: KaitaiStruct
    opcodes0StrInst*: Option[CastNested_Opcode_Strval]
    opcodes0StrValueInst*: string
    opcodes1IntInst*: Option[CastNested_Opcode_Intval]
    opcodes1IntValueInst*: Option[uint8]
  CastNested_Opcode* = ref object of KaitaiStruct
    code*: uint8
    body*: KaitaiStruct
    parent*: CastNested
  CastNested_Opcode_Intval* = ref object of KaitaiStruct
    value*: uint8
    parent*: CastNested_Opcode
  CastNested_Opcode_Strval* = ref object of KaitaiStruct
    value*: string
    parent*: CastNested_Opcode

proc read*(_: typedesc[CastNested], io: KaitaiStream, root: KaitaiStruct, parent: KaitaiStruct): CastNested
proc read*(_: typedesc[CastNested_Opcode], io: KaitaiStream, root: KaitaiStruct, parent: CastNested): CastNested_Opcode
proc read*(_: typedesc[CastNested_Opcode_Intval], io: KaitaiStream, root: KaitaiStruct, parent: CastNested_Opcode): CastNested_Opcode_Intval
proc read*(_: typedesc[CastNested_Opcode_Strval], io: KaitaiStream, root: KaitaiStruct, parent: CastNested_Opcode): CastNested_Opcode_Strval

proc opcodes0Str*(this: CastNested): CastNested_Opcode_Strval
proc opcodes0StrValue*(this: CastNested): string
proc opcodes1Int*(this: CastNested): CastNested_Opcode_Intval
proc opcodes1IntValue*(this: CastNested): uint8

proc read*(_: typedesc[CastNested], io: KaitaiStream, root: KaitaiStruct, parent: KaitaiStruct): CastNested =
  template this: untyped = result
  this = new(CastNested)
  let root = if root == nil: cast[CastNested](this) else: cast[CastNested](root)
  this.io = io
  this.root = root
  this.parent = parent

  block:
    var i: int
    while not this.io.isEof:
      let it = CastNested_Opcode.read(this.io, this.root, this)
      this.opcodes.add(it)
      inc i

proc opcodes0Str(this: CastNested): CastNested_Opcode_Strval = 
  if isSome(this.opcodes0StrInst):
    return get(this.opcodes0StrInst)
  let opcodes0StrInstExpr = CastNested_Opcode_Strval((CastNested_Opcode_Strval(this.opcodes[0].body)))
  this.opcodes0StrInst = opcodes0StrInstExpr
  if isSome(this.opcodes0StrInst):
    return get(this.opcodes0StrInst)

proc opcodes0StrValue(this: CastNested): string = 
  if this.opcodes0StrValueInst.len != 0:
    return this.opcodes0StrValueInst
  let opcodes0StrValueInstExpr = string((CastNested_Opcode_Strval(this.opcodes[0].body)).value)
  this.opcodes0StrValueInst = opcodes0StrValueInstExpr
  if this.opcodes0StrValueInst.len != 0:
    return this.opcodes0StrValueInst

proc opcodes1Int(this: CastNested): CastNested_Opcode_Intval = 
  if isSome(this.opcodes1IntInst):
    return get(this.opcodes1IntInst)
  let opcodes1IntInstExpr = CastNested_Opcode_Intval((CastNested_Opcode_Intval(this.opcodes[1].body)))
  this.opcodes1IntInst = opcodes1IntInstExpr
  if isSome(this.opcodes1IntInst):
    return get(this.opcodes1IntInst)

proc opcodes1IntValue(this: CastNested): uint8 = 
  if isSome(this.opcodes1IntValueInst):
    return get(this.opcodes1IntValueInst)
  let opcodes1IntValueInstExpr = uint8((CastNested_Opcode_Intval(this.opcodes[1].body)).value)
  this.opcodes1IntValueInst = opcodes1IntValueInstExpr
  if isSome(this.opcodes1IntValueInst):
    return get(this.opcodes1IntValueInst)

proc fromFile*(_: typedesc[CastNested], filename: string): CastNested =
  CastNested.read(newKaitaiFileStream(filename), nil, nil)

proc read*(_: typedesc[CastNested_Opcode], io: KaitaiStream, root: KaitaiStruct, parent: CastNested): CastNested_Opcode =
  template this: untyped = result
  this = new(CastNested_Opcode)
  let root = if root == nil: cast[CastNested](this) else: cast[CastNested](root)
  this.io = io
  this.root = root
  this.parent = parent

  let codeExpr = this.io.readU1()
  this.code = codeExpr
  case ord(this.code)
  of 73:
    let bodyExpr = CastNested_Opcode_Intval.read(this.io, this.root, this)
    this.body = bodyExpr
  of 83:
    let bodyExpr = CastNested_Opcode_Strval.read(this.io, this.root, this)
    this.body = bodyExpr
  else: discard

proc fromFile*(_: typedesc[CastNested_Opcode], filename: string): CastNested_Opcode =
  CastNested_Opcode.read(newKaitaiFileStream(filename), nil, nil)

proc read*(_: typedesc[CastNested_Opcode_Intval], io: KaitaiStream, root: KaitaiStruct, parent: CastNested_Opcode): CastNested_Opcode_Intval =
  template this: untyped = result
  this = new(CastNested_Opcode_Intval)
  let root = if root == nil: cast[CastNested](this) else: cast[CastNested](root)
  this.io = io
  this.root = root
  this.parent = parent

  let valueExpr = this.io.readU1()
  this.value = valueExpr

proc fromFile*(_: typedesc[CastNested_Opcode_Intval], filename: string): CastNested_Opcode_Intval =
  CastNested_Opcode_Intval.read(newKaitaiFileStream(filename), nil, nil)

proc read*(_: typedesc[CastNested_Opcode_Strval], io: KaitaiStream, root: KaitaiStruct, parent: CastNested_Opcode): CastNested_Opcode_Strval =
  template this: untyped = result
  this = new(CastNested_Opcode_Strval)
  let root = if root == nil: cast[CastNested](this) else: cast[CastNested](root)
  this.io = io
  this.root = root
  this.parent = parent

  let valueExpr = encode(this.io.readBytesTerm(0, false, true, true), "ASCII")
  this.value = valueExpr

proc fromFile*(_: typedesc[CastNested_Opcode_Strval], filename: string): CastNested_Opcode_Strval =
  CastNested_Opcode_Strval.read(newKaitaiFileStream(filename), nil, nil)

