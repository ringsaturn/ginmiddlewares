package xresponsetime

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Handler(c *gin.Context) {
	startTime := time.Now()
	c.Next()
	duration := int(time.Since(startTime).Milliseconds())
	c.Header("x-response-time", strconv.Itoa(duration))
}
