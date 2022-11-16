package entity

type User struct {
	id      int     `json:"id"`
	name    string  `json:"name"`
	account Account `json:"account"`
}
