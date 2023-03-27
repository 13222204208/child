package wechat

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	wechatReq "github.com/flipped-aurora/gin-vue-admin/server/model/wechat/request"
)

type AgencyService struct {
}

// CreateAgency 创建Agency记录
// Author [piexlmax](https://github.com/piexlmax)
func (agencyService *AgencyService) CreateAgency(agency wechat.Agency) (err error) {
	err = global.GVA_DB.Create(&agency).Error
	return err
}

// DeleteAgency 删除Agency记录
// Author [piexlmax](https://github.com/piexlmax)
func (agencyService *AgencyService) DeleteAgency(agency wechat.Agency) (err error) {
	err = global.GVA_DB.Delete(&agency).Error
	return err
}

// DeleteAgencyByIds 批量删除Agency记录
// Author [piexlmax](https://github.com/piexlmax)
func (agencyService *AgencyService) DeleteAgencyByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]wechat.Agency{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAgency 更新Agency记录
// Author [piexlmax](https://github.com/piexlmax)
func (agencyService *AgencyService) UpdateAgency(agency wechat.Agency) (err error) {
	err = global.GVA_DB.Save(&agency).Error
	return err
}

// GetAgency 根据id获取Agency记录
// Author [piexlmax](https://github.com/piexlmax)
func (agencyService *AgencyService) GetAgency(id uint) (agency wechat.Agency, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&agency).Error
	return
}

// GetAgencyInfoList 分页获取Agency记录
// Author [piexlmax](https://github.com/piexlmax)
func (agencyService *AgencyService) GetAgencyInfoList(info wechatReq.AgencySearch) (list []wechat.Agency, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&wechat.Agency{})
	var agencys []wechat.Agency
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&agencys).Error
	return agencys, total, err
}

func (agencyService *AgencyService) Apply(agency wechat.Agency) (err error) {
	//根据phoen查询是否已经申请过
	var agency1 wechat.Agency
	err = global.GVA_DB.Where("phone = ?", agency.Phone).First(&agency1).Error
	if err == nil {
		return errors.New("该手机号已经申请过")
	}
	err = global.GVA_DB.Create(&agency).Error
	return err
}
