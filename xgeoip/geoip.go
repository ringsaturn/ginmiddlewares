package xgeoip

import (
	"net"

	"github.com/gin-gonic/gin"
	"github.com/oschwald/geoip2-golang"
)

// New will return a gin.HandlerFunc and set "x-geoip-city" in ctx.
//
// By default it will read `X-Forwarded-For` as first priority,
// then gin's `RemoteIP`.
func New(reader *geoip2.Reader) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ip net.IP

		forwardedIP := net.ParseIP(c.Request.Header.Get("X-Forwarded-For"))
		if forwardedIP != nil {
			ip = forwardedIP
		} else {
			ip, _ = c.RemoteIP()
		}

		city, err := reader.City(ip)
		if err != nil {
			c.Set("x-geoip-city", nil)
		} else {
			c.Set("x-geoip-city", city)
		}
		c.Next()
	}
}
