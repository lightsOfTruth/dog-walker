package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lightsOfTruth/dog-walker/authentication"
)

const (
	authHeaderKey  = "authorization"
	authBearerType = "bearer"
	authPayloadKey = "auth_payload"
)

func authMiddleware(tokenMaker authentication.TokenCreator) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader(authHeaderKey)
		if len(authHeader) == 0 {
			err := errors.New("authorisation header not found")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		fields := strings.Fields(authHeader)
		if len(fields) < 2 {
			err := errors.New("invalid auth header")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return

		}

		authType := fields[0]
		if authType != authBearerType {
			err := fmt.Errorf("authorisation type is not supported %s", authType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		accessToken := fields[1]

		tokenPayload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		/**
		Setting the tokenPayload inside the context will make it accessible via
		cts.MustGet(authPayloadKey).(*token.Payload)
		This allows to authorise api calls based on payload credentials
		*/
		ctx.Set(authPayloadKey, tokenPayload)
		ctx.Next()
	}
}
