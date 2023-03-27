package controller

import (
	"bluebell/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

//跟社区相关

func CommunityHandler(c *gin.Context) {
	//查询到所有社区（id，name）以列表形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) //不轻易将数据库层的东西暴露给服务
		return
	}
	ResponseSuccess(c, data)
}

// CommunityDetailHandler 社区分类详情
func CommunityDatailHandler(c *gin.Context) {
	//1、获取社区id
	idStr := c.Param("id")                     //获取url参数（string型）
	id, err := strconv.ParseInt(idStr, 10, 64) //string转int64
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	//根据获取的id获取社区详情
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c, data)
}
