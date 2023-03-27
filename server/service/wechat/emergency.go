package wechat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	wechatReq "github.com/flipped-aurora/gin-vue-admin/server/model/wechat/request"
)

type EmergencyAlertService struct {
}

// CreateEmergencyAlert 创建EmergencyAlert记录
// Author [piexlmax](https://github.com/piexlmax)
func (emergencyAlertService *EmergencyAlertService) CreateEmergencyAlert(emergencyAlert wechat.EmergencyAlert) (err error) {
	err = global.GVA_DB.Create(&emergencyAlert).Error
	return err
}

// DeleteEmergencyAlert 删除EmergencyAlert记录
// Author [piexlmax](https://github.com/piexlmax)
func (emergencyAlertService *EmergencyAlertService) DeleteEmergencyAlert(emergencyAlert wechat.EmergencyAlert) (err error) {
	err = global.GVA_DB.Delete(&emergencyAlert).Error
	return err
}

// DeleteEmergencyAlertByIds 批量删除EmergencyAlert记录
// Author [piexlmax](https://github.com/piexlmax)
func (emergencyAlertService *EmergencyAlertService) DeleteEmergencyAlertByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]wechat.EmergencyAlert{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateEmergencyAlert 更新EmergencyAlert记录
// Author [piexlmax](https://github.com/piexlmax)
func (emergencyAlertService *EmergencyAlertService) UpdateEmergencyAlert(emergencyAlert wechat.EmergencyAlert) (err error) {
	err = global.GVA_DB.Save(&emergencyAlert).Error
	return err
}

// GetEmergencyAlert 根据id获取EmergencyAlert记录
// Author [piexlmax](https://github.com/piexlmax)
func (emergencyAlertService *EmergencyAlertService) GetEmergencyAlert(id uint) (emergencyAlert wechat.EmergencyAlert, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&emergencyAlert).Error
	return
}

// GetEmergencyAlertInfoList 分页获取EmergencyAlert记录
// Author [piexlmax](https://github.com/piexlmax)
func (emergencyAlertService *EmergencyAlertService) GetEmergencyAlertInfoList(info wechatReq.EmergencyAlertSearch) (list []wechat.EmergencyAlert, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&wechat.EmergencyAlert{})
	var emergencyAlerts []wechat.EmergencyAlert
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&emergencyAlerts).Error
	return emergencyAlerts, total, err
}

// save a emergencyAlert record
func (emergencyAlertService *EmergencyAlertService) Issue(emergencyAlert wechat.EmergencyAlert) (err error) {
	err = global.GVA_DB.Save(&emergencyAlert).Error
	return err
}
