package generics_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ValueData[T any] interface {
	GetValue1() T
	SetValue1(param T)
}

func ChangeValue[T any](param ValueData[T], value T) T {
	param.SetValue1(value)
	return param.GetValue1()
}

type Data[T any] struct {
  Value1 T
}

func (d *Data[T]) GetValue1() T {
	return d.Value1
}

func (d *Data[T]) SetValue1(param T) {
	d.Value1 = param
}

func TestGenericStructSuccess(t *testing.T) {
	data := Data[string]{}
	data.SetValue1("maulanazn")

	assert.Equal(t, "maulanazn", data.GetValue1())
}

func TestGenericStructFailed(t *testing.T) {
	data := Data[string]{}
	data.SetValue1("maulanazn")

	assert.NotEqual(t, "maulana", data.GetValue1())
}

func TestGenericInterface(t *testing.T) {
	data := Data[string]{}
	result := ChangeValue(&data, "maulanazn")

	assert.Equal(t, "maulanazn", result)
}
