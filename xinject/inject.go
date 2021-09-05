package xinject

import (
	"github.com/gin-gonic/gin"
	"github.com/ringsaturn/ginmiddlewares/internal/writer"
)

// Handler is the ResponseWriter inject func, this will replace gin's
// default writer, so we can change header after `c.Next()`
func Handler(c *gin.Context) {
	w := writer.NewResponseWriter(c)
	c.Writer = w
	defer w.Done(c)
	c.Next()
}
