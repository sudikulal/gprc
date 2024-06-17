package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	UserId primitive.ObjectID
}

func GetUserFromContext(c *gin.Context) (*User, error) {
	claims, exists := c.Get("userObj")

	if !exists {
		return nil, fmt.Errorf("user not found in context")
	}

	mapClaims, ok := claims.(jwt.MapClaims)

	if !ok {
		return nil, fmt.Errorf("invalid user claims type")
	}

	userId, ok := mapClaims["userId"].(string)

	if !ok {
		return nil, fmt.Errorf("userId not found in claims or is not a string")
	}

	objID, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return nil, fmt.Errorf("invalid userId format: %v", err)
	}

	return &User{
	 	UserId: objID,
	}, nil
}

func GetIdFromParam(c *gin.Context) (primitive.ObjectID){
	id := c.Param("id")
	objId, _ := primitive.ObjectIDFromHex(id)
	return objId
}
