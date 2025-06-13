package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"simple-go/internal/models"
	"simple-go/internal/services"
)

func GetAccessToken(ctx *gin.Context) {
	var accessTokenDto models.AccessTokenDto
	if err := ctx.ShouldBindJSON(&accessTokenDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("申请token: %v", accessTokenDto)

	if err := services.VerifyAppId(accessTokenDto.AppId, accessTokenDto.AppSecret); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	{
		result, err := services.DoGetAccessToken(ctx, accessTokenDto)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": result})
	}

}
