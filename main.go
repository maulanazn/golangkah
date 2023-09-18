package main

import (
	"fmt"
	"golangkah/mathoperation"
)

func main() {
	ints := map[string]int64{
		"maulana": 20,
		"lily":    18,
	}
	datanum := map[string]int{
		"satu": 1,
		"dua":  2,
	}
	mathoperation.DataKUA[string, int64](ints)
	fmt.Printf("%v", mathoperation.FindMin[string, int](datanum))
}
