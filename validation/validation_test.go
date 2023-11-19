package validation_test

import (
	"testing"

	"github.com/go-playground/validator"
)

func TestVarValidate(t *testing.T) {
	var validate *validator.Validate = validator.New()
	user := "maulanazn"

	if err := validate.Var(&user, "required"); err != nil {
		t.Error(err.Error())
		return
	}

	t.Logf("%s", user)
}

func TestMultipleVar(t *testing.T) {
	var validate *validator.Validate = validator.New()
	password := "maulanazn"
	confirmpassword := "maulanazn"

	if err := validate.VarWithValue(&password, &confirmpassword, "eqfield"); err != nil {
		t.Error(err.Error())
		return
	}

	t.Logf("%s", "yeahhh")
}

func TestMultipleTag(t *testing.T) {
	var validate *validator.Validate = validator.New()
	phone := "081028"

	if err := validate.Var(&phone, "eqfield,numeric"); err != nil {
		t.Error(err.Error())
		return
	}

	t.Logf("%s", phone)
}

func TestValidateStruct(t *testing.T) {
	type LoginRequest struct {
		Email string `validate:"required,email"`
		Password string `validate:"required,min=8,max=20"`
	}

	var validate *validator.Validate = validator.New()
	
	register := &LoginRequest{
		Email: "maulanazn@mail.com",
		Password: "maulanazn123",
	}

	err := validate.Struct(*register)
	if err != nil {
		t.Error(err.Error())
	}
	
	t.Log("success")
}

func TestValidationError(t *testing.T) {
	type LoginRequest struct {
		Email string `validate:"required,email"`
		Password string `validate:"required,min=8,max=20"`
	}

	var validate *validator.Validate = validator.New()
	
	register := &LoginRequest{
		Email: "maulanaznmail.com",
		Password: "maulanazn123",
	}

	err := validate.Struct(*register)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			if fieldError.Param() == "" {
				t.Error("error on ", fieldError.Field(), " must be ", fieldError.Tag())	
				return
			}
			t.Error("error on ", fieldError.Field(), " with error on tag ", fieldError.Tag(), " must have ", fieldError.Param())	
		}
	}
	
	t.Log("success")
}

func TestCrossValidation(t *testing.T) {
	type RegisterRequest struct {
		Username string `validate:"required,ascii"`
		Email string `validate:"required,email"`
		Password string `validate:"required,min=8,max=20"`
		ConfirmPassword string `validate:"eqfield=Password"`
	}

	var validate *validator.Validate = validator.New()
	
	register := &RegisterRequest{
		Username: "john blake",
		Email: "johnblake19@mail.com",
		Password: "maulanazn123",
		ConfirmPassword: "maulanazn123",
	}

	err := validate.Struct(*register)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			if fieldError.Param() == "" {
				t.Error("error on ", fieldError.Field(), " must be ", fieldError.Tag())	
				return
			}
			t.Error("error on ", fieldError.Field(), " with error on tag ", fieldError.Tag(), " must have ", fieldError.Param())	
		}
	}
	
	t.Log("success")
}
