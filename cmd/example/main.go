package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ringsaturn/ginmiddlewares/xresponsetime"
	"github.com/ringsaturn/ginmiddlewares/xservername"
)

func Ping(c *gin.Context) {
	c.String(200, "pong")
}

func main() {
	router := gin.New()

	router.Use(xresponsetime.Handler)
	router.Use(xservername.Handler)

	router.GET("/ping", Ping)

	server := http.Server{
		Addr:    "localhost:8999",
		Handler: router,
	}
	server.ListenAndServe()
}
