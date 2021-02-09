package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"omdbapi/objects"
	"os"
)

type OmdbService struct {
}

func OmdbServiceHandler() OmdbService {
	return OmdbService{}
}

func (o OmdbService) GetList(params objects.OmdbAPIListRequestObject) (objects.OmdbAPIListResponseObject, error) {

	if params.Page == 0 {
		params.Page = 1
	}

	url := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&s=%s&page=%d", os.Getenv("OMDB_API_KEY"), params.Search, params.Page)
	requestToOmdb, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return objects.OmdbAPIListResponseObject{}, err
	}
	httpClient := &http.Client{}
	response, err := httpClient.Do(requestToOmdb)
	if err != nil {
		return objects.OmdbAPIListResponseObject{}, err
	}
	fmt.Printf("%v \n", response.Body)
	if int(response.StatusCode/100) != 2 {
		return objects.OmdbAPIListResponseObject{}, fmt.Errorf("status %d", response.StatusCode)
	}
	var omdbAPIResponse objects.OmdbAPIListResponseObject
	err = json.NewDecoder(response.Body).Decode(&omdbAPIResponse)
	if err != nil {
		return objects.OmdbAPIListResponseObject{}, err
	}
	return omdbAPIResponse, nil
}

func (o OmdbService) GetDetail(id string) (objects.OmdbAPIDetailResponseObject, error) {
	url := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&i=%s", os.Getenv("OMDB_API_KEY"), id)
	requestToOmdb, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return objects.OmdbAPIDetailResponseObject{}, err
	}
	httpClient := &http.Client{}
	response, err := httpClient.Do(requestToOmdb)
	if err != nil {
		return objects.OmdbAPIDetailResponseObject{}, err
	}
	fmt.Printf("%v \n", response.Body)
	if int(response.StatusCode/100) != 2 {
		return objects.OmdbAPIDetailResponseObject{}, fmt.Errorf("status %d", response.StatusCode)
	}
	var omdbAPIResponse objects.OmdbAPIDetailResponseObject
	err = json.NewDecoder(response.Body).Decode(&omdbAPIResponse)
	if err != nil {
		return objects.OmdbAPIDetailResponseObject{}, err
	}
	return omdbAPIResponse, nil
}
