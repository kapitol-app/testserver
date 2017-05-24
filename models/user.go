package models

type User struct {
	UserId int `json:"userId"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastname"`
	Dob string `json:"dob"`
	Age int `json:"age"`
	IsActive bool `json:"isActive"`
}
