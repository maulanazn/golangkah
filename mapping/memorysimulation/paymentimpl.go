package memorysimulation

import "fmt"

type Payment struct {
	Id        string
	card_no   string
	card_type string
	user_id   string
}

func (payment Payment) GetCardNumber() string {
	return payment.card_no
}

func (payment Payment) GetCardType() string {
	return payment.card_type
}

func (payment Payment) GetInfo() int {
	result, err := fmt.Printf("Card number: %s \n Type: %s", payment.card_no, payment.card_type)
	if err != nil {
		fmt.Printf("%v", err)
	}

	return result
}
