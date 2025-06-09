package routes

import (
	"net/http"

	"Eventplanning.go/Api/models"
	"Eventplanning.go/Api/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "COuldn't parse the data"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Counldn't save the user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User succesfully created"})

}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't parse the request data"})
		return
	}

	err = user.ValidateUser()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't authenticate the user"})
		return
	}

	token, err := utils.GenrateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Can't authenticate the user"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "user logged in successfully", "token": token})
}
