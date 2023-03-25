package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
)

type WechatUserSearch struct {
	wechat.WechatUser
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type WechatCode struct {
	Code string `json:"code" form:"code"`
}

type WechatUserInfo struct {
	Name    string `json:"name" form:"name"`
	Phone   string `json:"phone" form:"phone"`
	Address string `json:"address" form:"address"`
	City    string `json:"city" form:"city"`
}
