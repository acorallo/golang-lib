package gin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Headers(h func(*gin.Context)) func(*gin.Context) {
	return func(c *gin.Context) {
		fmt.Println("hello from github.com/gin-gonic/gin Headers")
		h(c)
	}
}
