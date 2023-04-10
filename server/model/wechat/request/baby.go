package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
)

type BabySearch struct {
	wechat.Baby
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type SendCode struct {
	Phone string `json:"phone" form:"phone" binding:"required"`
}

// 验证手机号是否存在
type PhoneExist struct {
	Phone string `json:"phone" form:"phone" binding:"required"`
	Code  string `json:"code" form:"code" binding:"required"`
}
