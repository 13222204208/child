package wechat

import (
	"encoding/json"
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	wechatReq "github.com/flipped-aurora/gin-vue-admin/server/model/wechat/request"
)

type ScanService struct {
}

// CreateScan 创建Scan记录
// Author [piexlmax](https://github.com/piexlmax)
func (scanService *ScanService) CreateScan(scan wechat.Scan) (err error) {
	err = global.GVA_DB.Create(&scan).Error
	return err
}

// DeleteScan 删除Scan记录
// Author [piexlmax](https://github.com/piexlmax)
func (scanService *ScanService) DeleteScan(scan wechat.Scan) (err error) {
	err = global.GVA_DB.Delete(&scan).Error
	return err
}

// DeleteScanByIds 批量删除Scan记录
// Author [piexlmax](https://github.com/piexlmax)
func (scanService *ScanService) DeleteScanByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]wechat.Scan{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateScan 更新Scan记录
// Author [piexlmax](https://github.com/piexlmax)
func (scanService *ScanService) UpdateScan(scan wechat.Scan) (err error) {
	err = global.GVA_DB.Save(&scan).Error
	return err
}

// GetScan 根据id获取Scan记录
// Author [piexlmax](https://github.com/piexlmax)
func (scanService *ScanService) GetScan(id uint) (scan wechat.Scan, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&scan).Error
	return
}

// GetScanInfoList 分页获取Scan记录
// Author [piexlmax](https://github.com/piexlmax)
func (scanService *ScanService) GetScanInfoList(info wechatReq.ScanSearch) (list []wechat.Scan, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&wechat.Scan{})
	var scans []wechat.Scan
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&scans).Error
	return scans, total, err
}

// Contrast 对比 两张图片
func (scanService *ScanService) Contrast(pic, address string, uid uint) (info wechat.Baby, err error) {
	res, err := GetEmergencyBaby()
	if err != nil {
		return
	}
	for _, v := range res {
		img, err := ParseJsonArray(v.Feature)
		if err != nil {
			return info, err
		}
		// 对比图片
		if img == pic {
			info = v
			err = SaveScan(uid, v.ID, address)
			if err != nil {
				return info, err
			}
		} else {
			return info, errors.New("无法查询到相应的宝宝信息")
		}
	}
	return
}

// 查询出所有发布预警的信息
func GetEmergencyBaby() (baby []wechat.Baby, err error) {
	//查询出所有EmergencyAlert 中的 baby_id 并去重
	var emergencyAlert []wechat.EmergencyAlert
	err = global.GVA_DB.Model(&wechat.EmergencyAlert{}).Select("baby_id").Group("baby_id").Find(&emergencyAlert).Error
	if err != nil {
		return
	}
	//根据 baby_id 查询出所有的baby
	err = global.GVA_DB.Model(&wechat.Baby{}).Where("id in ?", emergencyAlert).Find(&baby).Error

	return
}

// 生物特征
type feature struct {
	Img  string `json:"img"`
	Time string `json:"time"`
}

// 解析一个json 数组字符串
func ParseJsonArray(jsonStr string) (img string, err error) {
	var res []feature
	err = json.Unmarshal([]byte(jsonStr), &res)
	img = res[0].Img
	return
}

// save scan info
func SaveScan(uid, babyId uint, address string) (err error) {
	var scan wechat.Scan
	*scan.Uid = uid
	*scan.BabyId = babyId
	scan.Address = address
	err = global.GVA_DB.Create(&scan).Error
	return err
}
