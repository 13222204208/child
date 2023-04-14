package wechat

import (
	"errors"
	"strconv"

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
	if err == nil {
		//发送信息给宝贝的家长
		err = SendMsgToParent(contact)
	}
	return err
}

// 发送信息给宝贝的家长
func SendMsgToParent(contact wechat.Contact) (err error) {

	appid := global.GVA_CONFIG.Wechat.Appid
	secret := global.GVA_CONFIG.Wechat.Secret
	if appid == "" || secret == "" {
		return errors.New("appid或secret为空")
	}
	accessToken, err := GetWechatAccessToken(appid, secret)
	if err != nil {
		return errors.New(err.Error())
	}

	//获取宝贝信息
	babyId := contact.BabyId
	babyService := BabyService{}
	baby, err := babyService.GetBaby(babyId)
	if err != nil {
		return err
	}
	//获取宝贝的家长信息
	parentService := WechatUserService{}
	parents, err := parentService.GetWechatUser(baby.Uid)
	if err != nil {
		return err
	}
	var t TemplateList
	t.Appid = appid
	t.Token = accessToken
	t.Secret = secret
	t.Openid = parents.Openid
	t.Phone = contact.Phone
	t.Remark = contact.Message
	//babyId 转为字符串
	t.BabyId = strconv.Itoa(int(babyId))
	t.Url = "https://huzhu.cnecip.com/wechat/#/pages/zhoubian/saomiao"
	//当前时间
	t.LostTime = contact.CreatedAt.Format("2006-01-02 15:04:05")
	err = SendTextMessageToUser(&t)
	return err
}
