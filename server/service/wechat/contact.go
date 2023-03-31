package wechat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	wechatReq "github.com/flipped-aurora/gin-vue-admin/server/model/wechat/request"
)

type ContactService struct {
}

// CreateContact 创建Contact记录
// Author [piexlmax](https://github.com/piexlmax)
func (contactService *ContactService) CreateContact(contact wechat.Contact) (err error) {
	err = global.GVA_DB.Create(&contact).Error
	return err
}

// DeleteContact 删除Contact记录
// Author [piexlmax](https://github.com/piexlmax)
func (contactService *ContactService) DeleteContact(contact wechat.Contact) (err error) {
	err = global.GVA_DB.Delete(&contact).Error
	return err
}

// DeleteContactByIds 批量删除Contact记录
// Author [piexlmax](https://github.com/piexlmax)
func (contactService *ContactService) DeleteContactByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]wechat.Contact{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateContact 更新Contact记录
// Author [piexlmax](https://github.com/piexlmax)
func (contactService *ContactService) UpdateContact(contact wechat.Contact) (err error) {
	err = global.GVA_DB.Save(&contact).Error
	return err
}

// GetContact 根据id获取Contact记录
// Author [piexlmax](https://github.com/piexlmax)
func (contactService *ContactService) GetContact(id uint) (contact wechat.Contact, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&contact).Error
	return
}

// GetContactInfoList 分页获取Contact记录
// Author [piexlmax](https://github.com/piexlmax)
func (contactService *ContactService) GetContactInfoList(info wechatReq.ContactSearch) (list []wechat.Contact, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&wechat.Contact{})
	var contacts []wechat.Contact
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&contacts).Error
	return contacts, total, err
}

// save contact
func (contactService *ContactService) SaveContact(contact wechat.Contact) (err error) {
	err = global.GVA_DB.Save(&contact).Error
	return err
}
