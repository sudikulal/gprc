package main

import (
    "github.com/gin-gonic/gin"

    "journal/routes"
)

func main() {
    router := gin.Default()

    routes.UserRoutes(router)
    
    router.Run(":8080")
}
