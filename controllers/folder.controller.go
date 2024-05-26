package controllers

import (
	"journal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetFoldersList(c *gin.Context) {
	userId := c.GetHeader("userId")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid user"})
		return
	}

	folderList, err := models.JournalModel.Find(c, bson.M{"user_id": userId})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to retrieve folders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"journalList": folderList})
}

func CreateFolder(c *gin.Context) {

}
func UpdateFolder(c *gin.Context) {

}
func DeleteFolder(c *gin.Context) {

}
