package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/lightsOfTruth/dog-walker/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	server.router = gin.Default()
	server.router.POST("/createuser", server.createUser)

	return server
}

// router is private to this api package only because it is lowercased
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(e error) gin.H {
	return gin.H{"errors": e.Error()}
}
