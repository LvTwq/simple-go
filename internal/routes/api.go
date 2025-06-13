package routes

import (
	"github.com/gin-gonic/gin"
	"simple-go/internal/handlers"
	"simple-go/internal/models"
	"simple-go/pkg/middleware"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")

	{
		api.GET("/users", handlers.GetUsers)
	}

	open := r.Group("/openApi")
	{
		open.POST("/manage/getAccessToken", handlers.GetAccessToken)
	}

	sms := r.Group("/sms")
	{
		sms.POST("/external/interface/sendSms.do", handlers.SendSms)
		sms.POST("/external/interface/getSendSmsResult.do", handlers.QuerySms)
	}

	jwt := r.Group("/jwt")
	{
		// 登录接口 - 生成 token
		jwt.POST("/login", func(c *gin.Context) {
			// _ 表示忽略不需要的返回值，这里忽略的事 异常
			token, _ := middleware.GenerateToken("lvmc")

			// []中的是key的类型，后面是value的类型
			ajax := models.NewAjaxResult(200, map[string]string{
				"token": token,
			})
			c.JSON(200, ajax)
		})
		jwt.Use(middleware.JWTAuthMiddleware())
		jwt.GET("/protected", func(c *gin.Context) {
			username, _ := c.Get("username")

			ajax := models.NewAjaxResult(200, map[string]string{
				"username": username.(string),
			})
			c.JSON(200, ajax)
		})

	}
}
