package generics_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Number interface {
	int | int64 | float64
}

func SetAndGetNoInline[T Number](param T) T {
	result := param
	return result
}

func SetAndGetInline[T interface{int | int64 | float64}](param T) T {
	result := param
	return result
}

func TestNonInline(t *testing.T) {
	assert.Exactly(t, 12, SetAndGetNoInline(12))
}

func TestInline(t *testing.T) {
	assert.Exactly(t, 12, SetAndGetInline(12))
}
