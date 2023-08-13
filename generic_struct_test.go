package belajargolanggenerics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data[T any] struct {
	First  T
	Second T
}

func TestData(t *testing.T) {
	data := Data[string]{
		First:  "Misno",
		Second: "Sugianto"}

	fmt.Println(data)
}

// generic method, can using named parameter
func (d *Data[_]) SayHello(name string) string {
	return "Hello " + name
}

func (d *Data[T]) ChangeFirst(first T) T {
	d.First = first
	return d.First
}

func TestGenericMethod(t *testing.T) {
	data := Data[string]{
		First:  "Misno",
		Second: "Sugianto",
	}

	assert.Equal(t, "misno", data.ChangeFirst("misno"))
	assert.Equal(t, "Hello misno", data.SayHello("misno"))
}
