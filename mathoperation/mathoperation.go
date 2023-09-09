package mathoperation

import (
	"log"
	"math/rand"
	"time"
)

var AvgNum = func(num ...int) int {
	var total int = 0
	var result int = 0
	for _, value := range num {
		total += value
		result = total / len(num)
	}

	return result
}

var FindMin = func(num ...int) int {
	var result int = 0
	for _, value := range num {
		if result < value {
			return value
		}
	}

	return result
}

var FindMax = func(numbers []int, max int) (int, func() []int) {
	var res []int

	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}

	return len(res), func() []int {
		return res
	}
}

var GenerateRandomNumber = func(min, max int) (int, error) {
	var value = rand.New(rand.NewSource(time.Now().Unix())).Int() % (max - min + 1) / min

	if min > max {
		log.Fatal("min must be lower than max")
	}

	return value, nil
}
