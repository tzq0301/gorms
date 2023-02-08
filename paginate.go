package gorms

import (
	"golang.org/x/exp/constraints"
	"gorm.io/gorm"
)

// Paginate convert pageNum and pageSize to offset and limit
//
//  Paginate(3, 10) -> db.Offset(20).Limit(10)
//
// If offset or limit is less than 1, then Paginate will fix it
//
// If some constraints are required for pageSize, please use PaginateWith
func Paginate[T constraints.Integer](pageNum, pageSize T) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset, limit := PageParamConvert(pageNum, pageSize)
		return db.Offset(int(offset)).Limit(int(limit))
	}
}

// PaginateWith convert pageNum and pageSize to offset and limit with the constraint of pageSize
func PaginateWith[T constraints.Integer](pageNum, pageSize T, pageSizeFn func(T) T) func(*gorm.DB) *gorm.DB {
	return Paginate(pageNum, pageSizeFn(pageSize))
}

// PageParamConvert convert the num and size of a page to the param offset and limit
//
// PageParamConvert will fix the value of the pageNum and pageSize if the value is less than 1
func PageParamConvert[T constraints.Integer](pageNum, pageSize T) (offset, limit T) {
	pageNum, pageSize = Max1(pageNum), Max1(pageSize)
	offset, limit = (pageNum-1)*pageSize, pageSize
	return
}

// PageParamConvertWith convert the num and size of a page to the param offset and limit with the adjust function for page size
func PageParamConvertWith[T constraints.Integer](pageNum, pageSize T, pageSizeFn func(T) T) (offset, limit T) {
	return PageParamConvert(pageNum, pageSizeFn(pageSize))
}
