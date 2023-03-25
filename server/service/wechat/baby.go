package wechat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	wechatReq "github.com/flipped-aurora/gin-vue-admin/server/model/wechat/request"
)

type BabyService struct {
}

// CreateBaby 创建Baby记录
// Author [piexlmax](https://github.com/piexlmax)
func (babyService *BabyService) CreateBaby(baby wechat.Baby) (err error) {
	err = global.GVA_DB.Create(&baby).Error
	return err
}

// DeleteBaby 删除Baby记录
// Author [piexlmax](https://github.com/piexlmax)
func (babyService *BabyService) DeleteBaby(baby wechat.Baby) (err error) {
	err = global.GVA_DB.Delete(&baby).Error
	return err
}

// DeleteBabyByIds 批量删除Baby记录
// Author [piexlmax](https://github.com/piexlmax)
func (babyService *BabyService) DeleteBabyByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]wechat.Baby{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateBaby 更新Baby记录
// Author [piexlmax](https://github.com/piexlmax)
func (babyService *BabyService) UpdateBaby(baby wechat.Baby) (err error) {
	err = global.GVA_DB.Save(&baby).Error
	return err
}

// GetBaby 根据id获取Baby记录
// Author [piexlmax](https://github.com/piexlmax)
func (babyService *BabyService) GetBaby(id uint) (baby wechat.Baby, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&baby).Error
	return
}

// GetBabyInfoList 分页获取Baby记录
// Author [piexlmax](https://github.com/piexlmax)
func (babyService *BabyService) GetBabyInfoList(info wechatReq.BabySearch) (list []wechat.Baby, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&wechat.Baby{})
	var babys []wechat.Baby
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&babys).Error
	return babys, total, err
}

// 保存baby信息
func (babyService *BabyService) SaveBabyInfo(baby wechat.Baby) (babyId uint, err error) {
	//根据baby的uid和name查询是否存在
	var babyInfo wechat.Baby
	err = global.GVA_DB.Where("uid = ? and name = ?", baby.Uid, baby.Name).First(&babyInfo).Error
	if err != nil {
		//不存在则创建
		err = global.GVA_DB.Save(&baby).Error
	}
	return
}
