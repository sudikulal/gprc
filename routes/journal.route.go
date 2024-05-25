package routes

import (
	"journal/controllers"
	"journal/middleware"

	"github.com/gin-gonic/gin"
)

func JournalRoutes(router *gin.Engine) {
    router.Use(middleware.Auth)
    
    journalGroup := router.Group("/journals")
    {
        journalGroup.GET("/:id", controllers.GetJournals)
        journalGroup.POST("/", controllers.CreateJournal)
        journalGroup.PUT("/:id", controllers.UpdateJournal)
        journalGroup.DELETE("/:id", controllers.DeleteJournal)
    }
}
