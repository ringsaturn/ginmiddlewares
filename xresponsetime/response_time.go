package xresponsetime

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ringsaturn/ginmiddlewares"
)

func Handler(c *gin.Context) {
	// Inject Writer
	w := ginmiddlewares.NewResponseWriter(c)
	c.Writer = w
	defer w.Done(c)

	// actual loginc
	startTime := time.Now()
	c.Next()
	duration := int(time.Since(startTime).Microseconds())
	c.Header("x-response-time", strconv.Itoa(duration))
}
