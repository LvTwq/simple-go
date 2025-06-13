package services

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math/rand"
	"simple-go/internal/constant"
	"simple-go/internal/models"
	"simple-go/pkg/database"
	"simple-go/pkg/redis"
	"strings"
	"time"
)

func VerifyAppId(appId string, secret string) error {
	var openApiIdModel models.OpenApiIdModel
	db := database.DB
	result := db.Table("tb_openapi_appid").Where("app_id = ?", appId).First(&openApiIdModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("invalid appId")
		}
		return result.Error
	}

	if !strings.EqualFold(openApiIdModel.AppSecret, secret) {
		return errors.New("invalid secret")
	}
	return nil
}

func DoGetAccessToken(ctx *gin.Context, dto models.AccessTokenDto) (models.AccessTokenVo, error) {
	key := constant.OpenToken + dto.AppId
	redisClient := redis.RedisClient
	oldToken := redisClient.Get(ctx, key).Val()
	// 有老token
	if !(len(oldToken) == 0) {
		tokenMap := redisClient.HGetAll(ctx, oldToken).Val()

		randomString := generateRandomString(16)
		refreshAccessToken := "t66_" + md5Hash("Enlink-Refresh-"+randomString+"-API")
		tokenMap["refreshToken"] = refreshAccessToken

		redisClient.HSet(ctx, refreshAccessToken, tokenMap)
		redisClient.Expire(ctx, refreshAccessToken, 7*24*3600)

		expire := redisClient.TTL(ctx, key).Val()
		return models.AccessTokenVo{
			Token:        oldToken,
			ExpireIn:     int64(expire.Seconds()),
			RefreshToken: refreshAccessToken,
		}, nil
	}

	// 生成token
	newAccessToken := "enaccess-" + md5Hash("Enlink-"+generateRandomString(16)+"-API")
	refreshAccessToken := "t66_" + md5Hash("Enlink-Refresh-"+generateRandomString(16)+"-API")
	newTokenMap := map[string]string{
		"appId":        dto.AppId,
		"appSecret":    dto.AppSecret,
		"refreshToken": refreshAccessToken,
	}
	redisClient.Set(ctx, key, newAccessToken, constant.ExpiresIn*time.Second)
	redisClient.HSet(ctx, newAccessToken, newTokenMap)
	redisClient.Expire(ctx, newAccessToken, constant.ExpiresIn*time.Second)

	return models.AccessTokenVo{
		Token:        newAccessToken,
		ExpireIn:     constant.ExpiresIn,
		RefreshToken: refreshAccessToken,
	}, nil
}

func md5Hash(text string) string {
	harsher := md5.New()
	harsher.Write([]byte(text))
	return hex.EncodeToString(harsher.Sum(nil))
}

func generateRandomString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
