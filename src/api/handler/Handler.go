package handler

import (
	"avito/src/domain/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.UserServiceInterface
}

func NewHandler(service *service.UserServiceInterface) *Handler {
	return &Handler{
		service: service,
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
