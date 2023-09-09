package main

import (
	"fmt"
	"golangkah/mathoperation"
	"log"
)

func main() {
	var groupOfNumber []int = []int{9, 12, 24, 33, 45, 51, 65, 76, 88, 99, 150, 72, 83, 82, 24, 23, 35, 27, 32, 21, 20, 100, 125}
	var max = 90
	var howMax, getNumbers = mathoperation.FindMax(groupOfNumber, max)
	var initNum = getNumbers()
	generateRandom, err := mathoperation.GenerateRandomNumber(1, 10)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(generateRandom)
	}

	fmt.Println(mathoperation.AvgNum(groupOfNumber...))
	fmt.Print(howMax, initNum)
}
