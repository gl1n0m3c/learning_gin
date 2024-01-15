package models

type Person struct {
	Id        int64  `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       uint8  `json:"age"`
}
