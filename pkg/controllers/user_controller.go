// controllers/user_controller.go
package controllers

import (
	"net/http"

	"github.com/go-firebase-api-template/pkg/models"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context, firestoreClient *firestore.Client) {
	// Parse request body to get user data
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add user document to Firestore
	_, _, err := firestoreClient.Collection("users").Add(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add user to Firestore"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User added successfully"})
}

func GetAllUsers(c *gin.Context, firestoreClient *firestore.Client) {

	// Query all users from Firestore
	usersSnapshot, err := firestoreClient.Collection("users").Documents(c).GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users from Firestore"})
		return
	}

	var users []models.User
	for _, doc := range usersSnapshot {
		var user models.User
		if err := doc.DataTo(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert user data"})
			return
		}
		user.ID = doc.Ref.ID
		users = append(users, user)
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func GetUserByID(c *gin.Context, firestoreClient *firestore.Client) {
	// Get user ID from the request params
	userID := c.Param("id")

	// Query user from Firestore by ID
	doc, err := firestoreClient.Collection("users").Doc(userID).Get(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user from Firestore"})
		return
	}

	// Convert Firestore document to User object
	var user models.User
	if err := doc.DataTo(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert user data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func EditUser(c *gin.Context, firestoreClient *firestore.Client) {
	// Get user ID from the request params
	userID := c.Param("id")

	// Parse request body to get updated user data
	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update user document in Firestore
	_, err := firestoreClient.Collection("users").Doc(userID).Set(c, updatedUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user in Firestore"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUser(c *gin.Context, firestoreClient *firestore.Client) {
	// Get user ID from the request params
	userID := c.Param("id")

	// Delete user document from Firestore
	_, err := firestoreClient.Collection("users").Doc(userID).Delete(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user from Firestore"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
