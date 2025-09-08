package server

import (
	"BrunoyamLesson6/internal"
	userDomain "BrunoyamLesson6/internal/domain/users/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserStorage interface {
	SaveUser(user userDomain.User) error
	GetUser(userReq userDomain.UserRequest) (userDomain.User, error)
}

type RentApi struct {
	srv *http.Server
	db  UserStorage
}

func NewServer(cfg internal.Config, db UserStorage) *RentApi {
	HttpSrv := http.Server{
		Addr: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
	}

	api := RentApi{srv: &HttpSrv, db: db}

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
	users.POST("/login", api.Login)
	users.POST("/register", api.Register)
	users.GET("/profile")
	users.GET("/cars")

	cars := router.Group("/cars")
	cars.GET("/list")
	cars.POST("/get_rent")
	cars.GET("/rent_cars")
	cars.POST("add_car")

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello world!")
	})

	api.srv.Handler = router
}
