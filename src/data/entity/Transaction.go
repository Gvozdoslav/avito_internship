package entity

const (
	reserved string = "reserved"
	accepted        = "accepted"
	done            = "done"
)

type Transaction struct {
	Id                int     `json:"Id" db:"id"`
	Amount            float64 `json:"amount" db:"amount"`
	FromId            int     `json:"from" db:"from_id"`
	ToId              int     `json:"to" db:"to_id"`
	TransactionStatus string  `json:"status" db:"status"`
}
