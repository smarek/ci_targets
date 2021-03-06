// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package test_formats

import (
	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
	"unicode/utf8"
	"strconv"
)

type ExprOpsParens struct {
	_io *kaitai.Stream
	_root *ExprOpsParens
	_parent interface{}
	_f_boolAnd bool
	boolAnd int
	_f_str0To4 bool
	str0To4 string
	_f_boolOr bool
	boolOr int
	_f_fE bool
	fE float64
	_f_fSumToInt bool
	fSumToInt int
	_f_strConcatRev bool
	strConcatRev string
	_f_fPi bool
	fPi float64
	_f_iM13 bool
	iM13 int
	_f_strConcatLen bool
	strConcatLen int
	_f_strConcatToI bool
	strConcatToI int
	_f_i42 bool
	i42 int8
	_f_iSumToStr bool
	iSumToStr string
	_f_boolEq bool
	boolEq int
	_f_str5To9 bool
	str5To9 string
	_f_strConcatSubstr2To7 bool
	strConcatSubstr2To7 string
}
func NewExprOpsParens() *ExprOpsParens {
	return &ExprOpsParens{
	}
}

func (this *ExprOpsParens) Read(io *kaitai.Stream, parent interface{}, root *ExprOpsParens) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	return err
}
func (this *ExprOpsParens) BoolAnd() (v int, err error) {
	if (this._f_boolAnd) {
		return this.boolAnd, nil
	}
	tmp1 := 0
	if  ((false) && (true))  {
		tmp1 = 1
	}
	this.boolAnd = int(tmp1)
	this._f_boolAnd = true
	return this.boolAnd, nil
}
func (this *ExprOpsParens) Str0To4() (v string, err error) {
	if (this._f_str0To4) {
		return this.str0To4, nil
	}
	this.str0To4 = string("01234")
	this._f_str0To4 = true
	return this.str0To4, nil
}
func (this *ExprOpsParens) BoolOr() (v int, err error) {
	if (this._f_boolOr) {
		return this.boolOr, nil
	}
	tmp2 := 0
	if  ((!(false)) || (false))  {
		tmp2 = 1
	}
	this.boolOr = int(tmp2)
	this._f_boolOr = true
	return this.boolOr, nil
}
func (this *ExprOpsParens) FE() (v float64, err error) {
	if (this._f_fE) {
		return this.fE, nil
	}
	this.fE = float64(2.72)
	this._f_fE = true
	return this.fE, nil
}
func (this *ExprOpsParens) FSumToInt() (v int, err error) {
	if (this._f_fSumToInt) {
		return this.fSumToInt, nil
	}
	tmp3, err := this.FPi()
	if err != nil {
		return 0, err
	}
	tmp4, err := this.FE()
	if err != nil {
		return 0, err
	}
	this.fSumToInt = int(int((tmp3 + tmp4)))
	this._f_fSumToInt = true
	return this.fSumToInt, nil
}
func (this *ExprOpsParens) StrConcatRev() (v string, err error) {
	if (this._f_strConcatRev) {
		return this.strConcatRev, nil
	}
	tmp5, err := this.Str0To4()
	if err != nil {
		return "", err
	}
	tmp6, err := this.Str5To9()
	if err != nil {
		return "", err
	}
	this.strConcatRev = string(kaitai.StringReverse(tmp5 + tmp6))
	this._f_strConcatRev = true
	return this.strConcatRev, nil
}
func (this *ExprOpsParens) FPi() (v float64, err error) {
	if (this._f_fPi) {
		return this.fPi, nil
	}
	this.fPi = float64(3.14)
	this._f_fPi = true
	return this.fPi, nil
}
func (this *ExprOpsParens) IM13() (v int, err error) {
	if (this._f_iM13) {
		return this.iM13, nil
	}
	this.iM13 = int(-13)
	this._f_iM13 = true
	return this.iM13, nil
}
func (this *ExprOpsParens) StrConcatLen() (v int, err error) {
	if (this._f_strConcatLen) {
		return this.strConcatLen, nil
	}
	tmp7, err := this.Str0To4()
	if err != nil {
		return 0, err
	}
	tmp8, err := this.Str5To9()
	if err != nil {
		return 0, err
	}
	this.strConcatLen = int(utf8.RuneCountInString(tmp7 + tmp8))
	this._f_strConcatLen = true
	return this.strConcatLen, nil
}
func (this *ExprOpsParens) StrConcatToI() (v int, err error) {
	if (this._f_strConcatToI) {
		return this.strConcatToI, nil
	}
	tmp9, err := this.Str0To4()
	if err != nil {
		return 0, err
	}
	tmp10, err := this.Str5To9()
	if err != nil {
		return 0, err
	}
	tmp11, err := strconv.ParseInt(tmp9 + tmp10, 10, 0)
	if err != nil {
		return 0, err
	}
	this.strConcatToI = int(tmp11)
	this._f_strConcatToI = true
	return this.strConcatToI, nil
}
func (this *ExprOpsParens) I42() (v int8, err error) {
	if (this._f_i42) {
		return this.i42, nil
	}
	this.i42 = int8(42)
	this._f_i42 = true
	return this.i42, nil
}
func (this *ExprOpsParens) ISumToStr() (v string, err error) {
	if (this._f_iSumToStr) {
		return this.iSumToStr, nil
	}
	tmp12, err := this.I42()
	if err != nil {
		return "", err
	}
	tmp13, err := this.IM13()
	if err != nil {
		return "", err
	}
	this.iSumToStr = string(strconv.FormatInt(int64((tmp12 + tmp13)), 10))
	this._f_iSumToStr = true
	return this.iSumToStr, nil
}
func (this *ExprOpsParens) BoolEq() (v int, err error) {
	if (this._f_boolEq) {
		return this.boolEq, nil
	}
	tmp14 := 0
	if false == true {
		tmp14 = 1
	}
	this.boolEq = int(tmp14)
	this._f_boolEq = true
	return this.boolEq, nil
}
func (this *ExprOpsParens) Str5To9() (v string, err error) {
	if (this._f_str5To9) {
		return this.str5To9, nil
	}
	this.str5To9 = string("56789")
	this._f_str5To9 = true
	return this.str5To9, nil
}
func (this *ExprOpsParens) StrConcatSubstr2To7() (v string, err error) {
	if (this._f_strConcatSubstr2To7) {
		return this.strConcatSubstr2To7, nil
	}
	tmp15, err := this.Str0To4()
	if err != nil {
		return "", err
	}
	tmp16, err := this.Str5To9()
	if err != nil {
		return "", err
	}
	this.strConcatSubstr2To7 = string(tmp15 + tmp16[2:7])
	this._f_strConcatSubstr2To7 = true
	return this.strConcatSubstr2To7, nil
}
