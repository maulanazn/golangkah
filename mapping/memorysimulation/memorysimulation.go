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

type Payment struct {
    Id string
    card_no string
    card_type string
    user_id string
}

func (payment Payment) GetId() string {
    return payment.Id
}

func (payment Payment) GetUserId() string {
    return payment.user_id
}

func (payment Payment) GetCardNumber() string {
    return payment.card_no
}

func (payment Payment) GetCardType() string {
    return payment.card_type
}

func (payment Payment) GetAllInfo() Payment {
    return payment
}

func GetPayment() {
    var userPayment Payment = Payment{
        Id: "1023ddss-dfdd-dfdf",
        user_id: "01002-2322-sdfs",
        card_no: "lkasldjflkasjf",
        card_type: "mastercard global",
    }
    
    ShowAllInfo(userPayment)
}
