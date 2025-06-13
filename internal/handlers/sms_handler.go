package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"simple-go/internal/models"
)

func SendSms(ctx *gin.Context) {
	var smsSendVo models.ReqSendSmsVo
	if err := ctx.ShouldBindJSON(&smsSendVo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("发送短信：%+v", smsSendVo)
	data := models.RespSendSmsResult{
		RetCode: "0000",
		RetMsg:  "请求成功",
		RetData: models.RetData{
			Code:   "SMS-230816121258664609959060",
			Result: 1,
		},
	}

	ctx.JSON(http.StatusOK, data)

}
func QuerySms(ctx *gin.Context) {
	var queryVo models.ReqQueryVo
	if err := ctx.ShouldBindJSON(&queryVo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("查询短信：%+v", queryVo)

	retData := models.Detail{
		Phone:  "15261811090",
		Status: 1,
	}

	retDataList := []models.Detail{
		retData,
	}
	detailJson, err1 := json.Marshal(retDataList)
	if err1 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}

	result := models.RespQueryVo{
		RetCode: "0000",
		RetMsg:  "请求成功",
		RetData: string(detailJson),
	}

	ctx.JSON(http.StatusOK, result)
}
