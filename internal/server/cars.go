package server

import (
	carErrors "BrunoyamLesson6/internal/domain/cars/errors"
	carDomain "BrunoyamLesson6/internal/domain/cars/models"
	"BrunoyamLesson6/internal/usecase/carusecase"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (srv *RentApi) getRent(ctx *gin.Context) {
	carID := ctx.Param("id")
	usecase := carusecase.NewCarUsecase(srv.db)
	car, err := usecase.GetCarByID(carID)
	if err != nil {
		if errors.Is(err, carErrors.ErrCarsNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, carErrors.ErrCarNotAvailable) {
			ctx.JSON(http.StatusConflict, gin.H{})
			return
		}
	}

	//userID, exist := ctx.Param("userID")
	//if !exist {
	//	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in request context"})
	//	return
	//}

	ctx.JSON(http.StatusOK, gin.H{"msg": "rent successful", "car": car})
}

func (srv *RentApi) getAllCars(ctx *gin.Context) {
	usecase := carusecase.NewCarUsecase(srv.db)
	cars, err := usecase.GetAllCars()
	if err != nil {
		if errors.Is(err, carErrors.ErrCarsNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": carErrors.ErrCarsNotFound})
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"cars": cars})
}

func (srv *RentApi) getAvailableCars(ctx *gin.Context) {
	usecase := carusecase.NewCarUsecase(srv.db)
	cars, err := usecase.GetAvailableCars()
	if err != nil {
		if errors.Is(err, carErrors.ErrCarsNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": carErrors.ErrCarsNotFound})
		}
		if errors.Is(err, carErrors.ErrAvailableCarsNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": carErrors.ErrAvailableCarsNotFound})
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "available cars", "cars": cars})
}

func (srv *RentApi) addCar(ctx *gin.Context) {
	var car carDomain.Car
	if err := ctx.ShouldBindJSON(&car); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usecase := carusecase.NewCarUsecase(srv.db)
	if err := usecase.AddCar(car); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusCreated, gin.H{"msg": "car added", "car": car})
}
