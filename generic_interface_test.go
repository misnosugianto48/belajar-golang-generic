package belajargolanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// combination generic interface + generic function
type GetterSetter[T any] interface {
	GetValue() T
	SetValue(value T)
}

func ChangeValue[T any](param GetterSetter[T], value T) T {
	param.SetValue(value)
	return param.GetValue()
}

// struct implementation,
// generic interface only use for struct generic implementation
type MyData[T any] struct {
	Value T
}

// struct method
func (d *MyData[T]) GetValue() T {
	return d.Value
}

func (d *MyData[T]) SetValue(value T) {
	d.Value = value
}

func TestInterface(t *testing.T) {
	myData := MyData[string]{}

	result := ChangeValue[string](&myData, "Misno")
	assert.Equal(t, "Misno", result)
	assert.Equal(t, "Misno", myData.Value)
}
