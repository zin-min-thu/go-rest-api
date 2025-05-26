package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zin-min-thu/go-rest-api/models"
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
