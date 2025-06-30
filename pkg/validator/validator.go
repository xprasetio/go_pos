package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(s interface{}) error {
	v := validator.New()
	_ = v.RegisterValidation("passwordcomplex", passwordComplex)
	return v.Struct(s)
}

func ValidateRegisterRequest(req interface{}) error {
	return ValidateStruct(req)
}

func ValidateLoginRequest(req interface{}) error {
	return ValidateStruct(req)
}

func passwordComplex(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	// Minimal 1 huruf besar, 1 angka, 1 karakter spesial
	upper := regexp.MustCompile(`[A-Z]`)
	digit := regexp.MustCompile(`[0-9]`)
	special := regexp.MustCompile(`[!@#\$%\^&\*\-_+=\[\]{}|;:'",.<>/?]`)
	return upper.MatchString(password) && digit.MatchString(password) && special.MatchString(password)
}

func FormatValidationError(err error) map[string]string {
	detail := map[string]string{}
	if verrs, ok := err.(validator.ValidationErrors); ok {
		for _, verr := range verrs {
			field := verr.Field()
			switch field {
			case "Name":
				if verr.Tag() == "required" {
					detail["error"] = "Field 'name' wajib diisi"
				}
			case "Date":
				if verr.Tag() == "required" {
					detail["error"] = "Field 'date' wajib diisi"
				}
			case "CategoryVillaID":
				if verr.Tag() == "required" {
					detail["error"] = "Field 'category_villa_id' wajib diisi"
				}
			case "RoomIdealCapacity":
				if verr.Tag() == "required" {
					detail["error"] = "Field 'room_ideal_capacity' wajib diisi"
				}
			case "NumberOfRoom":
				if verr.Tag() == "required" {
					detail["error"] = "Field 'number_of_room' wajib diisi"
				}
			case "WeekdayPattern":
				if verr.Tag() == "required" {
					detail["error"] = "Field 'weekday_pattern' wajib diisi"
				}
			case "BankName":
				if verr.Tag() == "required" {
					detail["error"] = "Field 'bank_name' wajib diisi"
				}
			case "BankAccName":
				if verr.Tag() == "required" {
					detail["error"] = "Field 'bank_acc_name' wajib diisi"
				}
			case "BankAccNumber":
				if verr.Tag() == "required" {
					detail["error"] = "Field 'bank_acc_number' wajib diisi"
				}
			case "Email":
				if verr.Tag() == "required" {
					detail["email"] = "Field 'email' wajib diisi"
				} else if verr.Tag() == "email" {
					detail["error"] = "Format email tidak valid"
				}
			case "Password":
				if verr.Tag() == "required" {
					detail["error"] = "Field 'password' wajib diisi"
				} else if verr.Tag() == "min" {
					detail["error"] = "Password minimal 8 karakter"
				} else if verr.Tag() == "passwordcomplex" {
					detail["error"] = "Password harus mengandung huruf besar, angka, dan karakter spesial"
				}
			}
		}
	}
	return detail
}
