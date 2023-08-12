package belajargolanggenerics

import (
	"fmt"
	"testing"
)

/*
1. generic can use for type
*/
type Bag[T any] []T

func PrintBag[T any](bag Bag[T]) {
	for _, v := range bag {
		fmt.Println(v)
	}
}

func TestBagString(t *testing.T) {
	names := Bag[string]{"misno", "sugianto"}
	PrintBag[string](names)
}
