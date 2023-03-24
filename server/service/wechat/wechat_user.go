package wechat

import (
	"errors"
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	wechatReq "github.com/flipped-aurora/gin-vue-admin/server/model/wechat/request"
	"github.com/golang-jwt/jwt/v4"
	w "github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram/config"
)

type WechatUserService struct {
}

// CreateWechatUser 创建WechatUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (wechatUserService *WechatUserService) CreateWechatUser(wechatUser wechat.WechatUser) (err error) {
	err = global.GVA_DB.Create(&wechatUser).Error
	return err
}

// DeleteWechatUser 删除WechatUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (wechatUserService *WechatUserService) DeleteWechatUser(wechatUser wechat.WechatUser) (err error) {
	err = global.GVA_DB.Delete(&wechatUser).Error
	return err
}

// DeleteWechatUserByIds 批量删除WechatUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (wechatUserService *WechatUserService) DeleteWechatUserByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]wechat.WechatUser{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateWechatUser 更新WechatUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (wechatUserService *WechatUserService) UpdateWechatUser(wechatUser wechat.WechatUser) (err error) {
	err = global.GVA_DB.Save(&wechatUser).Error
	return err
}

// GetWechatUser 根据id获取WechatUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (wechatUserService *WechatUserService) GetWechatUser(id uint) (wechatUser wechat.WechatUser, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&wechatUser).Error
	return
}

// GetWechatUserInfoList 分页获取WechatUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (wechatUserService *WechatUserService) GetWechatUserInfoList(info wechatReq.WechatUserSearch) (list []wechat.WechatUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&wechat.WechatUser{})
	var wechatUsers []wechat.WechatUser
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&wechatUsers).Error
	return wechatUsers, total, err
}

func (wechatUserService *WechatUserService) Login(code string) (token string, err error) {
	fmt.Println("code:", code)
	openid, err := GetWechatOpenid(code)
	if err != nil {
		return
	}
	if openid != "" {
		var wechatUser wechat.WechatUser
		err = global.GVA_DB.Where("openid = ?", openid).First(&wechatUser).Error
		if err != nil {
			return
		}
		uid := wechatUser.ID
		if uid > 0 {
			token, err = genToken(uid)
			if err != nil {
				return
			}
		} else {
			wechatUser.Openid = openid
			err = global.GVA_DB.Create(&wechatUser).Error
			if err != nil {
				return
			}
			token, err = genToken(wechatUser.ID)
			if err != nil {
				return
			}
		}
	}
	return
}

// GenToken 生成JWT
func genToken(id uint) (string, error) {
	// 创建一个我们自己的声明
	c := wechatReq.MyClaims{
		ID: id, // 自定义字段
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(wechatReq.TokenExpireDuration)), // 过期时间
			Issuer:    "yangpanda",                                                       // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(wechatReq.MySecret)
}

// 获取微信小程序的openid
func GetWechatOpenid(code string) (s string, err error) {
	appid := global.GVA_CONFIG.Wechat.Appid
	secret := global.GVA_CONFIG.Wechat.Secret
	if appid == "" || secret == "" {
		return s, errors.New("appid或secret为空")
	}
	wc := w.NewWechat()
	memory := cache.NewMemory()
	cfg := &config.Config{
		AppID:     appid,
		AppSecret: secret,
		Cache:     memory,
	}
	min := wc.GetMiniProgram(cfg)
	a := min.GetAuth()
	r, errs := a.Code2Session(code)

	if errs != nil {
		return s, errors.New(errs.Error())
	} else {
		return r.OpenID, nil
	}
}
