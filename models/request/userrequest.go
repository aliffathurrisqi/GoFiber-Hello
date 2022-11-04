package request

type UserCreate struct{
	Name 		string `json:"name" validate:"required,min=2"`
	Email 		string `json:"email" validate:"required,email"`
	City 		string `json:"city"`
	Province 	string `json:"province"`
}