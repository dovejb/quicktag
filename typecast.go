package quicktag

import (
	. "reflect"
	. "unsafe"
)

type emptyInterface struct {
	pt Pointer
	pv Pointer
}

func PointerOfType(t Type) Pointer {
	p := *(*emptyInterface)(Pointer(&t))
	return p.pv
}

func TypeCast(src interface{}, dstType Type) (dst interface{}) {
	srcType := TypeOf(src)
	eface := *(*emptyInterface)(Pointer(&src))
	if srcType.Kind() == Ptr {
		eface.pt = PointerOfType(PtrTo(dstType))
	} else {
		eface.pt = PointerOfType(dstType)
	}
	dst = *(*interface{})(Pointer(&eface))
	return
}
