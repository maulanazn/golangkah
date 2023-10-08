package memorysimulation

func GetPayment() {
	var userPayment Payment = Payment{
		Id:        "1023ddss-dfdd-dfdf",
		card_no:   "01002-2322-sdfs",
		user_id:   "lkasldjflkasjf",
		card_type: "mastercard global",
	}

	ShowInfo(userPayment)
}
