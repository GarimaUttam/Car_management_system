package service

import (
	"context"

	"github.com/GarimaUttam/Car_management_system/models"
)

type CarServiceInterface interface {
	GetCarByID(ctx context.Context, id string) (*models.Car, error)
	GetCarsByBrand(ctx context.Context, brand string, isEngine bool) ([]models.Car, error)
	CreateCar(ctx context.Context, car *models.CarRequest) (*models.Car, error)
	UpdateCar(ctx context.Context, id string, carReq *models.CarRequest) (*models.Car, error)
	DeleteCar(ctx context.Context, id string) (*models.Car, error) 
}

type EngineServiceInterface interface {
	GetEngineByID(ctx context.Context,id string) (*models.Engine, error) 
	CreateEngine(ctx context.Context, engineReq *models.EngineRequest) (*models.Engine, error)
	UpdateEngine(ctx context.Context, id string, engineReq *models.EngineRequest) (*models.Engine, error,)
	DeletedEngine(ctx context.Context, id string) (*models.Engine, error)
}
