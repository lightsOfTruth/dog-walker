package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/lightsOfTruth/dog-walker/db/sqlc"
	"github.com/lightsOfTruth/dog-walker/helpers"
)

type CreateUserRequestParams struct {
	FullName  string                `json:"full_name" binding:"required"`
	Email     string                `json:"email"`
	Password  string                `json:"password" binding:"required"`
	Contact   string                `json:"contact" binding:"required"`
	Dog       helpers.JsonNullInt32 `json:"dog"`
	Address   string                `json:"address"`
	City      string                `json:"city"`
	PostCode  string                `json:"post_code"`
	Longitude string                `json:"longitude" binding:"required"`
	Latitude  string                `json:"latitude" binding:"required"`
}

// leave out the password or any other sensitive data that would otherwise be included in createUserRequestParams
type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	FullName  string    `json:"full_name" binding:"required"`
	Email     string    `json:"email"`
	Contact   string    `json:"contact" binding:"required"`
	Dog       int32     `json:"dog"`
	Address   string    `json:"address"`
	City      string    `json:"city"`
	PostCode  string    `json:"post_code"`
	Longitude string    `json:"longitude" binding:"required"`
	Latitude  string    `json:"latitude" binding:"required"`
}

func FilteredUserResponse(user db.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		FullName:  user.FullName,
		Email:     user.Email,
		Contact:   user.Contact,
		Dog:       user.Dog.Int32,
		Address:   user.Address,
		City:      user.City,
		PostCode:  user.PostCode,
		Longitude: user.Longitude,
		Latitude:  user.Latitude,
	}

}

func (server *Server) createUser(ctx *gin.Context) {
	var requestType CreateUserRequestParams

	if err := ctx.ShouldBindJSON(&requestType); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errorResponse(err)})
		return
	}

	hashedPassword, err := helpers.HashPassword(requestType.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errorResponse(err)})
		return
	}

	createUserDBArgs := db.CreateUserParams{
		ID:        uuid.New(),
		FullName:  requestType.FullName,
		Email:     requestType.Email,
		Password:  hashedPassword,
		Contact:   requestType.Contact,
		Dog:       requestType.Dog.NullInt32,
		Address:   requestType.Address,
		City:      requestType.City,
		PostCode:  requestType.PostCode,
		Longitude: requestType.Longitude,
		Latitude:  requestType.Latitude,
	}

	user, err := server.store.CreateUser(ctx, createUserDBArgs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errorResponse(err)})
		return
	}

	ctx.JSON(http.StatusOK, FilteredUserResponse(user))

}

type LoginUserRequestParams struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        UserResponse `json:"user"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var requestType LoginUserRequestParams

	if err := ctx.ShouldBindJSON(&requestType); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errorResponse(err)})
		return
	}

	user, err := server.store.GetUserByEmail(ctx, requestType.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = helpers.CheckPassword(requestType.Password, user.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, err := server.tokenCreator.CreateToken(user.Email, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, LoginUserResponse{
		AccessToken: accessToken,
		User:        FilteredUserResponse(user),
	})

}
