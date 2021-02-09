package routes

import (
	"omdbapi/controllers"

	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine) {
	// defaultMiddleware := middleware.DefaultMiddleware{}

	omdb := controllers.OmdbControllerHandler()
	group := router.Group("/v1")
	group.GET("/movie-list", omdb.GetList)
	group.GET("/movie-detail/:id", omdb.GetDetail)

}
