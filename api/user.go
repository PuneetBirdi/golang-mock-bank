package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	db "github.com/PuneetBirdi/golang-bank/db/sqlc"
	"github.com/PuneetBirdi/golang-bank/util"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createUserRequest struct {
	FullName    string `json:"full_name" binding:"required"`
	Password    string `json:"password" binding:"required,min=6"`
	Email       string `json:"email" binding:"required,email"`
}

type userResponse struct {
	ID                int64     `json:"id"`
	FullName          string    `json:"full_name"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	Email             string    `json:"email"`
	CreatedAt         time.Time `json:"created_at"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse {
		ID: user.ID,
		FullName: user.FullName,
		PasswordChangedAt: user.PasswordChangedAt,
		Email: user.Email,
		CreatedAt: user.CreatedAt,
	}
}
func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	arg := db.CreateUserParams{
		FullName:    req.FullName,
		HashedPassword: hashedPassword,
		Email:  req.Email,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := newUserResponse(user)



	ctx.JSON(http.StatusOK, response)
}

type loginUserRequest struct {
	Password string `json:"password" binding:"required,min=6"`
	Email string `json:"email" binding:"required,email"`
}

type loginUserResponse struct {
	AccessToken string `json:"access_token"`
	User userResponse		`json:"user"`
}


func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	args := db.GetUserParams{
		Email: req.Email,
		ID: -1,
	}

	user, err := server.store.GetUser(ctx, args)
	if err != nil {
		fmt.Println(err)
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, _, err := server.tokenMaker.CreateToken(user.ID, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	response := loginUserResponse{
		AccessToken: accessToken,
		User: newUserResponse(user),
	}
	ctx.JSON(http.StatusOK, response)
}



