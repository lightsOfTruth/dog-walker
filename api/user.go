package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/lightsOfTruth/dog-walker/db/sqlc"
)


type CreateUserRequestParams struct {
		FullName  string        `json:"full_name" binding:"required"`
		Contact   string        `json:"contact" binding:"required"`
		Dog       int			`json:"dog"`
		Address   string        `json:"address"`
		City      string        `json:"city"`
		PostCode  string        `json:"post_code"`
		Longitude string        `json:"longitude" binding:"required"`
		Latitude  string        `json:"latitude" binding:"required"`


}

func (server *Server) createUser(ctx *gin.Context) {
	var requestType CreateUserRequestParams

	if err := ctx.ShouldBindJSON(&requestType); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errorResponse(err)})
		return
	}

	createUserDBArgs := db.CreateUserParams {
		ID: uuid.New(),
		FullName: requestType.FullName,
		Contact: requestType.Contact,
		Dog: sql.NullInt32{Int32: int32(requestType.Dog),Valid: true},
		Address: requestType.Address,
		City: requestType.City,
		PostCode: requestType.PostCode,
		Longitude: requestType.Longitude,
		Latitude: requestType.Latitude,
	}

	user, err := server.store.CreateUser(ctx,createUserDBArgs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errorResponse(err)})
		return
	}

	ctx.JSON(http.StatusOK, user)


}