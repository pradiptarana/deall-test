package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	usersRepo "deal-test/repository/users"
	usersTr "deal-test/transport/api/users"
	usersUC "deal-test/usecase/users"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// db := dbd.NewDBConnection()
	userRepo := usersRepo.NewUserRepository()
	userUC := usersUC.NewUserUC(userRepo)
	userTr := usersTr.NewUsersTransport(userUC)
	router := gin.Default()
	router.POST("/signup", userTr.SignUp)
	router.POST("/login", userTr.Login)

	router.ServeHTTP(w, r)
}
