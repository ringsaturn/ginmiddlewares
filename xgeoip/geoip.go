package xgeoip

import (
	"net"

	"github.com/gin-gonic/gin"
	"github.com/oschwald/geoip2-golang"
)

func New(reader *geoip2.Reader) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ip net.IP

		nginxIPAddr := c.Request.Header.Get("X-Forwarded-For")
		if nginxIPAddr != "" {
			ip = net.ParseIP(nginxIPAddr)
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
