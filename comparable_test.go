package fund_go_generics

import "testing"

func Equal[T comparable](param1, param2 T) bool {
	if param1 == param2 {
		return true
	} else {
		return false
	}
}

func Test_Comparable(t *testing.T) {
	assertEqual := Equal[int](1, 1)
	assertEqual2 := Equal[string]("Putri", "Putri")
	assertEqual3 := Equal[float64](1.2, 1.2)
	assertEqual4 := Equal[bool](true, false)

	if assertEqual == true {
		t.Log("Equal")
	}

	if assertEqual2 == true {
		t.Log("Equal")
	}

	if assertEqual3 == true {
		t.Log("Equal")
	}

	if assertEqual4 == false {
		t.Log("Not Equal")
	}
}
