package memorysimulation

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Before cleared")
	m.Run()
	fmt.Println("After cleared")
}

// func TestTable(t *testing.T) {
// 	data := []struct {
// 		name     string
// 		request  string
// 		expected string
// 	}{
// 		{
// 			name:     "Valid",
// 			request:  "mastercard global",
// 			expected: "mastercard global",
// 		},
// 		{
// 			name:     "NotValid",
// 			request:  "mastercard indonesia",
// 			expected: "mastercard indonesia",
// 		},
// 	}

// 	for _, d := range data {
// 		t.Run(d.name, func(t *testing.T) {
// 			result := CheckValue(d.request)
// 			assert.Equal(t, d.expected, result, "card type must be valid same")
// 		})
// 	}
// }

func BenchmarkTableHello(b *testing.B) {
	result := []struct {
		name    string
		message string
	}{
		{
			name:    "World",
			message: "Hello World",
		},
	}

	for _, tes := range result {
		b.Run(tes.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fmt.Println(tes.message, "number : ", i)
			}
		})
	}
}
