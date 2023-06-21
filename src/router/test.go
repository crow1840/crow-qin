package router

import (
	"github.com/gin-gonic/gin"
)

func LiuBei(c *gin.Context) {

	SetOK(c, "hello world")
}
