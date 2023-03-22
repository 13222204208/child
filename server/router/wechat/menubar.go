/*
 * @Author: 13222204208 573516293@qq.com
 * @Date: 2023-03-22 08:25:18
 * @LastEditors: 13222204208 573516293@qq.com
 * @LastEditTime: 2023-03-22 14:31:08
 * @FilePath: /server/router/wechat/menubar.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package wechat

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MenubarRouter struct {
}

// InitMenubarRouter 初始化 Menubar 路由信息
func (s *MenubarRouter) InitMenubarRouter(Router *gin.RouterGroup) {
	menubarRouter := Router.Group("menubar").Use(middleware.OperationRecord())
	menubarRouterWithoutRecord := Router.Group("menubar")
	var menubarApi = v1.ApiGroupApp.WechatApiGroup.MenubarApi
	{
		menubarRouter.POST("createMenubar", menubarApi.CreateMenubar)             // 新建Menubar
		menubarRouter.DELETE("deleteMenubar", menubarApi.DeleteMenubar)           // 删除Menubar
		menubarRouter.DELETE("deleteMenubarByIds", menubarApi.DeleteMenubarByIds) // 批量删除Menubar
		menubarRouter.PUT("updateMenubar", menubarApi.UpdateMenubar)              // 更新Menubar
	}
	{
		menubarRouterWithoutRecord.GET("findMenubar", menubarApi.FindMenubar)       // 根据ID获取Menubar
		menubarRouterWithoutRecord.GET("getMenubarList", menubarApi.GetMenubarList) // 获取Menubar列表
		menubarRouterWithoutRecord.GET("list", menubarApi.GetMenubarList)           // 获取Menubar列表
	}
}
