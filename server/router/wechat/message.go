/*
 * @Author: 13222204208 573516293@qq.com
 * @Date: 2023-03-22 14:42:43
 * @LastEditors: 13222204208 573516293@qq.com
 * @LastEditTime: 2023-03-22 15:28:32
 * @FilePath: /server/router/wechat/message.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package wechat

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MessageRouter struct {
}

// InitMessageRouter 初始化 Message 路由信息
func (s *MessageRouter) InitMessageRouter(Router *gin.RouterGroup) {
	messageRouter := Router.Group("message").Use(middleware.OperationRecord())
	messageRouterWithoutRecord := Router.Group("message")
	var messageApi = v1.ApiGroupApp.WechatApiGroup.MessageApi
	{
		messageRouter.POST("createMessage", messageApi.CreateMessage)             // 新建Message
		messageRouter.DELETE("deleteMessage", messageApi.DeleteMessage)           // 删除Message
		messageRouter.DELETE("deleteMessageByIds", messageApi.DeleteMessageByIds) // 批量删除Message
		messageRouter.PUT("updateMessage", messageApi.UpdateMessage)              // 更新Message
	}
	{
		messageRouterWithoutRecord.GET("findMessage", messageApi.FindMessage)       // 根据ID获取Message
		messageRouterWithoutRecord.GET("getMessageList", messageApi.GetMessageList) // 获取Message列表

		messageRouterWithoutRecord.GET("info", messageApi.GetMessageByCategory)
	}
}
