package controllers

import (
	"journal/constant"
	"journal/models"
	"journal/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetTodoList(c *gin.Context) {
	userObj, err := utils.GetUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	cursor, err := models.TodoModel.Find(c, bson.M{"user_id": userObj.UserId,
		"status": bson.M{"$in": []int{constant.STATUS["ACTIVE"], constant.STATUS["COMPLETED"]}}})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to retrieve todos"})
		return
	}
	defer cursor.Close(c)

	var todoList []models.TodoSchema

	if err = cursor.All(c, &todoList); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to decode todos"})
		return
	}

	c.JSON(http.StatusOK, todoList)
}

func CreateTodo(c *gin.Context) {
	userObj, err := utils.GetUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var todo models.TodoSchema

	if err := c.ShouldBind(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if todo.Title == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "title is empty"})
		return
	}

	todo.UserID = userObj.UserId
	todo.Status = constant.STATUS["ACTIVE"]
	todo.CreatedAt = time.Now()

	createTodo, err := models.TodoModel.InsertOne(c, todo)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createTodo)
}

func UpdateTodo(c *gin.Context) {
	userObj, err := utils.GetUserFromContext(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	todoId := utils.GetIdFromParam(c)

	var todo *models.TodoSchema

	if err := c.ShouldBind(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	updateData := bson.M{}

	if todo.Title != "" {
		updateData["title"] = todo.Title
	}

	if todo.Status != 0 {
		includes := false
		for _, value := range constant.STATUS {
			if value == todo.Status {
				includes = true
				break
			}
		}
		if !includes {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid status"})
			return
		}
		updateData["status"] = todo.Status
	}

	if len(updateData) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "update is empty"})
		return
	}

	result, err := models.TodoModel.UpdateOne(c, bson.M{"_id": todoId, "user_id": userObj.UserId}, bson.M{"$set": updateData})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "update failed"})
		return
	}

	c.JSON(http.StatusOK, result)
}
