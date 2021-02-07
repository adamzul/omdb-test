package services

import (
	"omdbapi/models"
	"omdbapi/repositories"
)

type SearchLogService struct {
	searchLogRepository repositories.SearchLogRepository
}

func SearchLogServiceHandler() SearchLogService {
	service := SearchLogService{
		searchLogRepository: repositories.SearchLogRepositoryHandler(),
	}
	return service
}

func (service *SearchLogService) Create(request models.SearchLogModel) (models.SearchLogModel, error) {

	result, err := service.searchLogRepository.Create(request)

	if err != nil {
		return models.SearchLogModel{}, err
	}
	return result, nil
}
