package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"query/utils"
)

func Tls() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tls := secure.New(secure.Options{
			SSLHost: utils.HttpHost+utils.HttpPort,
			SSLRedirect: true,
		})
		err := tls.Process(ctx.Writer,ctx.Request)
		if err != nil {
			return
		}
		ctx.Next()
	}
}