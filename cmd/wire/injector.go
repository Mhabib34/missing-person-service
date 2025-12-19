//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/Mhbib34/missing-person-service/internal/controller"
	"github.com/Mhbib34/missing-person-service/internal/database"
	"github.com/Mhbib34/missing-person-service/internal/repository"
	"github.com/Mhbib34/missing-person-service/internal/router"
	"github.com/Mhbib34/missing-person-service/internal/usecase"
	"github.com/Mhbib34/missing-person-service/internal/worker"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"gorm.io/gorm"
)


type App struct {
	DB     *gorm.DB
	Router *gin.Engine
	Worker *worker.ResizeImageJobWorker
}

func NewValidator() *validator.Validate {
	return validator.New()
}


var repositorySet = wire.NewSet(
	repository.NewMissingPersonRepository,
)

var usecaseSet = wire.NewSet(
	usecase.NewMissingPersonUsecase,
)

var controllerSet = wire.NewSet(
	controller.NewMissingPersonController,
)

var routerSet = wire.NewSet(
	router.SetupRouter,
)

func provideResizeImageWorker(db *gorm.DB) *worker.ResizeImageJobWorker {
	return worker.NewResizeImageJobWorker(db, 5)
}

func InitializeServer() (*App, error) {
	wire.Build(
		// Database
		database.Connect,

		// Validator
		NewValidator,

		// Layers
		repositorySet,
		usecaseSet,
		controllerSet,
		routerSet,

		// Worker
		provideResizeImageWorker,

		// App struct
		wire.Struct(new(App), "*"),
	)
	return nil, nil
}
