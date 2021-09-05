package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oschwald/geoip2-golang"
	"github.com/ringsaturn/ginmiddlewares/xgeoip"
)

func GeoIP(c *gin.Context) {
	val, ok := c.Get("x-geoip-city")
	if ok {
		c.JSON(200, val)
	}
}

func GeoIPLocation(c *gin.Context) {
	val, ok := c.Get("x-geoip-city")
	if !ok {
		c.String(404, "not found")
		return
	}

	city, ok := val.(*geoip2.City)
	if !ok {
		c.String(500, "Internal error")
		return
	}
	c.JSON(200, city.Location)
}

func main() {

	db, err := geoip2.Open("GeoIP2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := gin.New()

	router.Use(xgeoip.New(db))

	router.GET("/geoip/info", GeoIP)
	router.GET("/geoip/location", GeoIPLocation)

	server := http.Server{
		Addr:    "localhost:8999",
		Handler: router,
	}
	_ = server.ListenAndServe()
}
