package middleware

import (
	"github.com/gin-gonic/gin"
	"io"
)

func Recovery(f func(c *gin.Context, err interface{})) gin.HandlerFunc {
	return RecoveryWithWriter(f, gin.DefaultErrorWriter)
}

func RecoveryWithWriter(f func(c *gin.Context, err interface{}), out io.Writer) gin.HandlerFunc {
	return func(g *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				f(g, err)
			}
		}()
		g.Next()
	}
}
