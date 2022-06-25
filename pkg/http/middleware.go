package http

import (
	"context"
	"net/http"

	"github.com/arthureichelberger/trailrcore/pkg/jwt"
	"github.com/gin-gonic/gin"
)

type MiddlewareContextKey string

const ClaimsMiddlewareKey MiddlewareContextKey = "claims"

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")
		if len(auth) < 7 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		authToken := auth[7:]

		claims, err := jwt.Decode(authToken)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), ClaimsMiddlewareKey, claims))
		ctx.Next()
	}
}
