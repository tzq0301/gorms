package gorms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAS(t *testing.T) {
	a := assert.New(t)
	a.Equal("user AS u", AS("user", "u"))
}

func TestAlias(t *testing.T) {
	a := assert.New(t)
	a.Equal("user u", Alias("user", "u"))
}
