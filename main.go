package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	usersRepo "github.com/pradiptarana/deall-test/repository/users"
	usersTr "github.com/pradiptarana/deall-test/transport/api/users"
	usersUC "github.com/pradiptarana/deall-test/usecase/users"
)

func main() {
	// db := dbd.NewDBConnection()
	userRepo := usersRepo.NewUserRepository()
	userUC := usersUC.NewUserUC(userRepo)
	userTr := usersTr.NewUsersTransport(userUC)
	router := gin.New()
	router.NoRoute(func(c *gin.Context) {
		sb := &strings.Builder{}
		sb.WriteString("routing err: no route, try this:\n")
		for _, v := range router.Routes() {
			sb.WriteString(fmt.Sprintf("%s %s\n", v.Method, v.Path))
		}
		c.String(http.StatusBadRequest, sb.String())
	})
	router.Group("/api")
	router.POST("/signup", userTr.SignUp)
	router.POST("/login", userTr.Login)

	router.Run("localhost:8080")
}
