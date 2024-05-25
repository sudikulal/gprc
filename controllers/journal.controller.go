package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetJournals(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hello": "world"})
}

func CreateJournal(c *gin.Context) {

}
func UpdateJournal(c *gin.Context) {

}

func DeleteJournal(c *gin.Context) {

}
