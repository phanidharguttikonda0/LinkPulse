package middlewares

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	_ "golang.org/x/time/rate"
	"net/http"
	"sync"
	_ "sync"
	"time"
	_ "time"
)

var visitors = make(map[string]*rate.Limiter)
var mutex sync.Mutex

func getLimiter(ip string) *rate.Limiter {
	mutex.Lock()
	defer mutex.Unlock()
	limiter, exists := visitors[ip]
	if !exists {
		/*
			it allows only 1 request for 3 seconds and 20 requests a total for 1 minute
		*/
		limiter = rate.NewLimiter(rate.Every(time.Minute/20), 20)
		visitors[ip] = limiter
	}
	return limiter
}

func RateLimiterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		limiter := getLimiter(ip)
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests",
			})
			return
		}
		c.Next()
	}
}
