package entity

type Transaction struct {
	id     int     `json:"id"`
	amount float64 `json:"amount"`
	from   Account `json:"from"`
	to     Account `json:"to"`
}
