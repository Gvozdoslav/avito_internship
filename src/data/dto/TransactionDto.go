package dto

import "avito/src/data/entity"

type TransactionDto struct {
	Id     int     `json:"Id"`
	Amount float64 `json:"amount"`
	FromId int     `json:"from"`
	ToId   int     `json:"to"`
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
