package handler

import (
	"avito/src/domain/service"
	"github.com/gin-gonic/gin"
)

const (
	userIdContext                = "user_id"
	accountIdContext             = "account_id"
	senderTransactionIdContext   = "sender_transaction_id"
	recieverTransactionIdContext = "reciever_transaction_id"
)

type Handler struct {
	userService        *service.UserService
	accountService     *service.AccountService
	transactionService *service.TransactionService
}

func NewHandler(userService *service.UserService,
	accountService *service.AccountService,
	transactionService *service.TransactionService) *Handler {
	return &Handler{
		userService:        userService,
		accountService:     accountService,
		transactionService: transactionService,
	}
}

func (handler *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		user := api.Group("/user")
		{
			user.GET("/all", handler.getAllUsers)
			user.GET("/:user_id", handler.getUser)
			user.POST("/", handler.createUser)
			user.PUT("/:user_id", handler.updateUser)
			user.PUT("/book/:user_id", handler.bookService)
			user.PUT("/pay/:user_id", handler.payForService)
			user.DELETE("/:user_id", handler.deleteUser)

			account := user.Group(":user_id/account")
			{
				account.GET("/all", handler.getAllAccounts)
				account.GET("/:account_id", handler.getAccount)
				account.POST("/", handler.createAccount)
				account.PUT("/:account_id", handler.updateAccount)
				account.DELETE("/:account_id", handler.deleteAccount)
			}
		}
	}

	return router
}
