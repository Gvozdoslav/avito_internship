package entity

type User struct {
	id       int     `json:"id"`
	name     string  `json:"name"`
	username string  `json:"username"`
	password string  `json:"password"`
	account  Account `json:"account"`
}
