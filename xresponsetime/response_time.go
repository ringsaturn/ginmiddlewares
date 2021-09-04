package xresponsetime

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Handler(c *gin.Context) {
	// actual loginc
	startTime := time.Now()
	c.Next()
	duration := int(time.Since(startTime).Microseconds())
	c.Header("x-response-time", strconv.Itoa(duration))
}
