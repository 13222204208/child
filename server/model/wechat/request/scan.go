package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
)

type ScanSearch struct {
	wechat.Scan
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

// 扫描对比
type ScanCompare struct {
	Pic string `json:"pic" form:"pic" binding:"required"`
	//经度
	Lng float64 `json:"lng" form:"lng" binding:"required"`
	//纬度
	Lat float64 `json:"lat" form:"lat" binding:"required"`
}
