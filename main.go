package main

import (
	"errors"
	memorysimulation "golangkah/memorysim"
)

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "ups"
	}
}

func NewMap(name *string) map[string]string {
	if *name == "" {
		return nil
	} else {
		return map[string]string{
			"name": *name,
		}
	}
}

func Divide(value, initialitator int) (int, error) {
	if initialitator == 0 {
		return 0, errors.New("Initialiator cannot be 0")
	} else {
		result := value / initialitator
		return result, nil
	}
}

func main() {
	memorysimulation.GetPayment()
}
