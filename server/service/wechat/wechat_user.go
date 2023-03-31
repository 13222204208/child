package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	assist "github.com/13222204208/assist/wechat"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	wechatReq "github.com/flipped-aurora/gin-vue-admin/server/model/wechat/request"
	"github.com/golang-jwt/jwt/v4"
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

func (wechatUserService *WechatUserService) Login(userInfo wechatReq.WechatUserInfo) (token string, err error) {
	code := userInfo.Code
	fmt.Println("code:", code)
	openid, accessToken, err := GetWechatOpenid(code)
	if err != nil {
		return
	}

	url := "https://api.weixin.qq.com/sns/userinfo?access_token=" + accessToken + "&openid=" + openid + "&lang=zh_CN"
	res, err := GetUserInfo(url)
	if err != nil {
		return
	}
	//res解析到userinfo
	var info userinfo
	err = json.Unmarshal(res, &info)
	if err != nil {
		return
	}
	fmt.Println("openid:", openid)
	if err != nil {
		return
	}
	if openid != "" {
		var wechatUser wechat.WechatUser
		global.GVA_DB.Where("openid = ?", openid).First(&wechatUser)

		uid := wechatUser.ID
		if uid > 0 {
			token, err = genToken(uid)
			if err != nil {
				return
			} else {
				return token, err
			}
		} else {
			wechatUser.Openid = openid
			wechatUser.Nickname = info.Nickname
			wechatUser.Avatar = info.Headimgurl
			wechatUser.Name = userInfo.Name
			wechatUser.Phone = userInfo.Phone
			wechatUser.City = userInfo.City
			wechatUser.Address = userInfo.Address
			fmt.Println("Openid:", openid)
			err = global.GVA_DB.Create(&wechatUser).Error
			if err != nil {
				return
			}
			fmt.Println("wechatUser.ID:", wechatUser.ID)
			token, err = genToken(wechatUser.ID)
			fmt.Println("token:", token)
			if err != nil {
				return
			}
		}
	} else {
		err = errors.New("openid 为空")
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

// 获取微信公众号的openid
func GetWechatOpenid(code string) (s, a string, err error) {
	appid := global.GVA_CONFIG.Wechat.Appid
	secret := global.GVA_CONFIG.Wechat.Secret
	if appid == "" || secret == "" {
		return s, a, errors.New("appid或secret为空")
	}
	openid, accessToken, err := assist.GetWechatH5OpenidAndAccessToken(appid, secret, code)
	fmt.Println("openid:", openid, "err:", err)
	if err != nil {
		return s, a, errors.New(err.Error())
	} else {
		return openid, accessToken, nil
	}
}

type userinfo struct {
	Openid     string   `json:"openid"`
	Nickname   string   `json:"nickname"`
	Sex        int      `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	Headimgurl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	Unionid    string   `json:"unionid"`
}

func GetUserInfo(url string) (r []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	r, err = io.ReadAll(resp.Body)
	// r = string(body)
	// fmt.Println("返回的结果", r)
	return
}
