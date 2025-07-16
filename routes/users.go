package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zin-min-thu/go-rest-api/models"
	"github.com/zin-min-thu/go-rest-api/utils"
)

func signup(context *gin.Context) {

	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not save user", "err": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user created successful"})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}

	// context.JSON(http.StatusOK, user)
	// return

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "login successful", "token": token})

}
