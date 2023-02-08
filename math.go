package gorms

import "golang.org/x/exp/constraints"

func MaxInteger[T constraints.Integer](lhs, rhs T) T {
	if lhs >= rhs {
		return lhs
	} else {
		return rhs
	}
}

func Max0[T constraints.Integer](value T) T {
	return MaxInteger(value, 0)
}

func Max1[T constraints.Integer](value T) T {
	return MaxInteger(value, 1)
}
