package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestAuthMiddleware(t *testing.T) {

	server := newTestServer(t, nil)

	authPath := "/auth"

	server.router.GET(authPath, authMiddleware(server.tokenCreator), func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{})
	})

	recorder := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, authPath, nil)
	require.NoError(t, err)

	// add bearer [TOKEN] to header

	token, err := server.tokenCreator.CreateToken("testauth", time.Minute)
	require.NoError(t, err)

	authHeader := fmt.Sprintf("%s %s", authBearerType, token)
	request.Header.Set(authHeaderKey, authHeader)
	server.router.ServeHTTP(recorder, request)

	require.Equal(t, http.StatusOK, recorder.Code)

}
