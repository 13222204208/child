package wechat

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/13222204208/tool"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	wechatReq "github.com/flipped-aurora/gin-vue-admin/server/model/wechat/request"
)

type BabyService struct {
}

// CreateBaby 创建Baby记录
// Author [piexlmax](https://github.com/piexlmax)
func (babyService *BabyService) CreateBaby(baby wechat.Baby) (err error) {
	err = global.GVA_DB.Create(&baby).Error
	return err
}

// DeleteBaby 删除Baby记录
// Author [piexlmax](https://github.com/piexlmax)
func (babyService *BabyService) DeleteBaby(baby wechat.Baby) (err error) {
	err = global.GVA_DB.Delete(&baby).Error
	return err
}

// DeleteBabyByIds 批量删除Baby记录
// Author [piexlmax](https://github.com/piexlmax)
func (babyService *BabyService) DeleteBabyByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]wechat.Baby{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateBaby 更新Baby记录
// Author [piexlmax](https://github.com/piexlmax)
func (babyService *BabyService) UpdateBaby(baby wechat.Baby) (err error) {
	err = global.GVA_DB.Save(&baby).Error
	return err
}

// GetBaby 根据id获取Baby记录
// Author [piexlmax](https://github.com/piexlmax)
func (babyService *BabyService) GetBaby(id uint) (baby wechat.Baby, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&baby).Error
	return
}

// GetBabyInfoList 分页获取Baby记录
// Author [piexlmax](https://github.com/piexlmax)
func (babyService *BabyService) GetBabyInfoList(info wechatReq.BabySearch) (list []wechat.Baby, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&wechat.Baby{})
	var babys []wechat.Baby
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&babys).Error
	return babys, total, err
}

// 保存baby信息
func (babyService *BabyService) SaveBabyInfo(baby wechat.Baby) (babyId uint, err error) {
	//根据baby的uid和name查询是否存在
	if baby.ID != 0 {
		// err = global.GVA_DB.Save(&baby).Error
		err = global.GVA_DB.Model(&baby).Updates(baby).Error
		babyId = baby.ID
		return
	}

	var babyInfo wechat.Baby
	err = global.GVA_DB.Where("uid = ? and name = ?", baby.Uid, baby.Name).First(&babyInfo).Error
	babyId = babyInfo.ID
	fmt.Println("babyId", babyId)
	if err != nil {
		//不存在则创建
		err = global.GVA_DB.Save(&baby).Error
		babyId = baby.ID
		fmt.Println("babyId2", babyId)
	}
	return
}

// 给手机号发送验证码
func (babyService *BabyService) SendCode(phone string) (err error) {
	err = SendSms(phone)
	if err != nil {
		return errors.New(err.Error())
	}
	return err
}

var ctx = context.Background()

// 发送验证码
func SendSms(phoneNumber string) error {
	config := sdk.NewConfig()

	phoneKeyId := global.GVA_CONFIG.Phone.PhoneKeyId
	phoneKeySecret := global.GVA_CONFIG.Phone.PhoneKeySecret
	phoneSignName := global.GVA_CONFIG.Phone.PhoneSignName
	phoneTemplateCode := global.GVA_CONFIG.Phone.PhoneTemplateCode
	if phoneKeyId == "" || phoneKeySecret == "" {
		return errors.New("请先配置阿里云短信服务")
	}

	credential := credentials.NewAccessKeyCredential(phoneKeyId, phoneKeySecret)

	client, err := dysmsapi.NewClientWithOptions("cn-qingdao", config, credential)
	if err != nil {
		panic(err)
	}

	r := dysmsapi.CreateSendSmsRequest()

	r.SignName = phoneSignName
	r.TemplateCode = phoneTemplateCode
	r.PhoneNumbers = phoneNumber
	r.Scheme = "https"

	smsCode := tool.RandCode()
	code := "{\"code\":" + smsCode + "}"

	r.TemplateParam = code

	response, err := client.SendSms(r)

	if err != nil {
		return errors.New("发送验证码错误")
	} else {
		global.GVA_REDIS.Set(ctx, phoneNumber, smsCode, time.Second*60*15)
		fmt.Printf("response is %#v\n", response)
		return nil
	}

}

// 验证验证码是否正确
func VerifyPhoneCode(phoneNumber, phoneCode string) error {
	code, err := global.GVA_REDIS.Get(ctx, phoneNumber).Result()
	if err != nil {
		return errors.New("查询不到验证码")
	} else {
		fmt.Println("验证码", code, "存储的", phoneCode)
		if code == phoneCode {
			return nil
		} else {
			return errors.New("验证码不正确")
		}
	}
}

func (babyService *BabyService) Info(id string) (baby wechat.Baby, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&baby).Error
	return
}

// 根据uid获取 baby list
func (babyService *BabyService) GetBabyListByUid(uid uint) (babyList []wechat.Baby, err error) {

	err = global.GVA_DB.Where("uid = ?", uid).Find(&babyList).Error
	return
}
