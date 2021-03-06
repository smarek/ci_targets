import kaitai_struct_nim_runtime
import options

type
  SwitchManualInt* = ref object of KaitaiStruct
    `opcodes`*: seq[SwitchManualInt_Opcode]
    `parent`*: KaitaiStruct
  SwitchManualInt_Opcode* = ref object of KaitaiStruct
    `code`*: uint8
    `body`*: KaitaiStruct
    `parent`*: SwitchManualInt
  SwitchManualInt_Opcode_Intval* = ref object of KaitaiStruct
    `value`*: uint8
    `parent`*: SwitchManualInt_Opcode
  SwitchManualInt_Opcode_Strval* = ref object of KaitaiStruct
    `value`*: string
    `parent`*: SwitchManualInt_Opcode

proc read*(_: typedesc[SwitchManualInt], io: KaitaiStream, root: KaitaiStruct, parent: KaitaiStruct): SwitchManualInt
proc read*(_: typedesc[SwitchManualInt_Opcode], io: KaitaiStream, root: KaitaiStruct, parent: SwitchManualInt): SwitchManualInt_Opcode
proc read*(_: typedesc[SwitchManualInt_Opcode_Intval], io: KaitaiStream, root: KaitaiStruct, parent: SwitchManualInt_Opcode): SwitchManualInt_Opcode_Intval
proc read*(_: typedesc[SwitchManualInt_Opcode_Strval], io: KaitaiStream, root: KaitaiStruct, parent: SwitchManualInt_Opcode): SwitchManualInt_Opcode_Strval


proc read*(_: typedesc[SwitchManualInt], io: KaitaiStream, root: KaitaiStruct, parent: KaitaiStruct): SwitchManualInt =
  template this: untyped = result
  this = new(SwitchManualInt)
  let root = if root == nil: cast[SwitchManualInt](this) else: cast[SwitchManualInt](root)
  this.io = io
  this.root = root
  this.parent = parent

  block:
    var i: int
    while not this.io.isEof:
      let it = SwitchManualInt_Opcode.read(this.io, this.root, this)
      this.opcodes.add(it)
      inc i

proc fromFile*(_: typedesc[SwitchManualInt], filename: string): SwitchManualInt =
  SwitchManualInt.read(newKaitaiFileStream(filename), nil, nil)

proc read*(_: typedesc[SwitchManualInt_Opcode], io: KaitaiStream, root: KaitaiStruct, parent: SwitchManualInt): SwitchManualInt_Opcode =
  template this: untyped = result
  this = new(SwitchManualInt_Opcode)
  let root = if root == nil: cast[SwitchManualInt](this) else: cast[SwitchManualInt](root)
  this.io = io
  this.root = root
  this.parent = parent

  let codeExpr = this.io.readU1()
  this.code = codeExpr
  block:
    let on = this.code
    if on == 73:
      let bodyExpr = SwitchManualInt_Opcode_Intval.read(this.io, this.root, this)
      this.body = bodyExpr
    elif on == 83:
      let bodyExpr = SwitchManualInt_Opcode_Strval.read(this.io, this.root, this)
      this.body = bodyExpr

proc fromFile*(_: typedesc[SwitchManualInt_Opcode], filename: string): SwitchManualInt_Opcode =
  SwitchManualInt_Opcode.read(newKaitaiFileStream(filename), nil, nil)

proc read*(_: typedesc[SwitchManualInt_Opcode_Intval], io: KaitaiStream, root: KaitaiStruct, parent: SwitchManualInt_Opcode): SwitchManualInt_Opcode_Intval =
  template this: untyped = result
  this = new(SwitchManualInt_Opcode_Intval)
  let root = if root == nil: cast[SwitchManualInt](this) else: cast[SwitchManualInt](root)
  this.io = io
  this.root = root
  this.parent = parent

  let valueExpr = this.io.readU1()
  this.value = valueExpr

proc fromFile*(_: typedesc[SwitchManualInt_Opcode_Intval], filename: string): SwitchManualInt_Opcode_Intval =
  SwitchManualInt_Opcode_Intval.read(newKaitaiFileStream(filename), nil, nil)

proc read*(_: typedesc[SwitchManualInt_Opcode_Strval], io: KaitaiStream, root: KaitaiStruct, parent: SwitchManualInt_Opcode): SwitchManualInt_Opcode_Strval =
  template this: untyped = result
  this = new(SwitchManualInt_Opcode_Strval)
  let root = if root == nil: cast[SwitchManualInt](this) else: cast[SwitchManualInt](root)
  this.io = io
  this.root = root
  this.parent = parent

  let valueExpr = encode(this.io.readBytesTerm(0, false, true, true), "ASCII")
  this.value = valueExpr

proc fromFile*(_: typedesc[SwitchManualInt_Opcode_Strval], filename: string): SwitchManualInt_Opcode_Strval =
  SwitchManualInt_Opcode_Strval.read(newKaitaiFileStream(filename), nil, nil)

