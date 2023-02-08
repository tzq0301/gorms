package gorms

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPageParamConvert(t *testing.T) {
	Convey("PageParamConvert", t, func() {
		Convey("1-10", func() {
			offset, limit := PageParamConvert(1, 10)
			So(offset, ShouldEqual, 0)
			So(limit, ShouldEqual, 10)
		})
		Convey("2-10", func() {
			offset, limit := PageParamConvert(2, 10)
			So(offset, ShouldEqual, 10)
			So(limit, ShouldEqual, 10)
		})
		Convey("0-10", func() {
			offset, limit := PageParamConvert(0, 10)
			So(offset, ShouldEqual, 0)
			So(limit, ShouldEqual, 10)
		})
		Convey("1-0", func() {
			offset, limit := PageParamConvert(1, 0)
			So(offset, ShouldEqual, 0)
			So(limit, ShouldEqual, 1)
		})
	})
}

func TestPageParamConvertWith(t *testing.T) {
	Convey("PageParamConvertWith", t, func() {
		Convey("bound", func() {
			lowerBound := 5
			upperBound := 50
			pageSizeBoundFn := func(pageSize int) int {
				if pageSize > upperBound {
					return upperBound
				} else if pageSize < lowerBound {
					return lowerBound
				} else {
					return pageSize
				}
			}
			Convey("normal", func() {
				pageNum, pageSize := 2, 10
				offset, limit := PageParamConvertWith(pageNum, pageSize, pageSizeBoundFn)
				So(offset, ShouldEqual, (pageNum-1)*pageSize)
				So(limit, ShouldEqual, pageSize)
			})
			Convey("lower", func() {
				pageNum, pageSize := 2, 3
				offset, limit := PageParamConvertWith(pageNum, pageSize, pageSizeBoundFn)
				So(offset, ShouldEqual, (pageNum-1)*lowerBound)
				So(limit, ShouldEqual, lowerBound)
			})
			Convey("upper", func() {
				pageNum, pageSize := 2, 100
				offset, limit := PageParamConvertWith(pageNum, pageSize, pageSizeBoundFn)
				So(offset, ShouldEqual, (pageNum-1)*upperBound)
				So(limit, ShouldEqual, upperBound)
			})
		})
	})
}
