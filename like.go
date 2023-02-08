package gorms

import "gorm.io/gorm"

const (
	WildCardPercentage = "%"
	WildCardUnderLine  = "_"
)

// PercentagePrefix add "%" before the value
//
//  PercentagePrefix("Tim") -> "%Tim"
func PercentagePrefix(v string) string {
	return WildCardPercentage + v
}

// PercentageSuffix add "%" after the value
//
//  PercentageSuffix("Tim") -> "Tim%"
func PercentageSuffix(v string) string {
	return v + WildCardPercentage
}

// PercentageAround add % before and after the string
//
//  PercentageAround("Tim") -> "%Tim%"
//
// PercentageAround use string concatenation rather than fmt.Sprintf("%%%s%%", v), which is confusing and slow (no need to use format)
func PercentageAround(v string) string {
	return WildCardPercentage + v + WildCardPercentage
}

// Like generate `column LIKE pattern`
//
//  Like(column, pattern) -> db.Where("? LIKE ?", column, pattern)
//
// The order of the parameters (column & pattern) is consistent with the order of LIKE clause
func Like(column, pattern string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("? LIKE ?", column, pattern)
	}
}
