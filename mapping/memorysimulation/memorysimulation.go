package memorysimulation

import "fmt"

func SuperMem() {
	var numA = 12
	var numB *int = &numA

	fmt.Println("value of numb", *numB)
	fmt.Println("memory of numb", numB)
}
