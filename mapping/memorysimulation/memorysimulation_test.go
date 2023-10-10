package memorysimulation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("Before cleared")
	m.Run()
	fmt.Println("After cleared")
}

func TestTable(t *testing.T) {
	data := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Valid",
			request:  "mastercard global",
			expected: "mastercard global",
		},
		{
			name:     "NotValid",
			request:  "mastercard indonesia",
			expected: "mastercard indonesia",
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result := CheckValue(d.request)
			assert.Equal(t, d.expected, result, "card type must be valid same")
		})
	}
}
