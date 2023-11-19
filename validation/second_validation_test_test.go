package validation_test

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/go-playground/validator"
)

func TestNestedStruct(t *testing.T) {
	type Address struct {
		Street string `validate:"required,ascii"`
		Fax int `validate:"required,numeric"`
		Phone int `validate:"required,numeric"`
	}

	type User struct {
		Username string `validate:"required,ascii"`
		Gender string `validate:"required,alpha"`
		Address Address
	}	

	validate := validator.New()

	editUser := User{
		Username: "maulana zn",
		Gender: "male",
		Address: Address{
			Street: "Street 12 of low",
			Fax: 1020299292192,
			Phone: 120202932,
		},
	}

	if err := validate.Struct(&editUser); err != nil {
		t.Error(err.Error())
	}
}

func TestCollection(t *testing.T) {
	validate := validator.New()
	
	validate.RegisterAlias("mustnumber", "required,numeric")
	validate.RegisterAlias("character", "required,ascii")
	validate.RegisterValidation("username", MustValidUsername)

	type School struct {
		Name string `validate:"required"`
	}

	type Address struct {
		Street string `validate:"character"`
		Fax int `validate:"mustnumber"`
		Phone int `validate:"mustnumber"`
	}

	type User struct {
		Username string `validate:"required,username"`
		Gender string `validate:"required,alpha"`
		Addresses []Address `validate:"required,dive"`
		Hobbies []string `validate:"required,dive,min=1"`
		Schools map[string]School `validate:"dive,keys,required,min=1,endkeys,dive"`
		Wallet map[string]int `validate:"dive,keys,required,min=2,endkeys,required,gt=900"`
	}	

	editUser := User{
		Username: "MAULANA",
		Gender: "male",
		Addresses: []Address{
			{
			Street: "Street 12 of low",
			Fax: 123,
			Phone: 123,
			},
		},
		Hobbies: []string{"Writing", "Cycling", "Main TikTok -- Naima", "Masak"},
		Schools: map[string]School{
			"SD": {
				Name: "gtw",
			},
			"SMP": {
				Name: "gtw",
			},
			"SMA": {
				Name: "gtw",
			},
		},
		Wallet: map[string]int{
			"MANDIRI": 1000,
			"BCA": 1000,
		},
	}

	if err := validate.Struct(&editUser); err != nil {
		t.Error(err.Error())
	}
}

func MustValidUsername(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	if ok {
		if value != strings.ToUpper(value) {
			return false
		}

		if len(value) < 5 {
			return false
		}
	}

	return true
}

func MustValidPin(field validator.FieldLevel) bool {
	regexNumber := regexp.MustCompile("^[0-9]+$")
	length, err := strconv.Atoi(field.Param())
	if err != nil {
		panic(err)
	}

	value := field.Field().String()
	if !regexNumber.MatchString(value) {
		return false
	}	

	return len(value) == length
}

func MustEqualsIgnoreCase(field validator.FieldLevel) bool {
	value, _, _, ok := field.GetStructFieldOK2()
	if !ok {
		panic("Value is not ok")
	}

	firstValue := strings.ToUpper(field.Field().String())
	secondValue := strings.ToUpper(value.String())

	return firstValue == secondValue
}

func TestFieldParam(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("pin", MustValidPin)
	validate.RegisterValidation("field_equals_ignore_case", MustEqualsIgnoreCase)

	type SSOLogin struct {
		Username string `validate:"required,field_equals_ignore_case=Email"`
		Email string `validate:"required,email"`
		Pin string `validate:"required,pin=6"`
	}

	loginsso := SSOLogin{
		Username: "maulanazn@mail.com",
		Email: "maulanazn@mail.com",
		Pin: "123456",
	}

	if err := validate.Struct(&loginsso); err != nil {
		t.Error(err.Error())
		return
	}
	
	t.Log("your pin ", loginsso.Pin)
}

type RegisterRequest struct {
	Username string `validate:"required"`
	Email string `validate:"required,email"`
	Phone string `validate:"required,numeric"`
	Password string `validate:"required"`
}

func MustValidUsernameInterception(structLevel validator.StructLevel) {
	registerRequest := structLevel.Current().Interface().(RegisterRequest)

	if registerRequest.Username == registerRequest.Email || registerRequest.Username == registerRequest.Phone {
		log.Println("success")
	} else {
		structLevel.ReportError(registerRequest.Username, "Username", "Username", "username", "")	
	}
}

func TestStructLevel(t *testing.T) {
	validate := validator.New()
	validate.RegisterStructValidation(MustValidUsernameInterception, RegisterRequest{})

	registerRequest := RegisterRequest{
		Username: "123456",
		Email: "maulanazn@example.com",
		Phone: "123456",
		Password: "maulanazn123",
	}

	if err := validate.Struct(&registerRequest); err != nil {
		t.Error(err.Error())
	}

	t.Log("username ", registerRequest.Username)	
}
