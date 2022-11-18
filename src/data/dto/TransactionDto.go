package dto

import "avito/src/data/entity"

type TransactionDto struct {
	Id                int     `json:"Id"`
	ServiceId         int     `json:"service"`
	Amount            float64 `json:"amount"`
	FromId            int     `json:"from"`
	ToId              int     `json:"to"`
	TransactionStatus string  `json:"status"`
}

func (t *TransactionDto) ToTransactionEntity() *entity.Transaction {

	return &entity.Transaction{
		Id:     t.Id,
		Amount: t.Amount,
		FromId: t.FromId,
		ToId:   t.ToId,
	}
}

func ToTransactionDto(t *entity.Transaction) *TransactionDto {
	return &TransactionDto{
		Id:     t.Id,
		Amount: t.Amount,
		FromId: t.FromId,
		ToId:   t.ToId,
	}
}
