package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"backend/controllers"
	"backend/middlewares"

	swaggerFiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	fileController := &controllers.FileController{DB: db}
	paperController := &controllers.PaperController{DB: db}
	r.Use(middlewares.CorsMiddleware())

	r.POST("/file", fileController.UploadFile)
	r.POST("/files", fileController.UploadFiles)
	r.GET("/file/:uuid", fileController.GetFile)
	r.DELETE("/file/:uuid", fileController.DeleteFile)

	r.POST("/papers", paperController.UploadPaper)
	r.GET("/papers", paperController.GetAllPapers)
	r.GET("/papers/file/:id", paperController.GetPaperFile)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger/doc.json")))

	return r
}
