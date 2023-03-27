package wechat

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AgencyRouter struct {
}

// InitAgencyRouter 初始化 Agency 路由信息
func (s *AgencyRouter) InitAgencyRouter(Router *gin.RouterGroup) {
	agencyRouter := Router.Group("agency").Use(middleware.OperationRecord())
	agencyRouterWithoutRecord := Router.Group("agency")
	var agencyApi = v1.ApiGroupApp.WechatApiGroup.AgencyApi
	{
		agencyRouter.POST("createAgency", agencyApi.CreateAgency)             // 新建Agency
		agencyRouter.DELETE("deleteAgency", agencyApi.DeleteAgency)           // 删除Agency
		agencyRouter.DELETE("deleteAgencyByIds", agencyApi.DeleteAgencyByIds) // 批量删除Agency
		agencyRouter.PUT("updateAgency", agencyApi.UpdateAgency)              // 更新Agency
	}
	{
		agencyRouterWithoutRecord.GET("findAgency", agencyApi.FindAgency)       // 根据ID获取Agency
		agencyRouterWithoutRecord.GET("getAgencyList", agencyApi.GetAgencyList) // 获取Agency列表
	}
	{
		agencyRouterWithoutRecord.POST("apply", agencyApi.Apply) //申请成为代理
	}
}
