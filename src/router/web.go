package router

import (
	"crow-qin/web"
	"github.com/gin-gonic/gin"
)

func indexHtml(c *gin.Context) {
	c.Writer.WriteHeader(200)
	b, _ := web.Content.ReadFile("index.html")
	_, _ = c.Writer.Write(b)
	c.Writer.Header().Add("Accept", "text/html")
	c.Writer.Flush()
}
