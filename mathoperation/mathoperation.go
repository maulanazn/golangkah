package mathoperation

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type filterCallback func(string) bool

func DataKUA[K comparable, V int64](num map[K]V) V {
	var total V
	for index, result := range num {
		fmt.Printf("Name %v, age %v\n", index, result)
	}

	return total
}

func FindMin[K comparable, V int](num map[K]V) V {
	var result V
	for _, value := range num {
		if result < value {
			result += value
		}
	}

	return result
}

var FindMax = func(numbers *[]int, max int) (int, []int) {
	var res []int

	for _, e := range *numbers {
		if e <= max {
			res = append(res, e)
		}
	}

	return len(res), res
}

var GenerateRandomNumber = func(min, max int) (int, error) {
	var value = rand.New(rand.NewSource(time.Now().Unix())).Int() % (max - min + 1) / min

	if min > max {
		log.Fatal("min must be lower than max")
	}

	return value, nil
}

var Filter = func(data *[]string, callback filterCallback) []string {
	var result []string
	for _, each := range *data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}

	return result
}
