package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Print[T any](param T) T {
	fmt.Println(param)
	return param
}

func Test_Type(t *testing.T) {
	var name string = Print[string]("Putri")
	assert.Equal(t, "Putri", name)

	var umur int = Print[int](21)
	assert.Equal(t, 21, umur)

}
