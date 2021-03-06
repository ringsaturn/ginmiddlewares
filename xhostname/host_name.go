package xhostname

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Handler(c *gin.Context) {
	c.Next()
	name, err := os.Hostname()
	if err != nil {
		name = "unknown"
	}
	c.Header("x-host-name", name)
}
