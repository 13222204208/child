package wechat

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ScanRouter struct {
}

// InitScanRouter 初始化 Scan 路由信息
func (s *ScanRouter) InitScanRouter(Router *gin.RouterGroup) {
	scanRouter := Router.Group("scan").Use(middleware.OperationRecord())
	scanRouterWithoutRecord := Router.Group("scan")
	scanPrivateRouter := Router.Group("scan").Use(middleware.JWTAuth())
	var scanApi = v1.ApiGroupApp.WechatApiGroup.ScanApi
	{
		scanRouter.POST("createScan", scanApi.CreateScan)             // 新建Scan
		scanRouter.DELETE("deleteScan", scanApi.DeleteScan)           // 删除Scan
		scanRouter.DELETE("deleteScanByIds", scanApi.DeleteScanByIds) // 批量删除Scan
		scanRouter.PUT("updateScan", scanApi.UpdateScan)              // 更新Scan
	}
	{
		scanRouterWithoutRecord.GET("findScan", scanApi.FindScan)       // 根据ID获取Scan
		scanRouterWithoutRecord.GET("getScanList", scanApi.GetScanList) // 获取Scan列表
	}
	{
		scanPrivateRouter.POST("contrast", scanApi.Contrast) // 扫码对比
	}
}
