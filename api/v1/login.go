package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"query/middleware"
	"query/model"
	"query/utils/errmsg"
)

// 登录后台
func Login(ctx *gin.Context)  {
	var data model.User
	ctx.ShouldBindJSON(&data)
	var token string

	data,code := model.Login(data.Username,data.Password)

	if code == errmsg.SUCCESS {
		middleware.SetToken(ctx,data)
	} else {
		ctx.JSON(http.StatusOK,gin.H{
			"status" : code,
			"username": data.Username,
			"id": data.ID,
			"message" : errmsg.GetErrMsg(code),
			"token" : token,
		})
	}
}