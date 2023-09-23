package fund_go_generics

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func MultiplePrintOut[T1, T2 any](name T1, age T2) (T1, T2) {
	fmt.Println(name, age)
	return name, age
}

func MultiPrint[T1 any, T2 any](param1 T1, param2 T2) {
	fmt.Println(param1)
	fmt.Println(param2)
}

func TestMultiple(t *testing.T) {
	name, age := MultiplePrintOut[string, int]("Putri", 21)
	assert.Equal(t, "Putri", name)
	assert.Equal(t, 21, age)

	name, age = MultiplePrintOut[string, int]("Chika", 26)
	assert.Equal(t, "Chika", name)
	assert.Equal(t, 26, age)

	MultiPrint("Putri", 21)
	MultiPrint("Chika", 26)
}
