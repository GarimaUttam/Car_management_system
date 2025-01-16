package engine

import (
	"context"

	"github.com/GarimaUttam/Car_management_system/models"
	"github.com/GarimaUttam/Car_management_system/store"
)

type EngineService struct {
	store store.EngineStoreInterface
}

func NewEngineService(store store.EngineStoreInterface) *EngineService {
	return &EngineService{
		store: store,
	}
}

func(s *EngineService) GetEngineByID(ctx context.Context,id string) (*models.Engine, error) {
	engine, err := s.store.EngineById(ctx,id)
	if err != nil {
		return nil,err
	}
	return &engine, nil
}

func (s *EngineService) CreateEngine(ctx context.Context, engineReq *models.EngineRequest) (*models.Engine, error){
	if err := models.ValidateEngineRequest(*engineReq); err != nil {
		return nil, err
	}

	createdEngine, err := s.store.EngineCreate(ctx, engineReq)
	if err != nil {
		return nil, err 
	}
	return &createdEngine, nil
}

func (s *EngineService) UpdateEngine(ctx context.Context, id string, engineReq *models.EngineRequest) (*models.Engine, error,) {
	if err := models.ValidateEngineRequest(*engineReq); err != nil {
		return nil, err
	}
	updateEngine, err := s.store.EngineUpdate(ctx, id, engineReq)
	if err != nil {
		return nil, err
	}
	return &updateEngine, nil
}
func (s *EngineService) DeletedEngine(ctx context.Context, id string) (*models.Engine, error) {
	deletedEngine, err := s.store.EngineDelete(ctx, id)
	if err != nil {
		return nil, err
	}
	return &deletedEngine, nil
}