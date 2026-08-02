package helpers

import (
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
	"github.com/zgierz/klimson/backend/khttp"
)

func keyFunc(ctx *gin.Context) string {
	return ctx.ClientIP()
}

func errorHandler(ctx *gin.Context, info ratelimit.Info) {
	khttp.TooManyRequestsResponse(ctx, time.Until(info.ResetTime).String())
}

var publicRateLimitRuleConfig = ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
	Rate:  time.Minute,
	Limit: 10,
})
var adminRateLimitRuleConfig = ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
	Rate:  time.Second,
	Limit: 20,
})

func PublicRateLimiter() gin.HandlerFunc {
	return ratelimit.RateLimiter(publicRateLimitRuleConfig, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      keyFunc,
	})
}

func AdminRateLimiter() gin.HandlerFunc {
	return ratelimit.RateLimiter(adminRateLimitRuleConfig, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      keyFunc,
	})
}
