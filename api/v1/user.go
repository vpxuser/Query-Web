package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"query/model"
	"query/utils/errmsg"
	"query/utils/validator"
	"strconv"
)

// 添加用户
func AddUser(ctx *gin.Context) {
	var data model.User
	ctx.ShouldBindJSON(&data)
	msg,code := validator.Validate(&data)
	if code != errmsg.SUCCESS {
		ctx.JSON(http.StatusOK,gin.H{
			"status" : code,
			"message" : msg,
		})
	} else {
		code = model.FindUserByUsername(data.Username)
		if code == errmsg.SUCCESS {
			code = model.AddUser(&data)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
	}
}

// 删除用户
func DeleteUserById(ctx *gin.Context) {
	id,_ := strconv.Atoi(ctx.Param("id"))

	code := model.DeleteUserById(id)

	ctx.JSON(http.StatusOK,gin.H{
		"status" : code,
		"message" : errmsg.GetErrMsg(code),
	})
}

// 编辑用户
func EditUserById(ctx *gin.Context) {
	var data model.User
	id,_ := strconv.Atoi(ctx.Param("id"))
	ctx.ShouldBindJSON(&data)
	code := model.FindUserByUsername(data.Username)
	if code == errmsg.SUCCESS {
		model.EditUserById(id,&data)
	}

	ctx.JSON(http.StatusOK,gin.H{
		"status" : code,
		"message" : errmsg.GetErrMsg(code),
	})
}

// 查询单个用户
func FindUserById(ctx *gin.Context)  {
	id,_ := strconv.Atoi(ctx.Param("id"))
	var data = make(map[string]interface{})
	user,code := model.FindUserById(id)
	data["username"] = user.Username
	ctx.JSON(http.StatusOK,gin.H{
		"status": code,
		"data": data,
		"total": 1,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询用户列表
func GetUserListByUsername(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))
	username := ctx.Query("username")
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize < 1:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data,total := model.GetUserListByUsername(username,pageSize, pageNum)

	code := errmsg.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total": total,
		"message": errmsg.GetErrMsg(code),
	})
}

// 修改密码
func RePassword(ctx *gin.Context)  {
	var data model.User
	id,_ := strconv.Atoi(ctx.Param("id"))
	ctx.ShouldBindJSON(&data)

	code := model.RePassword(id,&data)

	ctx.JSON(http.StatusOK,gin.H{
		"status": code,
		"message": errmsg.GetErrMsg(code),
	})
}