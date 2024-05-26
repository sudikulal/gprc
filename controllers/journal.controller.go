package controllers

import (
	"journal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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

}
func UpdateJournal(c *gin.Context) {

}

func DeleteJournal(c *gin.Context) {

}
