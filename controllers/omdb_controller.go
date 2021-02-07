package controllers

import (
	"net/http"
	"omdbapi/models"
	"omdbapi/objects"
	"omdbapi/services"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type OmdbController struct {
	omdbService      services.OmdbService
	searchLogService services.SearchLogService
}

func OmdbControllerHandler() OmdbController {
	return OmdbController{
		omdbService:      services.OmdbServiceHandler(),
		searchLogService: services.SearchLogServiceHandler(),
	}
}

func (o OmdbController) GetList(context *gin.Context) {
	var params objects.OmdbAPIListRequestObject
	var err error

	queryParams := context.Request.URL.Query()
	params.Page, err = strconv.Atoi(queryParams.Get("page"))
	params.Search = queryParams.Get("search")
	if err != nil {
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	if params.Page == 0 {
		params.Page = 1
	}
	response, err := o.omdbService.GetList(params)

	if err != nil {
		context.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	go func(params objects.OmdbAPIListRequestObject) {

		time.Sleep(5 * time.Second)
		searchLog := models.SearchLogModel{
			SearchKey: params.Search,
			Page:      params.Page,
		}

		_, err = o.searchLogService.Create(searchLog)
		if err != nil {
			context.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
	}(params)

	context.JSON(http.StatusOK, response)
}
