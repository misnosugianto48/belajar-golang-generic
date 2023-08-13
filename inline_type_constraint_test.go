package belajargolanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func FindMin[T interface{ int | int64 | float64 }](first, second T) T {
	if second > first {
		return first
	}

	return second
}

func TestFinfINline(t *testing.T) {
	assert.Equal(t, int64(100), FindMin[int64](int64(100), int64(200)))
}

// generic type parameter
func GetFirst[T []E, E any](slice T) E {
	first := slice[0]
	return first
}

func TestGetFirst(t *testing.T) {
	names := []string{
		"Misno", "Sugianto",
	}
	first := GetFirst[[]string, string](names)
	assert.Equal(t, "Misno", first)
}
