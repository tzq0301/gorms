package gorms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAsc(t *testing.T) {
	a := assert.New(t)
	a.Equal("Hello ASC", Asc("Hello"))
}

func TestDesc(t *testing.T) {
	a := assert.New(t)
	a.Equal("Hello DESC", Desc("Hello"))
}

func TestOrders(t *testing.T) {
	a := assert.New(t)
	a.Equal("Hello, HelloWorld", Orders("Hello", "HelloWorld"))
	a.Equal("Hello DESC, World ASC", Orders("Hello"+" DESC", "World"+" ASC"))
}
