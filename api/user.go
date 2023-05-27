package api

import (
	"net/http"
	"time"

	db "github.com/PuneetBirdi/golang-bank/db/sqlc"
	"github.com/PuneetBirdi/golang-bank/util"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createUserRequest struct {
	FullName    string `json:"full_name" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Email string `json:"email" binding:"required,email"`
}

type createUserResponse struct {
	ID                int64     `json:"id"`
	FullName          string    `json:"full_name"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	Email             string    `json:"email"`
	CreatedAt         time.Time `json:"created_at"`
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

	response := createUserResponse{
		ID: user.ID,
		FullName: user.FullName,
		PasswordChangedAt: user.PasswordChangedAt,
		Email: user.Email,
		CreatedAt: user.CreatedAt,
	}


	ctx.JSON(http.StatusOK, response)
}


