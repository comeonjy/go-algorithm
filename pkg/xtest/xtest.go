package xtest

import (
	"reflect"
	"testing"
)

func Assert(t *testing.T, src1 interface{}, src2 interface{}) {
	if !reflect.DeepEqual(src1, src2) {
		t.Error("not equal")
	}
}
