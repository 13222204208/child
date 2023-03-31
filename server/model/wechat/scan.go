// 自动生成模板Scan
package wechat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Scan 结构体
type Scan struct {
	global.GVA_MODEL
	Uid     *uint  `json:"uid" form:"uid" gorm:"column:uid;comment:扫描人ID;"`
	Source  string `json:"source" form:"source" gorm:"column:source;comment:扫描来源默认为微信;default:微信;"`
	Method  string `json:"method" form:"method" gorm:"column:method;comment:默认为拍照识别;default:拍照识别;"`
	Address string `json:"address" form:"address" gorm:"column:address;comment:;"`
	BabyId  *uint  `json:"babyId" form:"babyId" gorm:"column:baby_id;comment:被扫描的宝贝ID;"`
	Pic     string `json:"pic" form:"pic" gorm:"column:pic;comment:扫描的图片;"`
	Status  string `json:"status" form:"status" gorm:"column:status;comment:扫描状态;"`
}

// TableName Scan 表名
func (Scan) TableName() string {
	return "wechat_scan_record"
}
