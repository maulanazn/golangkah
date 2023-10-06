package memorysimulation

import "fmt"

type PaymentContract interface {
	GetId() string
	GetUserId() string
	GetCardNumber() string
	GetCardType() string
	GetAllInfo() Payment
}

func ShowUserId(paymentContr PaymentContract) {
	fmt.Println("User id payment : ", paymentContr.GetUserId())
}

func ShowCardNumber(paymentContr PaymentContract) {
	fmt.Println("Card Num : ", paymentContr.GetCardNumber())
}

func ShowCardType(paymentContr PaymentContract) {
	fmt.Println("Card you use : ", paymentContr.GetCardType())
}

func ShowAllInfo(paymentContr PaymentContract) {
	fmt.Printf("Id user: %s\nUser id: %s\n Card number: %s\nCard Type: %s", paymentContr.GetId(), paymentContr.GetUserId(), paymentContr.GetCardNumber(), paymentContr.GetCardType())
}
