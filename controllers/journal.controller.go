package controllers

import (
	"journal/models"
	"journal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetJournalsList(c *gin.Context) {
	userId := c.GetHeader("userId")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid user"})
		return
	}

	folderId := c.Query("folder_id")

	findQuery := bson.M{"user_id": userId}

	if folderId != "" {
		findQuery["folder_id"] = folderId
	}

	findOptions := options.Find().SetProjection(bson.M{
		"_id":        1,
		"title":      1,
		"folder_id":  1,
		"day_rating": 1,
		"created_at": 1,
	})

	journalList, err := models.JournalModel.Find(c, findQuery, findOptions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to retrieve journals"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"journalList": journalList})
}

func GetJournalsDetail(c *gin.Context) {
	var journal *models.JournalSchema

	userId := c.GetHeader("userId")

	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid user"})
		return
	}

	journalId := c.Param("id")
	if journalId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	if err := models.JournalModel.FindOne(c, bson.M{"user_id": userId, "journal_id": journalId}).Decode(&journal); err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"message": "journal not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to retrieve journal"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"journalDetail": journal})
}

func CreateJournal(c *gin.Context) {
	var journalData models.JournalSchema

	userId, err := primitive.ObjectIDFromHex(c.GetHeader("userId"))

	if userId == primitive.NilObjectID || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid user"})
		return
	}

	if err := c.ShouldBind(&journalData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	journalData.UserID = userId

	if journalData.FolderID == primitive.NilObjectID {
		c.JSON(http.StatusBadRequest, gin.H{"message": "folder id cannot be empty"})
		return
	} else {
		var folderData models.FolderSchema
		if err := models.FolderModel.FindOne(c, bson.M{"_id": journalData.FolderID}).Decode(&folderData); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if folderData.ID == primitive.NilObjectID {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid folder id"})
			return
		}
	}

	encryptedData, err := utils.Encrypt(journalData.Content)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error while processing content"})
		return
	} else {
		journalData.Content = encryptedData
	}

	result, err := models.JournalModel.InsertOne(c, journalData)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}

func UpdateJournal(c *gin.Context) {
	userId, err := primitive.ObjectIDFromHex(c.GetHeader("userId"))

	if userId == primitive.NilObjectID || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid user"})
		return
	}

	journalId, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	var journal *models.JournalSchema

	c.ShouldBind(&journal)

	result, err := models.FolderModel.UpdateOne(c, bson.M{"_id": journalId, "user_id": userId}, journal)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "update failed"})
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}

func DeleteJournal(c *gin.Context) {
	userId := c.GetHeader("userId")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid user"})
		return
	}

	journalId:= c.Param("id")

	result, err := models.FolderModel.DeleteOne(c, bson.M{"_id": journalId, "user_id": userId})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "delete failed"})
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}
