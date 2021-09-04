package xinject

import (
	"github.com/gin-gonic/gin"
	"github.com/ringsaturn/ginmiddlewares"
)

func Handler(c *gin.Context) {
	// Inject Writer
	w := ginmiddlewares.NewResponseWriter(c)
	c.Writer = w
	defer w.Done(c)
	c.Next()
}
