package server

import (
	userErrors "BrunoyamLesson6/internal/domain/users/errors"
	"BrunoyamLesson6/internal/domain/users/models"
	"BrunoyamLesson6/internal/usecase/userusecase"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (srv *RentApi) register(ctx *gin.Context) {
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

func (srv *RentApi) login(ctx *gin.Context) {
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

	access, err := srv.jwtSigner.NewAccessToken(user.Uuid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	refresh, err := srv.jwtSigner.NewRefreshToken(user.Uuid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.SetCookie("refresh_token", refresh, 3600*24*7, "/", "127.0.0.1:8080", false, true)
	ctx.JSON(http.StatusOK, gin.H{"user": user, "access": access})
}

func (srv *RentApi) profile(ctx *gin.Context) {

	userID, exist := ctx.Get("userID")
	if !exist {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User ID not found in request context"})
		return
	}
	user, err := userusecase.UserStorage(srv.db).GetUserByID(userID.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
