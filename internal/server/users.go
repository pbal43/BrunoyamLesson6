package server

import (
	userErrors "BrunoyamLesson6/internal/domain/users/errors"
	"BrunoyamLesson6/internal/domain/users/models"
	"BrunoyamLesson6/internal/usecase/userusecase"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (srv *RentApi) Register(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	usecase := userusecase.NewUserUsecase(srv.db)
	if err := usecase.SaveUser(user); err != nil {
		if errors.Is(err, userErrors.ErrorUserIsAlreadyExist) {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (srv *RentApi) Login(ctx *gin.Context) {
	var usReq models.UserRequest
	if err := ctx.ShouldBindJSON(&usReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	usecase := userusecase.NewUserUsecase(srv.db)
	user, err := usecase.LoginUser(usReq)
	if err != nil {
		if errors.Is(err, userErrors.ErrorInvalidPassword) || errors.Is(err, userErrors.ErrorUserNotExist) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
