package wechat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	wechatReq "github.com/flipped-aurora/gin-vue-admin/server/model/wechat/request"
)

type IdentityService struct {
}

// CreateIdentity 创建Identity记录
// Author [piexlmax](https://github.com/piexlmax)
func (identityService *IdentityService) CreateIdentity(identity wechat.Identity) (err error) {
	err = global.GVA_DB.Create(&identity).Error
	return err
}

// DeleteIdentity 删除Identity记录
// Author [piexlmax](https://github.com/piexlmax)
func (identityService *IdentityService) DeleteIdentity(identity wechat.Identity) (err error) {
	err = global.GVA_DB.Delete(&identity).Error
	return err
}

// DeleteIdentityByIds 批量删除Identity记录
// Author [piexlmax](https://github.com/piexlmax)
func (identityService *IdentityService) DeleteIdentityByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]wechat.Identity{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateIdentity 更新Identity记录
// Author [piexlmax](https://github.com/piexlmax)
func (identityService *IdentityService) UpdateIdentity(identity wechat.Identity) (err error) {
	err = global.GVA_DB.Save(&identity).Error
	return err
}

// GetIdentity 根据id获取Identity记录
// Author [piexlmax](https://github.com/piexlmax)
func (identityService *IdentityService) GetIdentity(id uint) (identity wechat.Identity, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&identity).Error
	return
}

// GetIdentityInfoList 分页获取Identity记录
// Author [piexlmax](https://github.com/piexlmax)
func (identityService *IdentityService) GetIdentityInfoList(info wechatReq.IdentitySearch) (list []wechat.Identity, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&wechat.Identity{})
	var identitys []wechat.Identity
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&identitys).Error
	return identitys, total, err
}

// get identity list status 1
func (identityService *IdentityService) GetIdentityList() (list []wechat.Identity, err error) {
	err = global.GVA_DB.Where("status = ?", 1).Find(&list).Error
	return
}
