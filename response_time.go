package ginmiddlewares

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func ResponseTimeHeader(c *gin.Context) {
	startTime := time.Now()
	c.Next()
	duration := int(time.Since(startTime).Microseconds())
	c.Header("x-response-time", strconv.Itoa(duration))
}
