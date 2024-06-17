package routes

import (
	"journal/controllers"
	"journal/middleware"

	"github.com/gin-gonic/gin"
)

func FolderRoutes(router *gin.Engine) {
    router.Use(middleware.Auth)
    
    folderGroup := router.Group("/folders")
    {
        folderGroup.GET("/", controllers.GetFoldersList)
        folderGroup.POST("/", controllers.CreateFolder)
        folderGroup.PUT("/:id", controllers.UpdateFolder)
        folderGroup.DELETE("/:id", controllers.DeleteFolder)
    }
}
