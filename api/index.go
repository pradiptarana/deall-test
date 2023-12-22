package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	usersRepo "github.com/pradiptarana/deall-test/repository/users"
	usersTr "github.com/pradiptarana/deall-test/transport/api/users"
	usersUC "github.com/pradiptarana/deall-test/usecase/users"
)

var (
	app *gin.Engine
)

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

// init gin app
func init() {
	userRepo := usersRepo.NewUserRepository()
	userUC := usersUC.NewUserUC(userRepo)
	userTr := usersTr.NewUsersTransport(userUC)
	app = gin.New()

	// Handling routing errors
	app.NoRoute(func(c *gin.Context) {
		sb := &strings.Builder{}
		sb.WriteString("routing err: no route, try this:\n")
		for _, v := range app.Routes() {
			sb.WriteString(fmt.Sprintf("%s %s\n", v.Method, v.Path))
		}
		c.String(http.StatusBadRequest, sb.String())
	})

	r := app.Group("/")

	// register route

	r.POST("/api/signup", userTr.SignUp)
	r.POST("/api/login", userTr.Login)
	r.POST("/ai/ping", Ping)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// userRepo := usersRepo.NewUserRepository()
	// userUC := usersUC.NewUserUC(userRepo)
	// userTr := usersTr.NewUsersTransport(userUC)
	// router := gin.New()
	// // Handling routing errors
	// router.NoRoute(func(c *gin.Context) {
	// 	sb := &strings.Builder{}
	// 	sb.WriteString("routing err: no route, try this:\n")
	// 	for _, v := range router.Routes() {
	// 		sb.WriteString(fmt.Sprintf("%s %s\n", v.Method, v.Path))
	// 	}
	// 	c.String(http.StatusBadRequest, sb.String())
	// })
	// router.POST("/api/signup", userTr.SignUp)
	// router.POST("/api/login", userTr.Login)
	// router.POST("/ai/ping", Ping)

	app.ServeHTTP(w, r)
}
