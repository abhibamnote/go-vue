package models

import (
	"gorm.io/gorm"
	"github.com/go-playground/validator/v10"
)

type User struct {
	gorm.Model
	FirstName   string
	LastName	string
	Email   	string		`gorm:"size:191;uniqueIndex"`
	Phone 		*string
	CountryCode *string
	Password 	string
}

var validate = validator.New() 

type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

func ValidateStruct[T any](payload T) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

type CreateUserSchema struct {
	FirstName	string	`json:"firstName" validate:"required"`
	LastName	string	`json:"lastName" validate:"required"`
	Email		string	`json:"email" validate:"required"`
	Phone		string	`json:"phone,omitempty"`
	CountryCode	string 	`json:"countryCode,omitempty"`
	Password	string	`json:"password" validate:"required"`
}

type UpdateUserSchema struct {
	FirstName	string	`json:"firstName,omitempty"`
	LastName	string	`json:"lastName,omitempty"`
	Email		string	`json:"email,omitempty"`
	Phone		string	`json:"phone,omitempty"`
	CountryCode	string 	`json:"countryCode,omitempty"`
	Password	string	`json:"password,omitempty"`
}

type LoginUserSchema struct {
	Email	 string		`json:"email" validate:"required"`
	Password string		`json:"password"`
}