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
	userId := c.GetHeader("userId")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid user"})
		return
	}

	var folder models.FolderSchema

	if err := c.ShouldBind(&folder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var folderExist models.FolderSchema
	if err := models.FolderModel.FindOne(c, bson.M{"folder_name": folder.FolderName, "user_id": userId}).Decode(&folderExist); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Folder already exists"})
		return
	} else if err != mongo.ErrNoDocuments {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	createFolder, err := models.FolderModel.InsertOne(c, folder)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"folder": createFolder})

}

func UpdateFolder(c *gin.Context) {

}
func DeleteFolder(c *gin.Context) {

}
