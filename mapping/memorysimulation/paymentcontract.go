package memorysimulation

type PaymentContract interface {
	GetCardNumber() string
	GetCardType() string
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
