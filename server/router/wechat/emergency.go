package wechat

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EmergencyAlertRouter struct {
}

// InitEmergencyAlertRouter 初始化 EmergencyAlert 路由信息
func (s *EmergencyAlertRouter) InitEmergencyAlertRouter(Router *gin.RouterGroup) {
	emergencyAlertRouter := Router.Group("emergencyAlert").Use(middleware.OperationRecord())
	emergencyAlertRouterWithoutRecord := Router.Group("emergencyAlert")

	emergencyAlertPrivateRouter := Router.Group("emergency").Use(middleware.JWTAuthMiddleware())

	var emergencyAlertApi = v1.ApiGroupApp.WechatApiGroup.EmergencyAlertApi
	{
		emergencyAlertRouter.POST("createEmergencyAlert", emergencyAlertApi.CreateEmergencyAlert)             // 新建EmergencyAlert
		emergencyAlertRouter.DELETE("deleteEmergencyAlert", emergencyAlertApi.DeleteEmergencyAlert)           // 删除EmergencyAlert
		emergencyAlertRouter.DELETE("deleteEmergencyAlertByIds", emergencyAlertApi.DeleteEmergencyAlertByIds) // 批量删除EmergencyAlert
		emergencyAlertRouter.PUT("updateEmergencyAlert", emergencyAlertApi.UpdateEmergencyAlert)              // 更新EmergencyAlert
	}
	{
		emergencyAlertRouterWithoutRecord.GET("findEmergencyAlert", emergencyAlertApi.FindEmergencyAlert)       // 根据ID获取EmergencyAlert
		emergencyAlertRouterWithoutRecord.GET("getEmergencyAlertList", emergencyAlertApi.GetEmergencyAlertList) // 获取EmergencyAlert列表
	}

	{
		//发布预警通知
		emergencyAlertPrivateRouter.POST("issue", emergencyAlertApi.Issue)
		//获取预警通知列表
		emergencyAlertPrivateRouter.GET("list", emergencyAlertApi.List)

		//根据id获取预警通知
		emergencyAlertPrivateRouter.GET(":id", emergencyAlertApi.FindEmergencyAlertById)
	}
}
