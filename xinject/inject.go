package xinject

import (
	"github.com/gin-gonic/gin"
	"github.com/ringsaturn/ginmiddlewares/internal/writer"
)

func Handler(c *gin.Context) {
	// Inject Writer
	w := writer.NewResponseWriter(c)
	c.Writer = w
	defer w.Done(c)
	c.Next()
}
