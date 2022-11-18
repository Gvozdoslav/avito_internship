package entity

type Account struct {
	Id      int     `json:"Id" db:"id"`
	Balance float64 `json:"Balance" db:"balance"`
}
