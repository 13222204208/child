package wechat

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AffirmRouter struct {
}

// InitAffirmRouter 初始化 Affirm 路由信息
func (s *AffirmRouter) InitAffirmRouter(Router *gin.RouterGroup) {
	affirmRouter := Router.Group("affirm").Use(middleware.OperationRecord())
	affirmRouterWithoutRecord := Router.Group("affirm")
	var affirmApi = v1.ApiGroupApp.WechatApiGroup.AffirmApi
	{
		affirmRouter.POST("createAffirm", affirmApi.CreateAffirm)             // 新建Affirm
		affirmRouter.DELETE("deleteAffirm", affirmApi.DeleteAffirm)           // 删除Affirm
		affirmRouter.DELETE("deleteAffirmByIds", affirmApi.DeleteAffirmByIds) // 批量删除Affirm
		affirmRouter.PUT("updateAffirm", affirmApi.UpdateAffirm)              // 更新Affirm
	}
	{
		affirmRouterWithoutRecord.GET("findAffirm", affirmApi.FindAffirm)       // 根据ID获取Affirm
		affirmRouterWithoutRecord.GET("getAffirmList", affirmApi.GetAffirmList) // 获取Affirm列表
		affirmRouterWithoutRecord.GET("list", affirmApi.List)                   // 获取Affirm列表
	}
}
