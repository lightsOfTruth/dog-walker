package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lightsOfTruth/dog-walker/authentication"
	db "github.com/lightsOfTruth/dog-walker/db/sqlc"
	"github.com/lightsOfTruth/dog-walker/helpers"
)

type Server struct {
	store        db.Store
	config       helpers.Config
	tokenCreator authentication.TokenCreator
	router       *gin.Engine
}

func NewServer(config helpers.Config, store db.Store) (*Server, error) {
	tokenCreator, err := authentication.NewJWTCreator(config.SecretKey)
	if err != nil {
		return nil, fmt.Errorf("token creation failed %w", err)
	}

	server := &Server{
		store:        store,
		config:       config,
		tokenCreator: tokenCreator}

	server.initRouter()

	return server, nil
}

func (server *Server) initRouter() {
	server.router = gin.Default()
	server.router.POST("/createuser", server.createUser)
	server.router.POST("/user/login", server.loginUser)
	server.router.Group("/").Use(authMiddleware(server.tokenCreator))

	// all routes that should use auth middleware need authRoutes.POST instead of server.router
}

// router is private to this api package only because it is lowercased
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(e error) gin.H {
	return gin.H{"errors": e.Error()}
}
