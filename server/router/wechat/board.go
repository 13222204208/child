/*
 * @Author: 13222204208 573516293@qq.com
 * @Date: 2023-03-22 14:57:03
 * @LastEditors: 13222204208 573516293@qq.com
 * @LastEditTime: 2023-03-22 15:19:30
 * @FilePath: /server/router/wechat/board.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package wechat

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BoardRouter struct {
}

// InitBoardRouter 初始化 Board 路由信息
func (s *BoardRouter) InitBoardRouter(Router *gin.RouterGroup) {
	boardRouter := Router.Group("board").Use(middleware.OperationRecord())
	boardRouterWithoutRecord := Router.Group("board")
	var boardApi = v1.ApiGroupApp.WechatApiGroup.BoardApi
	{
		boardRouter.POST("createBoard", boardApi.CreateBoard)             // 新建Board
		boardRouter.DELETE("deleteBoard", boardApi.DeleteBoard)           // 删除Board
		boardRouter.DELETE("deleteBoardByIds", boardApi.DeleteBoardByIds) // 批量删除Board
		boardRouter.PUT("updateBoard", boardApi.UpdateBoard)              // 更新Board
	}
	{
		boardRouterWithoutRecord.GET("findBoard", boardApi.FindBoard)       // 根据ID获取Board
		boardRouterWithoutRecord.GET("getBoardList", boardApi.GetBoardList) // 获取Board列表

		boardRouterWithoutRecord.GET("list", boardApi.GetList) // 获取Board列表
	}
}
