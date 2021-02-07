package repositories

import (
	"errors"
	"fmt"
	"omdbapi/config"
	"omdbapi/models"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type SearchLogRepository struct {
	DB    *gorm.DB
	Model models.SearchLogModel
}

func SearchLogRepositoryHandler() SearchLogRepository {
	db := config.GetCon()

	repository := SearchLogRepository{
		DB:    db,
		Model: models.SearchLogModel{},
	}
	return repository
}

func (r *SearchLogRepository) Create(model models.SearchLogModel) (models.SearchLogModel, error) {
	model.UUID = fmt.Sprintf("%s", uuid.NewV4())
	dt := time.Now()
	model.CreatedAt = dt.String()
	if err := r.DB.Create(model).Error; err != nil {
		return models.SearchLogModel{}, err
	}
	return r.GetByUUID(model.UUID)
}

func (r *SearchLogRepository) GetByUUID(UUID string) (models.SearchLogModel, error) {
	var model models.SearchLogModel
	r.DB.Where("uuid = ?", UUID).First(&model)
	if model.UUID == "" {
		return models.SearchLogModel{}, errors.New("data not found")
	}
	return model, nil
}
