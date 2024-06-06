package main

import (
    "github.com/gin-gonic/gin"

    "journal/routes"
)

func main() {
    router := gin.Default()

    routes.UserRoutes(router)
    routes.JournalRoutes(router)
    routes.FolderRoutes(router)

    router.Run(":8080")
}
