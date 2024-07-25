package controllers

import (
	"journal/config"
	"journal/models"
	"journal/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func UserRegister(c *gin.Context) {
	var user models.UserSchema

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	emailId, password, userName, registerType := user.EmailId, user.Password, user.UserName, user.RegisterType

	if (registerType == config.REGISTER_TYPE["EMAIL"] && (emailId == "" || password == "")) ||
		(registerType == config.REGISTER_TYPE["USER_NAME"] && (userName == "" || password == "")) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Required fields are empty"})
		return
	}

	filter := bson.M{}
	if registerType == config.REGISTER_TYPE["EMAIL"] {
		filter["email_id"] = emailId
	} else if registerType == config.REGISTER_TYPE["USER_NAME"] {
		filter["user_name"] = userName
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid register type"})
		return
	}

	var foundUser models.UserSchema
	if err := models.UserModel.FindOne(c, filter).Decode(&foundUser); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User already exists"})
		return
	} else if err != mongo.ErrNoDocuments {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Password encryption failed"})
		return
	}
	user.Password = string(hashedPassword)
	user.CreatedAt = time.Now()

	createUser, err := models.UserModel.InsertOne(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	objectID := createUser.InsertedID.(primitive.ObjectID).Hex()

	userObj := map[string]interface{}{
		"userId": objectID,
	}

	token, err := utils.CreateJwtToken(userObj, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"name":        user.UserName,
		"user_id":     objectID,
		"accessToken": token,
	})
}

func UserLogin(c *gin.Context) {
	var user models.UserSchema
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	emailId, password, userName, registerType := user.EmailId, user.Password, user.UserName, user.RegisterType

	if (registerType == config.REGISTER_TYPE["EMAIL"] && (emailId == "" || password == "")) ||
		(registerType == config.REGISTER_TYPE["USER_NAME"] && (userName == "" || password == "")) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Required fields are empty"})
		return
	}

	filter := bson.M{}
	if registerType == config.REGISTER_TYPE["EMAIL"] {
		filter["email_id"] = emailId
	} else if registerType == config.REGISTER_TYPE["USER_NAME"] {
		filter["user_name"] = userName
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid register type"})
		return
	}

	var foundUser models.UserSchema
	if err := models.UserModel.FindOne(c, filter).Decode(&foundUser); err == mongo.ErrNoDocuments {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid email or password"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid email or password"})
		return
	}

	objectID := (foundUser.ID).Hex()

	userObj := map[string]interface{}{
		"userId": objectID,
	}

	token, err := utils.CreateJwtToken(userObj, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":        user.UserName,
		"user_id":     objectID,
		"accessToken": token,
	})
}

func UserLogout(c *gin.Context) {

}
func VerifyEmail(c *gin.Context) {

}
