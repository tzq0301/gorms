package gorms

const (
	AliasAs = "AS"
)

// AS set alias with "AS"
//
// If reserved word "AS" is not need, use Alias
//
//  AS("user", "u") -> "user AS u"
func AS(v, alias string) string {
	return v + space + AliasAs + space + alias
}

// Alias the value
//
// If reserved word "AS" is need, use AS
//
//  Alias("user", "u") -> "user u"
func Alias(v, alias string) string {
	return v + space + alias
}
