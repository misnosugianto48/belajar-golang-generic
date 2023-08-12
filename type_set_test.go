package belajargolanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Age int

/*
	1. type set is an interface
*/
/*
	2. type ExampleName interface { int|string|bool| }
*/
/*
	3. ~int( type approximation ) = all type have this data type can compatible or inheritance
*/
type Number interface {
	~int | int16 | int32 | int64 | float32 | float64
}

// all data type number can inheritance to compatile with different function
func Min[T Number](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestMin(t *testing.T) {
	assert.Equal(t, 100, Min[int](100, 200))
	assert.Equal(t, int64(100), Min[int64](int64(100), int64(200)))
	assert.Equal(t, float64(100), Min[float64](float64(100), float64(200)))
	assert.Equal(t, Age(100), Min[Age](Age(100), Age(200)))
}

/*
1. type inference a feature from golang generic
*/
func TestMinTypeInference(t *testing.T) {
	assert.Equal(t, 100, Min(100, 200))
	assert.Equal(t, int64(100), Min(int64(100), int64(200)))
	assert.Equal(t, float64(100), Min(float64(100), float64(200)))
	assert.Equal(t, Age(100), Min(Age(100), Age(200)))
}
