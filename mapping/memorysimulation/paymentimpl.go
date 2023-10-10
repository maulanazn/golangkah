package memorysimulation

import (
	"fmt"
)

type Payment []struct {
	Id        string
	card_no   string
	card_type string
	user_id   string
}

func (payment Payment) GetCardNumber() string {
	for _, up := range payment {
		return up.card_no
	}

	return ""
}

func (payment Payment) GetCardType() string {
	for _, up := range payment {
		return up.card_type
	}

	return ""
}

func (payment Payment) GetInfo() int {
	for _, up := range payment {
		result, err := fmt.Printf("Card number: %s \n Type: %s", up.card_no, up.card_type)
		if err != nil {
			fmt.Printf("%v", err)
		} else {
			return result
		}
	}

	return 0
}

func (payment Payment) GetCardById(id string) string {
	return id
}
