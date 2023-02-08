package gorms

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSignedIntegerToString(t *testing.T) {
	Convey("SignedIntegerToString", t, func() {
		Convey("int", func() {
			So("-1", ShouldEqual, SignedIntegerToString(-1))
			So("0", ShouldEqual, SignedIntegerToString(0))
			So("1", ShouldEqual, SignedIntegerToString(1))
			So("10", ShouldEqual, SignedIntegerToString(10))
			So("433", ShouldEqual, SignedIntegerToString(433))
		})
		Convey("int8", func() {
			So("-1", ShouldEqual, SignedIntegerToString(int8(-1)))
			So("0", ShouldEqual, SignedIntegerToString(int8(0)))
			So("1", ShouldEqual, SignedIntegerToString(int8(1)))
			So("10", ShouldEqual, SignedIntegerToString(int8(10)))
			So("123", ShouldEqual, SignedIntegerToString(int8(123)))
		})
		Convey("int16", func() {
			So("-1", ShouldEqual, SignedIntegerToString(int16(-1)))
			So("0", ShouldEqual, SignedIntegerToString(int16(0)))
			So("1", ShouldEqual, SignedIntegerToString(int16(1)))
			So("10", ShouldEqual, SignedIntegerToString(int16(10)))
			So("433", ShouldEqual, SignedIntegerToString(int16(433)))
		})
		Convey("int32", func() {
			So("-1", ShouldEqual, SignedIntegerToString(int32(-1)))
			So("0", ShouldEqual, SignedIntegerToString(int32(0)))
			So("1", ShouldEqual, SignedIntegerToString(int32(1)))
			So("10", ShouldEqual, SignedIntegerToString(int32(10)))
			So("433", ShouldEqual, SignedIntegerToString(int32(433)))
		})
		Convey("int64", func() {
			So("-1", ShouldEqual, SignedIntegerToString(int64(-1)))
			So("0", ShouldEqual, SignedIntegerToString(int64(0)))
			So("1", ShouldEqual, SignedIntegerToString(int64(1)))
			So("10", ShouldEqual, SignedIntegerToString(int64(10)))
			So("433", ShouldEqual, SignedIntegerToString(int64(433)))
		})
	})
}
