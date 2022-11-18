package handler

import (
	"avito/src/domain/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	userService    *service.UserService
	accountService *service.AccountService
}

func NewHandler(userService *service.UserService, accountService *service.AccountService) *Handler {
	return &Handler{
		userService:    userService,
		accountService: accountService,
	}
}

func (handler *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		user := api.Group("/user")
		{
			user.GET("/all", handler.getAllUsers)
			user.GET("/:id", handler.getUser)
			user.POST("/", handler.createUser)
			user.PUT("/:id", handler.updateUser)
			user.PUT("/pay/:id", handler.payForService)
			user.DELETE("/:id", handler.deleteUser)

			account := user.Group(":id/account")
			{
				account.GET("/all", handler.getAllAccounts)
				account.GET("/:id", handler.getAccount)
				account.POST("/", handler.createAccount)
				account.PUT("/:id", handler.updateAccount)
			}
		}
	}

	return router
}
