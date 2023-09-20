package routes

import (
	"context"
	"net/http"

	"github.com/ZFP-Gaming/minion/initializers"
	"github.com/ZFP-Gaming/minion/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func FindUserByName(c *gin.Context) {
	var member models.Member
	var body struct {
		Name string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fields are not correct"})
		return
	}

	filter := bson.M{"name": body.Name}

	collection := initializers.DB.Collection("members")

	err := collection.FindOne(context.Background(), filter).Decode(&member)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, member)
}

func AllUsers(c *gin.Context) {
	var members []models.Member

	collection := initializers.DB.Collection("members")

	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while fetching users"})
		return
	}

	if err = cursor.All(context.Background(), &members); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while decoding users"})
		return
	}

	c.JSON(200, members)
}
