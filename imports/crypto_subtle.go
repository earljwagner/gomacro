// this file was generated by gomacro command: import "crypto/subtle"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	pkg "crypto/subtle"
	. "reflect"
)

func init() {
	Binds["crypto/subtle"] = map[string]Value{
		"ConstantTimeByteEq":	ValueOf(pkg.ConstantTimeByteEq),
		"ConstantTimeCompare":	ValueOf(pkg.ConstantTimeCompare),
		"ConstantTimeCopy":	ValueOf(pkg.ConstantTimeCopy),
		"ConstantTimeEq":	ValueOf(pkg.ConstantTimeEq),
		"ConstantTimeLessOrEq":	ValueOf(pkg.ConstantTimeLessOrEq),
		"ConstantTimeSelect":	ValueOf(pkg.ConstantTimeSelect),
	}
	Types["crypto/subtle"] = map[string]Type{
	}
}