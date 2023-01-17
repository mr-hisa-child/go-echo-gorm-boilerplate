//go:build wireinject
// +build wireinject

package di

import (
	"gorm.io/gorm"
	"henotech.net/presentation/controller"

	"henotech.net/infrastructure/datasource"
	"henotech.net/usecase"

	"github.com/google/wire"
)

func Init(db *gorm.DB) controller.TaskHandler {
	wire.Build(
		datasource.NewTaskRepository,
		usecase.NewTaskUsecase,
		controller.NewTaskHandler,
	)
	return nil
}
