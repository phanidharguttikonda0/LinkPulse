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
			it allows only 20 requests per minute, if more than that then it stop working still next minute
			Means once we sent 20 requests per minute then we need to wait for 3 seconds to empty one bucket ,
			in total 1 minute need to be waited to send again 20 request continuously, else for every 3 seconds
			we can send only one request if have doesn't wait
		*/
		limiter = rate.NewLimiter(rate.Every(time.Minute/30), 30)
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
