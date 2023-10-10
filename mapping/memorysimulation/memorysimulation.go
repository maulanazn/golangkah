package memorysimulation

import (
	"fmt"
)

func GetPayment() {
	userPayment := []Payment{
		{
			{
				"asdf",
				"a1022033",
				"mastercard global",
				"821832",
			},
		},
	}

	fmt.Println("Hello world")
	ShowCardById(userPayment, "lkjlasdf")
}
