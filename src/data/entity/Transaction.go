package entity

const (
	reserved int = 0
	accepted     = 1
	done         = 2
)

type Transaction struct {
	Id                int     `json:"id" db:"id"`
	ServiceId         int     `json:"service" db:"service_id"`
	Amount            float64 `json:"amount" db:"amount"`
	FromId            int     `json:"from" db:"from_id"`
	ToId              int     `json:"to" db:"to_id"`
	TransactionStatus string  `json:"status" db:"status"`
}
