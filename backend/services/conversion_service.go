package services

import (
	"context"
	"image"
	"time"

	"backend/converters"
	"backend/models"
	"backend/repository"
)

type ConversionService struct {
	repo *repository.ConversionRepository
}

func NewConversionService(repo *repository.ConversionRepository) *ConversionService {
	return &ConversionService{repo: repo}
}

func (s *ConversionService) ProcesarImagen(ctx context.Context, img image.Image, tipo string, filename string) (models.ConversionRecord, error) {
	conv, err := converters.NewConverter(tipo)
	if err != nil {
		return models.ConversionRecord{}, err
	}

	resultado, err := conv.Convert(img)
	if err != nil {
		return models.ConversionRecord{}, err
	}

	record := models.ConversionRecord{
		FileName:  filename,
		Type:      models.ConversionType(tipo),
		Result:    resultado,
		CreatedAt: time.Now(),
	}
	return s.repo.Save(ctx, record)
}

func (s *ConversionService) ListarConversiones(ctx context.Context) ([]models.ConversionRecord, error) {
	return s.repo.FindAll(ctx)
}

func (s *ConversionService) ObtenerConversion(ctx context.Context, id string) (models.ConversionRecord, error) {
	return s.repo.FindByID(ctx, id)
}
