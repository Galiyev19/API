package dto

import "API/internal/api/validator"

type RegisterUser struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func ValidateEmail(v *validator.Validator, email string) {
	v.Check(email != "", "email", "must be provided")
	v.Check(v.Matches(email, validator.EmailRX), "email", "must be a valid email")
}

func ValidatorPassword(v *validator.Validator, password string) {
	v.Check(password != "", "password", "must be provided")
	v.Check(len(password) >= 8, "password", "must be at least 8 character")
}

func ValidateUser(v *validator.Validator, u *RegisterUser) {
	ValidateEmail(v, u.Email)
	ValidatorPassword(v, u.Password)
}
