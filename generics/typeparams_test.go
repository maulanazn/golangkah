package generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func length[T1 int, T2 string](key T1, value T2) T2 {
	return value
}

func Min[T Number](value1, value2 T) T {
	if value1 < value2 {
		return value1
	} else {
		return value2
	}
}

func IsSame[T comparable](value1, value2 T) bool {
	if value1 == value2 {
		return true
	} else {
		return false
	}
}

func TestLength(t *testing.T) {
	assert.Equal(t, "tes", length(1, "tes"))
	assert.Exactly(t, 19, Min(19, 22))
	assert.True(t, IsSame(19,19))
}
