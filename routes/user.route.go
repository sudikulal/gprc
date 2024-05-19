package routes

import (
    "journal/controllers"
    "github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
    userGroup := router.Group("/user")
    {
        userGroup.POST("/register", controllers.UserRegister)
        userGroup.POST("/login", controllers.UserLogin)
        userGroup.POST("/logout", controllers.UserLogout)
        userGroup.POST("/verifyEmail", controllers.VerifyEmail)
    }
}
