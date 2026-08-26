package services

import (
	"context"
	"errors"
	"image"
	"time"

	"backend/converters"
	"backend/models"
	"backend/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
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

func (s *ConversionService) UpdateConversion(ctx context.Context, id, filename, tipo string) (models.ConversionRecord, error) {
	if filename == "" || tipo == " " {
		return models.ConversionRecord{}, errors.New("filename and type is obligatories for update ;v")
	}

	update := bson.M{
		"filename": filename,
		"type":     tipo,
	}

	return s.repo.UpdateByID(ctx, id, update)
}

func (s *ConversionService) UpdateTan(ctx context.Context, id string, filename, tipo *string) (models.ConversionRecord, error) {
	update := bson.M{}

	if filename != nil {
		update["filename"] = *filename
	}
	if tipo != nil {
		update["type"] = *tipo
	}

	if len(update) == 0 {
		return models.ConversionRecord{}, errors.New("not send data for update")

	}
	return s.repo.UpdateByID(ctx, id, update)
}

func (s *ConversionService) DeleteConversion(ctx context.Context, id string) error {
	return s.repo.DeleteByID(ctx, id)
}
