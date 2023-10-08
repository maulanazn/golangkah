package memorysimulation

import (
	"fmt"
	"testing"
)

func TestSuccessGetPayment(t *testing.T) {
	var userPayment Payment = Payment{
		Id:        "1023ddss-dfdd-dfdf",
		card_no:   "01002-2322-sdfs",
		user_id:   "lkasldjflkasjf",
		card_type: "mastercard global",
	}

	result := ShowCardNum(userPayment)

	if result != "01002-2322-sdfs" {
		t.Fatal("Wrong card number")
	}

	fmt.Println(result)
}

func TestFailedGetPayment(t *testing.T) {
	var userPayment Payment = Payment{
		Id:        "1023ddss-dfdd-dfdf",
		card_no:   "01002-2322-adbc",
		user_id:   "lkasldjflkasjf",
		card_type: "mastercard global",
	}

	result := ShowCardNum(userPayment)

	if result == "01002-2322-sdfs" {
		t.Fatal("Right card number, enter the wrong card")
	}

	fmt.Println(result)
}
