package belajargolanggenerics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// constraint for type parameter = T interface{} or any === all data type
func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestLength(t *testing.T) {
	var result string = Length[string]("Misno")
	assert.Equal(t, "Misno", result)

	var resultNumber int = Length[int](200)
	assert.Equal(t, 200, resultNumber)
}
