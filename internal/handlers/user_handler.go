package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-go/internal/services"
)

func GetUsers(context *gin.Context) {
	users := services.GetAllUsers()
	context.JSON(http.StatusOK, gin.H{"users": users})
}
