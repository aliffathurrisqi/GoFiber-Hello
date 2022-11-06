package request

import (
	"time"
)

type UserCreate struct{
	Name 		string `json:"name" validate:"required,min=2"`
	Email 		string `json:"email" validate:"required,email"`
	City 		string `json:"city"`
	Province 	string `json:"province"`
}

type UserEdit struct{
	Name 		string `json:"name" validate:"required,min=2"`
	City 		string `json:"city"`
	Province 	string `json:"province"`
	Updated_at 	time.Time `json:"updated_at,omitempty" db:"updated_at"`
}