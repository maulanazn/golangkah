package memorysimulation

type Payment struct {
	Id        string
	card_no   string
	card_type string
	user_id   string
}

func (payment *Payment) GetId() string {
	return payment.Id
}

func (payment *Payment) GetUserId() string {
	return payment.user_id
}

func (payment *Payment) GetCardNumber() string {
	return payment.card_no
}

func (payment *Payment) GetCardType() string {
	return payment.card_type
}

func (payment *Payment) GetAllInfo() Payment {
	return *payment
}
