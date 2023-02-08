# gorms - util for gorm

Looking forward to your pull requests!

## Alias

```go
gorms.AS("user", "u")    // "user AS u"
gorms.Alias("user", "u") // "user u"
```

## Connection

```go
db, err := gorms.InMemorySqlite()
db, err := gorms.InMemorySqliteWithInitSQL("CREATE TABLE ...")
```

## Like

```go
gorms.PercentagePrefix("Tim") // "%Tim"
gorms.PercentageSuffix("Tim") // "Tim%"
gorms.PercentageAround("Tim") // "%Tim%"

gorms.Like(gorms.column, pattern)         // db.Where("? LIKE ?", column, %pattern%)
gorms.BeforeLike(gorms.column, pattern)   // db.Where("? LIKE ?", column, %pattern)
gorms.AfterLike(gorms.column, pattern)    // db.Where("? LIKE ?", column, pattern%)
```

## Order

```go
gorms.Desc("id") // "id DESC"
gorms.Asc("id")  // "id ASC"
gorms.Orders("id DESC", "create_time ASC") // "id DESC, create_time ASC"
```

## Paginate

```go
gorms.Paginate(3, 10) // db.Offset(20).Limit(10)
gorms.PageParamConvert(2, 10) // offset=10 & limit=10

// pageSizeBoundFn is used for the constraints of pageSize
// check paginate_test.go for usage in detail 
gorms.PaginateWith(pageNum, pageSize, pageSizeBoundFn)
gorms.PageParamConvertWith(pageNum, pageSize, pageSizeBoundFn)
```

## String

```go
gorms.SignedIntegerToString(123) // "123"
```