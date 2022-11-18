package entity

type User struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Account  int    `json:"account" db:"account_id"`
}
