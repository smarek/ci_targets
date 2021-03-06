import kaitai_struct_nim_runtime
import options

type
  ExprIfIntOps* = ref object of KaitaiStruct
    `skip`*: seq[byte]
    `it`*: int16
    `boxed`*: int16
    `parent`*: KaitaiStruct
    `isEqPrimInst`*: bool
    `isEqBoxedInst`*: bool

proc read*(_: typedesc[ExprIfIntOps], io: KaitaiStream, root: KaitaiStruct, parent: KaitaiStruct): ExprIfIntOps

proc isEqPrim*(this: ExprIfIntOps): bool
proc isEqBoxed*(this: ExprIfIntOps): bool

proc read*(_: typedesc[ExprIfIntOps], io: KaitaiStream, root: KaitaiStruct, parent: KaitaiStruct): ExprIfIntOps =
  template this: untyped = result
  this = new(ExprIfIntOps)
  let root = if root == nil: cast[ExprIfIntOps](this) else: cast[ExprIfIntOps](root)
  this.io = io
  this.root = root
  this.parent = parent

  let skipExpr = this.io.readBytes(int(2))
  this.skip = skipExpr
  if true:
    let itExpr = this.io.readS2le()
    this.it = itExpr
  if true:
    let boxedExpr = this.io.readS2le()
    this.boxed = boxedExpr

proc isEqPrim(this: ExprIfIntOps): bool = 
  if this.isEqPrimInst != nil:
    return this.isEqPrimInst
  let isEqPrimInstExpr = bool(this.it == 16705)
  this.isEqPrimInst = isEqPrimInstExpr
  if this.isEqPrimInst != nil:
    return this.isEqPrimInst

proc isEqBoxed(this: ExprIfIntOps): bool = 
  if this.isEqBoxedInst != nil:
    return this.isEqBoxedInst
  let isEqBoxedInstExpr = bool(this.it == this.boxed)
  this.isEqBoxedInst = isEqBoxedInstExpr
  if this.isEqBoxedInst != nil:
    return this.isEqBoxedInst

proc fromFile*(_: typedesc[ExprIfIntOps], filename: string): ExprIfIntOps =
  ExprIfIntOps.read(newKaitaiFileStream(filename), nil, nil)

