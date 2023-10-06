package memorysimulation

func GetPayment() {
	var userPayment Payment = Payment{
		Id:        "1023ddss-dfdd-dfdf",
		user_id:   "01002-2322-sdfs",
		card_no:   "lkasldjflkasjf",
		card_type: "mastercard global",
	}

	ShowCardNumber(&userPayment)
}
