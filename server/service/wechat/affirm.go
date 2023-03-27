package wechat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	wechatReq "github.com/flipped-aurora/gin-vue-admin/server/model/wechat/request"
)

type AffirmService struct {
}

// CreateAffirm 创建Affirm记录
// Author [piexlmax](https://github.com/piexlmax)
func (affirmService *AffirmService) CreateAffirm(affirm wechat.Affirm) (err error) {
	err = global.GVA_DB.Create(&affirm).Error
	return err
}

// DeleteAffirm 删除Affirm记录
// Author [piexlmax](https://github.com/piexlmax)
func (affirmService *AffirmService) DeleteAffirm(affirm wechat.Affirm) (err error) {
	err = global.GVA_DB.Delete(&affirm).Error
	return err
}

// DeleteAffirmByIds 批量删除Affirm记录
// Author [piexlmax](https://github.com/piexlmax)
func (affirmService *AffirmService) DeleteAffirmByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]wechat.Affirm{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAffirm 更新Affirm记录
// Author [piexlmax](https://github.com/piexlmax)
func (affirmService *AffirmService) UpdateAffirm(affirm wechat.Affirm) (err error) {
	err = global.GVA_DB.Save(&affirm).Error
	return err
}

// GetAffirm 根据id获取Affirm记录
// Author [piexlmax](https://github.com/piexlmax)
func (affirmService *AffirmService) GetAffirm(id uint) (affirm wechat.Affirm, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&affirm).Error
	return
}

// GetAffirmInfoList 分页获取Affirm记录
// Author [piexlmax](https://github.com/piexlmax)
func (affirmService *AffirmService) GetAffirmInfoList(info wechatReq.AffirmSearch) (list []wechat.Affirm, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&wechat.Affirm{})
	var affirms []wechat.Affirm
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["sort"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	err = db.Limit(limit).Offset(offset).Find(&affirms).Error
	return affirms, total, err
}

// 查询status==1的数据
func (affirmService *AffirmService) GetAffirmList() (list []wechat.Affirm, err error) {

	db := global.GVA_DB.Model(&wechat.Affirm{})

	err = db.Where("status = ?", 1).Find(&list).Error
	return
}
