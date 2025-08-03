package services

import (
	"context"
	"smart-city/internal/app/camera/dto"
	camRepositories "smart-city/internal/app/camera/repository"
	premiseRepositories "smart-city/internal/app/premise/repository"
	"smart-city/internal/models"

	"github.com/google/uuid"
)

type Service struct {
	cameraRepo  camRepositories.CameraRepository
	premiseRepo premiseRepositories.PremiseRepository
}

func NewCameraService(cameraRepo camRepositories.CameraRepository, premiseRepo premiseRepositories.PremiseRepository) *Service {
	return &Service{cameraRepo: cameraRepo, premiseRepo: premiseRepo}
}

func (s *Service) CreateCamera(ctx context.Context, createCameraDto *dto.CreateCameraDto) (*models.Camera, error) {
	Camera := &models.Camera{
		Name:      createCameraDto.Name,
		Location:  createCameraDto.Location,
		StreamURL: createCameraDto.StreamURL,
		IsActive:  true,
	}
	if createCameraDto.PremiseID != "" {
		premiseID, err := uuid.Parse(createCameraDto.PremiseID)

		if err != nil {
			return nil, err
		}

		premise, err := s.premiseRepo.GetPremiseByID(ctx, premiseID.String())

		if err != nil {
			return nil, err
		}
		Camera.Premise = premise
		Camera.PremiseID = premiseID

	}

	return s.cameraRepo.CreateCamera(ctx, Camera)
}

func (s *Service) GetCameras(ctx context.Context) ([]models.Camera, error) {
	return s.cameraRepo.GetCameras(ctx)
}
