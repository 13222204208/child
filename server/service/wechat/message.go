/*
 * @Author: 13222204208 573516293@qq.com
 * @Date: 2023-03-22 14:42:43
 * @LastEditors: 13222204208 573516293@qq.com
 * @LastEditTime: 2023-03-22 15:26:19
 * @FilePath: /server/service/wechat/message.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package wechat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	wechatReq "github.com/flipped-aurora/gin-vue-admin/server/model/wechat/request"
)

type MessageService struct {
}

// CreateMessage 创建Message记录
// Author [piexlmax](https://github.com/piexlmax)
func (messageService *MessageService) CreateMessage(message wechat.Message) (err error) {
	err = global.GVA_DB.Create(&message).Error
	return err
}

// DeleteMessage 删除Message记录
// Author [piexlmax](https://github.com/piexlmax)
func (messageService *MessageService) DeleteMessage(message wechat.Message) (err error) {
	err = global.GVA_DB.Delete(&message).Error
	return err
}

// DeleteMessageByIds 批量删除Message记录
// Author [piexlmax](https://github.com/piexlmax)
func (messageService *MessageService) DeleteMessageByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]wechat.Message{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateMessage 更新Message记录
// Author [piexlmax](https://github.com/piexlmax)
func (messageService *MessageService) UpdateMessage(message wechat.Message) (err error) {
	err = global.GVA_DB.Save(&message).Error
	return err
}

// GetMessage 根据id获取Message记录
// Author [piexlmax](https://github.com/piexlmax)
func (messageService *MessageService) GetMessage(id uint) (message wechat.Message, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&message).Error
	return
}

// GetMessageInfoList 分页获取Message记录
// Author [piexlmax](https://github.com/piexlmax)
func (messageService *MessageService) GetMessageInfoList(info wechatReq.MessageSearch) (list []wechat.Message, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&wechat.Message{})
	var messages []wechat.Message
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&messages).Error
	return messages, total, err
}

// get message list by category
func (messageService *MessageService) GetMessageListByCategory(category int) (info wechat.Message, err error) {
	err = global.GVA_DB.Where("category = ?", category).First(&info).Error
	return
}
