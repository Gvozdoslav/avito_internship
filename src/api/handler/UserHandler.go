package handler

import (
	"avito/src/api/response"
	"avito/src/data/dto"
	_ "avito/src/domain/service"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (handler *Handler) getUser(context *gin.Context) {

	userId, err := getUserIdFromContext(context)
	if err != nil {
		response.NewErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := handler.userService.GetUser(userId)
	if err != nil {
		response.NewErrorResponse(context, http.StatusInternalServerError, err.Error())
	}

	context.JSON(http.StatusOK, user)
}

func (handler *Handler) getAllUsers(context *gin.Context) {

	users, err := handler.userService.GetAllUsers()
	if err != nil {
		response.NewErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, users)
}

func (handler *Handler) createUser(context *gin.Context) {

	input := new(dto.UserDto)
	if err := context.BindJSON(input); err != nil {
		response.NewErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	responseUserDto, err := handler.userService.CreateUser(input)
	if err != nil {
		response.NewErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusCreated, map[string]interface{}{
		"id":       responseUserDto.Id,
		"name":     responseUserDto.Name,
		"username": responseUserDto.Username,
		"account":  responseUserDto.Account,
	})
}

func (handler *Handler) updateUser(context *gin.Context) {

	userId, err := getUserIdFromContext(context)
	if err != nil {
		response.NewErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	input := new(dto.UserDto)
	if err := context.BindJSON(input); err != nil {
		response.NewErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	input.Id = userId
	responseUserDto, err := handler.userService.UpdateUser(userId, input)
	if err != nil {
		response.NewErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"id":       responseUserDto.Id,
		"name":     responseUserDto.Name,
		"username": responseUserDto.Username,
		"account":  responseUserDto.Account,
	})
}

func (handler *Handler) bookService(context *gin.Context) {

	input := new(dto.TransactionDto)
	if err := context.BindJSON(input); err != nil {
		response.NewErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	responseTransactionDto, err := handler.transactionService.BookForService(input)
	if err != nil {
		response.NewErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	err = handler.accountService.BookForService(input.FromId, input.ToId, input.Amount)
	if err != nil {
		response.NewErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusCreated, map[string]interface{}{
		"Message": "Transaction with id " + strconv.Itoa(responseTransactionDto.Id) + " has been booked successfully",
	})
}

func (handler *Handler) payForService(context *gin.Context) {

	input := new(dto.TransactionDto)
	if err := context.BindJSON(input); err != nil {
		response.NewErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	responseTransactionDto, err := handler.transactionService.PayForService(input)
	if err != nil {
		response.NewErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	err = handler.accountService.PayForService(input.FromId, input.ToId, input.Amount)
	if err != nil {
		response.NewErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusCreated, map[string]interface{}{
		"Message": "Transaction with id " + strconv.Itoa(responseTransactionDto.Id) + " has been payed successfully",
	})
}

func (handler *Handler) deleteUser(context *gin.Context) {

	userId, err := getUserIdFromContext(context)
	if err != nil {
		response.NewErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	isOk, err := handler.userService.DeleteUser(userId)
	if !isOk || err != nil {
		response.NewErrorResponse(context, http.StatusInternalServerError, err.Error())
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"Message": "user has been deleted successfully",
	})
}

func getUserIdFromContext(c *gin.Context) (int, error) {

	id := c.Param(userIdContext)
	if id == "" {
		return 0, errors.New("entity id not found")
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return 0, errors.New("entity id is of invalid type")
	}

	return idInt, nil
}
