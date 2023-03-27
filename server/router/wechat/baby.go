package wechat

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BabyRouter struct {
}

// InitBabyRouter 初始化 Baby 路由信息
func (s *BabyRouter) InitBabyRouter(Router *gin.RouterGroup, privateRouter *gin.RouterGroup) {
	babyRouter := Router.Group("baby").Use(middleware.OperationRecord())
	babyRouterWithoutRecord := Router.Group("baby")

	babyPrivateRouter := privateRouter.Group("baby").Use(middleware.JWTAuthMiddleware())

	var babyApi = v1.ApiGroupApp.WechatApiGroup.BabyApi
	{
		babyRouter.POST("createBaby", babyApi.CreateBaby)             // 新建Baby
		babyRouter.DELETE("deleteBaby", babyApi.DeleteBaby)           // 删除Baby
		babyRouter.DELETE("deleteBabyByIds", babyApi.DeleteBabyByIds) // 批量删除Baby
		babyRouter.PUT("updateBaby", babyApi.UpdateBaby)              // 更新Baby
	}
	{
		babyRouterWithoutRecord.GET("findBaby", babyApi.FindBaby)       // 根据ID获取Baby
		babyRouterWithoutRecord.GET("getBabyList", babyApi.GetBabyList) // 获取Baby列表
	}
	{
		babyPrivateRouter.POST("save", babyApi.SaveBaby) // 保存Baby
		babyPrivateRouter.GET(":id", babyApi.Info)
	}
}
