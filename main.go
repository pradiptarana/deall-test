package main

import (
	"github.com/gin-gonic/gin"

	usersRepo "github.com/pradiptarana/deall-test/repository/users"
	usersTr "github.com/pradiptarana/deall-test/transport/api/users"
	usersUC "github.com/pradiptarana/deall-test/usecase/users"
)

func main() {
	userRepo := usersRepo.NewUserRepository()
	userUC := usersUC.NewUserUC(userRepo)
	userTr := usersTr.NewUsersTransport(userUC)
	router := gin.New()
	router.POST("/signup", userTr.SignUp)
	router.POST("/login", userTr.Login)

	router.Run("localhost:8080")
}
