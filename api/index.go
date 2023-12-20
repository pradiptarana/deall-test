package main

import (
	"github.com/gin-gonic/gin"

	usersRepo "deal-test/repository/users"
	usersTr "deal-test/transport/api/users"
	usersUC "deal-test/usecase/users"
)

func Handler(c *gin.Context) {
	// db := dbd.NewDBConnection()
	userRepo := usersRepo.NewUserRepository()
	userUC := usersUC.NewUserUC(userRepo)
	userTr := usersTr.NewUsersTransport(userUC)
	router := gin.Default()
	router.POST("/signup", userTr.SignUp)
	router.POST("/login", userTr.Login)

	router.Run("localhost:8080")
}
