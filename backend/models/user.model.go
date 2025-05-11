package models

import (
	"time"
	"gorm.io/gorm"
	"github.com/go-playground/validator/v10"
)

type User struct {
	ID			uint		`gorm:"primaryKey"`
	FirstName   string
	LastName	string
	Email   	string		`gorm:"size:191;uniqueIndex"`
	Phone 		*string
	CountryCode *string
	Password 	string
	CreatedAt	time.Time
	UpdatedAt	time.Time		
	*gorm.Model
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
	FirstName	string	`json:"first_name" validate:"required"`
	LastName	string	`json:"last_name" validate:"required"`
	Email		string	`json:"email" validate:"required"`
	Phone		string	`json:"phone,omitempty"`
	CountryCode	string 	`json:"country_code,omitempty"`
	Password	string	`json:"password" validate:"required"`
}

type UpdateUserSchema struct {
	FirstName	string	`json:"first_name,omitempty"`
	LastName	string	`json:"last_name,omitempty"`
	Email		string	`json:"email,omitempty"`
	Phone		string	`json:"phone,omitempty"`
	CountryCode	string 	`json:"country_code,omitempty"`
	Password	string	`json:"password,omitempty"`
}

type LoginUserSchema struct {
	Email	 string		`json:"email" validate:"required"`
	Password string		`json:"password" validate:"required`
}