package api

import (
	"fmt"
	"net/http"
	"strings"

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
	router := gin.New()
	// Handling routing errors
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

	router.ServeHTTP(w, r)
}
