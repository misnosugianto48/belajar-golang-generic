package belajargolanggenerics

import (
	"fmt"
	"testing"
)

// multiple type parameter must have different named constraint
func MultipleParameter[T1 any, T2 any](param1 T1, param2 T2) {
	fmt.Println(param1)
	fmt.Println(param2)
}

func TestMultipleParameter(t *testing.T) {
	MultipleParameter[string, int]("Misno", 22)
	MultipleParameter[int, string](23, "Aminah")
}
