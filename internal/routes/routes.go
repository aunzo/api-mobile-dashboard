package routes

import (
	"github.com/aunz/api-mobile-dashboard-golang/internal/handlers"
	"github.com/gin-gonic/gin"

	_ "github.com/aunz/api-mobile-dashboard-golang/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(engine *gin.Engine, buildInfoHandler *handlers.BuildInfoHandler) {
	engine.StaticFile("/", "./static/index.html")

	engine.GET("/build-info", buildInfoHandler.BuildInfoList)
	engine.GET("/build-info/csv", buildInfoHandler.BuildInfoListCSV)
	engine.POST("/build-info", buildInfoHandler.CreateBuildInfo)

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
