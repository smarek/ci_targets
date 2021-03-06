import kaitai_struct_nim_runtime
import options

type
  SwitchManualStr* = ref object of KaitaiStruct
    `opcodes`*: seq[SwitchManualStr_Opcode]
    `parent`*: KaitaiStruct
  SwitchManualStr_Opcode* = ref object of KaitaiStruct
    `code`*: string
    `body`*: KaitaiStruct
    `parent`*: SwitchManualStr
  SwitchManualStr_Opcode_Intval* = ref object of KaitaiStruct
    `value`*: uint8
    `parent`*: SwitchManualStr_Opcode
  SwitchManualStr_Opcode_Strval* = ref object of KaitaiStruct
    `value`*: string
    `parent`*: SwitchManualStr_Opcode

proc read*(_: typedesc[SwitchManualStr], io: KaitaiStream, root: KaitaiStruct, parent: KaitaiStruct): SwitchManualStr
proc read*(_: typedesc[SwitchManualStr_Opcode], io: KaitaiStream, root: KaitaiStruct, parent: SwitchManualStr): SwitchManualStr_Opcode
proc read*(_: typedesc[SwitchManualStr_Opcode_Intval], io: KaitaiStream, root: KaitaiStruct, parent: SwitchManualStr_Opcode): SwitchManualStr_Opcode_Intval
proc read*(_: typedesc[SwitchManualStr_Opcode_Strval], io: KaitaiStream, root: KaitaiStruct, parent: SwitchManualStr_Opcode): SwitchManualStr_Opcode_Strval


proc read*(_: typedesc[SwitchManualStr], io: KaitaiStream, root: KaitaiStruct, parent: KaitaiStruct): SwitchManualStr =
  template this: untyped = result
  this = new(SwitchManualStr)
  let root = if root == nil: cast[SwitchManualStr](this) else: cast[SwitchManualStr](root)
  this.io = io
  this.root = root
  this.parent = parent

  block:
    var i: int
    while not this.io.isEof:
      let it = SwitchManualStr_Opcode.read(this.io, this.root, this)
      this.opcodes.add(it)
      inc i

proc fromFile*(_: typedesc[SwitchManualStr], filename: string): SwitchManualStr =
  SwitchManualStr.read(newKaitaiFileStream(filename), nil, nil)

proc read*(_: typedesc[SwitchManualStr_Opcode], io: KaitaiStream, root: KaitaiStruct, parent: SwitchManualStr): SwitchManualStr_Opcode =
  template this: untyped = result
  this = new(SwitchManualStr_Opcode)
  let root = if root == nil: cast[SwitchManualStr](this) else: cast[SwitchManualStr](root)
  this.io = io
  this.root = root
  this.parent = parent

  let codeExpr = encode(this.io.readBytes(int(1)), "ASCII")
  this.code = codeExpr
  block:
    let on = this.code
    if on == "I":
      let bodyExpr = SwitchManualStr_Opcode_Intval.read(this.io, this.root, this)
      this.body = bodyExpr
    elif on == "S":
      let bodyExpr = SwitchManualStr_Opcode_Strval.read(this.io, this.root, this)
      this.body = bodyExpr

proc fromFile*(_: typedesc[SwitchManualStr_Opcode], filename: string): SwitchManualStr_Opcode =
  SwitchManualStr_Opcode.read(newKaitaiFileStream(filename), nil, nil)

proc read*(_: typedesc[SwitchManualStr_Opcode_Intval], io: KaitaiStream, root: KaitaiStruct, parent: SwitchManualStr_Opcode): SwitchManualStr_Opcode_Intval =
  template this: untyped = result
  this = new(SwitchManualStr_Opcode_Intval)
  let root = if root == nil: cast[SwitchManualStr](this) else: cast[SwitchManualStr](root)
  this.io = io
  this.root = root
  this.parent = parent

  let valueExpr = this.io.readU1()
  this.value = valueExpr

proc fromFile*(_: typedesc[SwitchManualStr_Opcode_Intval], filename: string): SwitchManualStr_Opcode_Intval =
  SwitchManualStr_Opcode_Intval.read(newKaitaiFileStream(filename), nil, nil)

proc read*(_: typedesc[SwitchManualStr_Opcode_Strval], io: KaitaiStream, root: KaitaiStruct, parent: SwitchManualStr_Opcode): SwitchManualStr_Opcode_Strval =
  template this: untyped = result
  this = new(SwitchManualStr_Opcode_Strval)
  let root = if root == nil: cast[SwitchManualStr](this) else: cast[SwitchManualStr](root)
  this.io = io
  this.root = root
  this.parent = parent

  let valueExpr = encode(this.io.readBytesTerm(0, false, true, true), "ASCII")
  this.value = valueExpr

proc fromFile*(_: typedesc[SwitchManualStr_Opcode_Strval], filename: string): SwitchManualStr_Opcode_Strval =
  SwitchManualStr_Opcode_Strval.read(newKaitaiFileStream(filename), nil, nil)

