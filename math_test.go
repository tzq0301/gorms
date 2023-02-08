package gorms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Max0(t *testing.T) {
	a := assert.New(t)

	a.Equal(1, Max0(1))
	a.Equal(0, Max0(0))
	a.Equal(0, Max0(-1))
}

func Test_Max1(t *testing.T) {
	a := assert.New(t)

	a.Equal(2, Max1(2))
	a.Equal(1, Max1(1))
	a.Equal(1, Max1(0))
}
