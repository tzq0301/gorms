package gorms

import "strings"

// Note: GORM Order cannot avoid SQL injection
//
// To support some features, some inputs are not escaped, be careful when using user’s input with those methods
// https://gorm.io/docs/security.html#SQL-injection-Methods

const (
	OrderDesc = "DESC"
	OrderAsc  = "ASC"
)

// Desc order
//
//  Desc("id") -> "id DESC"
func Desc(order string) string {
	return order + space + OrderDesc
}

// Asc order
//
//  Asc("id") -> "id ASC"
func Asc(order string) string {
	return order + space + OrderAsc
}

// Orders make a concatenation
//
//  Orders("id DESC", "create_time ASC") -> ""id DESC, create_time ASC""
func Orders(orders ...string) string {
	return strings.Join(orders, comma+space)
}
