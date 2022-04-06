package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func authMiddleware() gin.HandlerFunc {
	authorizationHeaderKey := "test"
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("Authorization Header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("Invalid Authorization Header Format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != "bearer" {
			err := fmt.Errorf("Unsupported Authorization Type %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

	}
}
