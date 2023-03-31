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
	Pic     string `json:"pic" form:"pic" binding:"required"`
	Address string `json:"address" form:"address" binding:"required"`
}
