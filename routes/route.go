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
	// group.GET("user/:uuid", user.Get)
	// group.PUT("user/:uuid", user.Update)
	// group.DELETE("user/:uuid", user.Delete)

}
