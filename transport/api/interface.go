package api

import "github.com/gin-gonic/gin"

type UsersTransport interface {
	Disburse(c *gin.Context)
}
