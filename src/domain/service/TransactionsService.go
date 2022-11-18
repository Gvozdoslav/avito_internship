package service

import (
	"avito/src/data/dto"
	"avito/src/data/repository"
)

type TransactionService struct {
	transactionRepository *repository.TransactionRepository
}

func NewTransactionService(transactionRepository *repository.TransactionRepository) *TransactionService {
	return &TransactionService{
		transactionRepository: transactionRepository,
	}
}

func (t *TransactionService) BookForService(transactionDto *dto.TransactionDto) (*dto.TransactionDto, error) {

	transactionDto.TransactionStatus = "reserved"
	transaction, err := t.transactionRepository.Create(transactionDto.ToTransactionEntity())
	if err != nil {
		return nil, err
	}

	transactionResponseDto := dto.ToTransactionDto(transaction)
	return transactionResponseDto, nil
}

func (t *TransactionService) PayForService(transactionDto *dto.TransactionDto) (*dto.TransactionDto, error) {

	transactionDto.TransactionStatus = "accepted"
	transaction, err := t.transactionRepository.Update(transactionDto.Id, transactionDto.ToTransactionEntity())
	if err != nil {
		return nil, err
	}

	transactionDto = dto.ToTransactionDto(transaction)
	transactionDto.TransactionStatus = "done"

	transaction, err = t.transactionRepository.Update(transactionDto.Id, transactionDto.ToTransactionEntity())
	if err != nil {
		return nil, err
	}

	transactionResponseDto := dto.ToTransactionDto(transaction)
	return transactionResponseDto, nil
}
