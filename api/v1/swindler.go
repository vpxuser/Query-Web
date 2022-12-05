package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"query/model"
	"query/utils/errmsg"
	"query/utils/validator"
	"strconv"
)

//添加骗子信息
func AddSwindler(ctx *gin.Context) {
	var data model.Swindler
	ctx.ShouldBindJSON(&data)
	msg,code := validator.Validate(&data)
	if code != errmsg.SUCCESS {
		ctx.JSON(http.StatusOK,gin.H{
			"status" : code,
			"message" : msg,
		})
	} else {
		code = model.AddSwindler(&data)
		ctx.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
	}
}

func DeleteSwindlerById(ctx *gin.Context) {
	id,_ := strconv.Atoi(ctx.Param("id"))

	code := model.DeleteSwindlerById(id)

	ctx.JSON(http.StatusOK,gin.H{
		"status" : code,
		"message" : errmsg.GetErrMsg(code),
	})
}

func EditSwindlerById(ctx *gin.Context) {
	var data model.Swindler
	id,_ := strconv.Atoi(ctx.Param("id"))
	ctx.ShouldBindJSON(&data)
	code := model.EditSwindlerById(id,&data)

	ctx.JSON(http.StatusOK,gin.H{
		"status" : code,
		"message" : errmsg.GetErrMsg(code),
	})
}

func FindSwindlerById(ctx *gin.Context)  {
	id,_ := strconv.Atoi(ctx.Param("id"))
	var data = make(map[string]interface{})
	swindler,code := model.FindSwindlerById(id)
	data["wechat"] = swindler.Wechat
	data["phone"] = swindler.Phone
	ctx.JSON(http.StatusOK,gin.H{
		"status": code,
		"data": data,
		"total": 1,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetSwindlerListByTieba(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))
	tieba := ctx.Query("tieba")
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize < 1:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	var data []model.Swindler
	var total int64

	if tieba == "" {
		data,total = model.GetSwindlerList(pageSize, pageNum)
	} else {
		data,total = model.GetSwindlerListByTieba(tieba,pageSize, pageNum)
	}

	code := errmsg.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total": total,
		"message": errmsg.GetErrMsg(code),
	})
}