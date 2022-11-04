package request

type UserCreate struct{
	Name 		string `json:"name"`
	Email 		string `json:"email"`
	City 		string `json:"city"`
	Province 	string `json:"province"`
}