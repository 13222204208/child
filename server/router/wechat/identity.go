package wechat

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type IdentityRouter struct {
}

// InitIdentityRouter 初始化 Identity 路由信息
func (s *IdentityRouter) InitIdentityRouter(Router *gin.RouterGroup) {
	identityRouter := Router.Group("identity").Use(middleware.OperationRecord())
	identityRouterWithoutRecord := Router.Group("identity")
	var identityApi = v1.ApiGroupApp.WechatApiGroup.IdentityApi
	{
		identityRouter.POST("createIdentity", identityApi.CreateIdentity)             // 新建Identity
		identityRouter.DELETE("deleteIdentity", identityApi.DeleteIdentity)           // 删除Identity
		identityRouter.DELETE("deleteIdentityByIds", identityApi.DeleteIdentityByIds) // 批量删除Identity
		identityRouter.PUT("updateIdentity", identityApi.UpdateIdentity)              // 更新Identity
	}
	{
		identityRouterWithoutRecord.GET("findIdentity", identityApi.FindIdentity)       // 根据ID获取Identity
		identityRouterWithoutRecord.GET("getIdentityList", identityApi.GetIdentityList) // 获取Identity列表

		identityRouterWithoutRecord.GET("list", identityApi.List) // 获取Identity列表
	}
}
