package http

import (
	"pari/internal/config"

	cCrud "pari/internal/controller/crud"
	"pari/internal/domain"
	rCrud "pari/internal/repository/crud"
	uCrud "pari/internal/usecase/crud"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	repoCrud domain.CrudRepository

	usecaseCrud domain.CrudUsecase
)

func InitializeRepositories(db *gorm.DB) {

	repoCrud = rCrud.New(db)

}

func InitializeUsecases(config config.Config) {

	usecaseCrud = uCrud.New(repoCrud)

}

func InitializeControllers(router *gin.Engine, config config.Config) {

	publicApiV1_0 := router.Group("public/api/v1.0")
	{
		registerPublicApiV1_0(publicApiV1_0, config)
	}

}

func registerPublicApiV1_0(v1_0 *gin.RouterGroup, config config.Config) {
	controllerCrud := cCrud.New(usecaseCrud)

	RoutesCrud := v1_0.Group("crud")
	{
		RoutesCrud.GET("/all", controllerCrud.GetAll)
		RoutesCrud.GET("/", controllerCrud.GetById)
		RoutesCrud.POST("/items", controllerCrud.Post)
		RoutesCrud.PUT("/items", controllerCrud.Put)
		RoutesCrud.DELETE("/items", controllerCrud.Delete)
	}

}
