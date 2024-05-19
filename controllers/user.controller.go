package controllers

import (
	"journal/config"
	"journal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func UserRegister(c *gin.Context) {
	var user models.UserSchema
	var err error
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	emailId, password, userName, registerType := user.EmailId, user.Password, user.UserName, user.RegisterType

	var foundUser models.UserSchema

	switch registerType {
	case config.REGISTER_TYPE["EMAIL"]:
		{
			if emailId == "" || password == "" {
				c.JSON(http.StatusBadRequest, gin.H{"message": "emailId/password is empty"})
				return
			}

			err = models.UserModel.FindOne(c, bson.M{"email_id": emailId}).Decode(&foundUser)

			if foundUser.UserName != "" {
				c.JSON(http.StatusBadRequest, gin.H{"message": "userName is already exist"})
				return
			}
		}
	case config.REGISTER_TYPE["USER_NAME"]:
		{
			if userName == "" || password == "" {
				c.JSON(http.StatusBadRequest, gin.H{"message": "userName/password is empty"})
				return
			}

			err = models.UserModel.FindOne(c, bson.M{"user_name": userName}).Decode(&foundUser)

			if foundUser.UserName != "" {
				c.JSON(http.StatusBadRequest, gin.H{"message": "userName is already exist"})
				return
			}
		}

	default:
		{
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid register type"})
			return
		}

	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	createUser, err := models.UserModel.InsertOne(c, user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createUser)
}

func UserLogin(c *gin.Context) {
	var user models.UserSchema
	var err error
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	emailId, password, userName, registerType := user.EmailId, user.Password, user.UserName, user.RegisterType

	var foundUser models.UserSchema

	switch registerType {
	case config.REGISTER_TYPE["EMAIL"]:
		{
			if emailId == "" || password == "" {
				c.JSON(http.StatusBadRequest, gin.H{"message": "Email/password is empty"})
				return
			}

			err = models.UserModel.FindOne(c, bson.M{"email_id": emailId}).Decode(&foundUser)

		}
	case config.REGISTER_TYPE["USER_NAME"]:
		{
			if userName == "" || password == "" {
				c.JSON(http.StatusBadRequest, gin.H{"message": "userName/password is empty"})
				return
			}

			err = models.UserModel.FindOne(c, bson.M{"user_name": userName}).Decode(&foundUser)
		}

	default:
		{
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid register type"})
			return
		}

	}

	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": foundUser})
}

func UserLogout(c *gin.Context) {

}
func VerifyEmail(c *gin.Context) {

}
