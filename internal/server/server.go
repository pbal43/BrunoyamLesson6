package server

import (
	"BrunoyamLesson6/internal"
	carDomain "BrunoyamLesson6/internal/domain/cars/models"
	userDomain "BrunoyamLesson6/internal/domain/users/models"
	"BrunoyamLesson6/internal/server/auth"
	"BrunoyamLesson6/internal/server/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type UserStorage interface {
	SaveUser(user userDomain.User) error
	GetUser(userReq userDomain.UserRequest) (userDomain.User, error)
	GetUserByID(uid string) (userDomain.User, error)
}

type CarStorage interface {
	GetAllCars() ([]carDomain.Car, error)
	GetCarByID(id string) (carDomain.Car, error)
	GetAvailableCars() ([]carDomain.Car, error)
	AddCar(car carDomain.Car) error
	UpdateAvailable(cid string) error
}

type Storage interface {
	UserStorage
	CarStorage
}

type RentApi struct {
	srv       *http.Server
	db        Storage
	jwtSigner auth.HS256Signer
}

func NewServer(cfg internal.Config, db Storage) *RentApi {
	signer := auth.HS256Signer{
		Secret:     []byte("ultraSecretKey123"),
		Issuer:     "rentService",
		Audience:   "rentClient",
		AccessTTL:  15 * time.Minute,
		RefreshTTL: 24 * 7 * time.Hour,
	}
	HttpSrv := http.Server{
		Addr: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
	}

	api := RentApi{srv: &HttpSrv, db: db, jwtSigner: signer}

	api.configRouter()

	return &api
}

func (api *RentApi) Run() error {
	return api.srv.ListenAndServe()
}

func (api *RentApi) ShutDown() error {
	return nil
}

func (api *RentApi) configRouter() {
	router := gin.Default()

	users := router.Group("/users")
	users.POST("/login", api.login)
	users.POST("/register", api.register)
	users.GET("/profile", middleware.AuthMiddleware(api.jwtSigner), api.profile)
	users.GET("/cars")

	cars := router.Group("/cars")
	cars.GET("/list", api.getAllCars)
	cars.POST("/get-rent/:id", middleware.AuthMiddleware(api.jwtSigner), api.getRent)
	cars.GET("/rent-cars", middleware.AuthMiddleware(api.jwtSigner), api.getAvailableCars)
	cars.POST("add-car", middleware.AuthMiddleware(api.jwtSigner), api.addCar)

	router.POST("/refresh", api.refresh)
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello world!")
	})

	api.srv.Handler = router
}

func (api *RentApi) refresh(ctx *gin.Context) {
	refreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	claims, err := api.jwtSigner.ParseRefreshToken(refreshToken, auth.ParseOptions{
		ExpectedIssuer:   api.jwtSigner.Issuer,
		ExpectedAudience: api.jwtSigner.Audience,
		AllowMethods:     []string{"HS256"},
		Leeway:           60 * time.Second,
	})
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	access, err := api.jwtSigner.NewAccessToken(claims.Subject)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	newRefresh, err := api.jwtSigner.NewRefreshToken(claims.Subject)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.SetCookie("refresh_token", newRefresh, 3600*24*7, "/", "127.0.0.1:8080", false, true)
	ctx.JSON(http.StatusOK, gin.H{"access_token": access})
}

// делаем команаду docker-compose up
// подключаем бобра
