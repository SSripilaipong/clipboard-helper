package rsltutil

import (
	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"
)

// ResultOf2 collapses two rslt.Of values into a tuple, short-circuiting on the first error.
func ResultOf2[T1, T2 any](x1 rslt.Of[T1], x2 rslt.Of[T2]) rslt.Of[tuple.Of2[T1, T2]] {
	if x1.IsErr() {
		return rslt.Error[tuple.Of2[T1, T2]](x1.Error())
	}
	if x2.IsErr() {
		return rslt.Error[tuple.Of2[T1, T2]](x2.Error())
	}
	return rslt.Value(tuple.New2(x1.Value(), x2.Value()))
}
