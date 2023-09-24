package datamap

import (
	"fmt"
	"golangkah/mathoperation"
)

func Datamap() {
	var arrNum = map[string]int64{
		"satu": 1,
		"dua":  2,
	}
	fmt.Println(mathoperation.DataKUA(arrNum))
}
