package gorms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPercentagePrefix(t *testing.T) {
	a := assert.New(t)
	a.Equal("%123", PercentagePrefix("123"))
}

func TestPercentageAround(t *testing.T) {
	a := assert.New(t)
	a.Equal("%123%", PercentageAround("123"))
}

func TestPercentageSuffix(t *testing.T) {
	a := assert.New(t)
	a.Equal("123%", PercentageSuffix("123"))
}
