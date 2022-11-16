package main

import (
	handler2 "avito/src/api/handler"
	"avito/src/data/configs"
	"avito/src/data/repository"
	server2 "avito/src/data/server"
	"avito/src/domain/service"
	"log"
)

func main() {

	db, err := configs.NewPostgresDb(nil)
	if err != nil {
		log.Fatalf("error occured while connecting to db: %s", err.Error())
	}

	accountRepository := repository.NewAccountRepository(db)
	transactionRepository := repository.NewTransactionRepository(db)
	userRepository := repository.NewUserRepository(db)

	userService := service.UserServiceInterface(*service.NewUserService(userRepository, accountRepository, transactionRepository))
	handler := handler2.NewHandler(&userService)
	srv := new(server2.Server)
	if err := srv.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running server: %s", err.Error())
	}
}
