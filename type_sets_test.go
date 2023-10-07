package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Age int

type Number interface {
	~int | int8 | int16 | int32 | int64
}

func Min[T Number](param1, param2 T) T {
	if param1 < param2 {
		return param1
	} else {
		return param2
	}
}

func TestNumber(t *testing.T) {
	minimal := Min[int](1, 2)
	minn := Min[Age](1, 2)
	assert.Equal(t, 1, minimal)
	assert.Equal(t, Age(1), minn)
}
