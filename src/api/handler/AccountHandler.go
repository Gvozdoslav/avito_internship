package handler

import (
	"avito/src/api/response"
	"avito/src/data/dto"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (handler *Handler) getAccount(context *gin.Context) {

	accountId, err := getAccountIdFromContext(context)
	if err != nil {
		response.NewErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	account, err := handler.accountService.GetAccount(accountId)
	if err != nil {
		response.NewErrorResponse(context, http.StatusInternalServerError, err.Error())
	}

	context.JSON(http.StatusOK, account)
}

func (handler *Handler) getAllAccounts(context *gin.Context) {

	accounts, err := handler.accountService.GetAllAccounts()
	if err != nil {
		response.NewErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, accounts)
}

func (handler *Handler) createAccount(context *gin.Context) {

	input := new(dto.AccountDto)
	if err := context.BindJSON(input); err != nil {
		response.NewErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	responseAccountDto, err := handler.accountService.CreateAccount(input)
	if err != nil {
		response.NewErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusCreated, map[string]interface{}{
		"id":      responseAccountDto.Id,
		"balance": responseAccountDto.Balance,
	})
}

func (handler *Handler) updateAccount(context *gin.Context) {

	accountId, err := getAccountIdFromContext(context)
	if err != nil {
		response.NewErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	input := new(dto.AccountDto)
	if err := context.BindJSON(input); err != nil {
		response.NewErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	input.Id = accountId
	responseAccountDto, err := handler.accountService.UpdateAccount(accountId, input)
	if err != nil {
		response.NewErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"id":      responseAccountDto.Id,
		"balance": responseAccountDto.Balance,
	})
}

func (handler *Handler) deleteAccount(context *gin.Context) {

	accountId, err := getAccountIdFromContext(context)
	if err != nil {
		response.NewErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	isOk, err := handler.accountService.Delete(accountId)
	if !isOk || err != nil {
		response.NewErrorResponse(context, http.StatusInternalServerError, err.Error())
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"Message": "account has been deleted successfully",
	})
}

func getAccountIdFromContext(c *gin.Context) (int, error) {

	id := c.Param(accountIdContext)
	if id == "" {
		return 0, errors.New("entity id not found")
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return 0, errors.New("entity id is of invalid type")
	}

	return idInt, nil
}
