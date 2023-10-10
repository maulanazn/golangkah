package memorysimulation

type PaymentContract interface {
	GetCardNumber() string
	GetCardType() string
	GetCardById(id string) string
	GetInfo() int
}

func ShowCardNum(paymentContract PaymentContract) string {
	return paymentContract.GetCardNumber()
}

func ShowCardType(paymentContract PaymentContract) string {
	return paymentContract.GetCardType()
}

func ShowInfo(paymentContract PaymentContract) int {
	return paymentContract.GetInfo()
}

func ShowCardById(paymentContract PaymentContract, id string) string {
	return paymentContract.GetCardById(id)
}
