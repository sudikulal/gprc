package controllers

import (
	"context"
	"net/http"

	"journal/models"
	"journal/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetFoldersList(c *gin.Context) {
	userObj, err := utils.GetUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	cursor, err := models.FolderModel.Find(c, bson.M{"user_id": userObj.UserId}, options.Find())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to retrieve folders"})
		return
	}
	defer cursor.Close(context.Background())

	var folderList []models.FolderSchema
	for cursor.Next(context.Background()) {
		var folder models.FolderSchema
		if err = cursor.Decode(&folder); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to decode folder"})
			return
		}
		folderList = append(folderList, folder)
	}

	if err = cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cursor error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"journalList": folderList})
}


func CreateFolder(c *gin.Context) {
	userObj, err := utils.GetUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var folder models.FolderSchema

	if err := c.ShouldBind(&folder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if folder.FolderName == "" {
		c.JSON(http.StatusOK, gin.H{"message": "folder name is empty"})
		return
	}

	folder.UserID = userObj.UserId

	var folderExist models.FolderSchema
	if err := models.FolderModel.FindOne(c, bson.M{"folder_name": folder.FolderName, "user_id": userObj.UserId}).Decode(&folderExist); err == nil {
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

	c.JSON(http.StatusCreated, createFolder)

}

func UpdateFolder(c *gin.Context) {
	userObj, err := utils.GetUserFromContext(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	folderId := utils.GetIdFromParam(c)

	var folder *models.FolderSchema

	c.ShouldBind(&folder)

	if folder.FolderName == "" {
		c.JSON(http.StatusOK, gin.H{"message": "folder name is empty"})
		return
	}

	result, err := models.FolderModel.UpdateOne(c, bson.M{"_id": folderId, "user_id": userObj.UserId}, bson.M{"$set": bson.M{"folder_name": folder.FolderName}})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "update failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})

}

func DeleteFolder(c *gin.Context) {
	userId := c.GetHeader("userId")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid user"})
		return
	}

	folderId := c.Param("id")

	result, err := models.FolderModel.DeleteOne(c, bson.M{"_id": folderId, "user_id": userId})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "delete failed"})
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}
