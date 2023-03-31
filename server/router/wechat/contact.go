package wechat

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ContactRouter struct {
}

// InitContactRouter 初始化 Contact 路由信息
func (s *ContactRouter) InitContactRouter(Router *gin.RouterGroup) {
	contactRouter := Router.Group("contact").Use(middleware.OperationRecord())
	contactRouterWithoutRecord := Router.Group("contact")

	contactPrivateRouter := Router.Group("contact").Use(middleware.JWTAuthMiddleware())

	var contactApi = v1.ApiGroupApp.WechatApiGroup.ContactApi
	{
		contactRouter.POST("createContact", contactApi.CreateContact)             // 新建Contact
		contactRouter.DELETE("deleteContact", contactApi.DeleteContact)           // 删除Contact
		contactRouter.DELETE("deleteContactByIds", contactApi.DeleteContactByIds) // 批量删除Contact
		contactRouter.PUT("updateContact", contactApi.UpdateContact)              // 更新Contact
	}
	{
		contactRouterWithoutRecord.GET("findContact", contactApi.FindContact)       // 根据ID获取Contact
		contactRouterWithoutRecord.GET("getContactList", contactApi.GetContactList) // 获取Contact列表
	}
	{
		contactPrivateRouter.POST("save", contactApi.SaveContact) // 保存Contact
	}
}
