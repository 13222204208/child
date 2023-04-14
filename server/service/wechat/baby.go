package wechat

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/13222204208/copilot"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tencentErrors "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	wechatReq "github.com/flipped-aurora/gin-vue-admin/server/model/wechat/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wechat/response"
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

	if baby.Status == 1 {
		//删除宝贝预警信息
		err = global.GVA_DB.Where("baby_id = ?", baby.ID).Delete(&wechat.EmergencyAlert{}).Error
	}
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
		if baby.Image != "" {
			faceId, err := response.SaveFaceId(baby.Image, baby.ID)
			if err != nil {
				return 0, err
			}
			if faceId != "" {
				baby.FaceId = faceId
			}
		}

		err = global.GVA_DB.Model(&baby).Updates(baby).Error
		babyId = baby.ID
		return
	}

	var babyInfo wechat.Baby
	err = global.GVA_DB.Where("uid = ? and name = ?", baby.Uid, baby.Name).First(&babyInfo).Error
	babyId = babyInfo.ID
	fmt.Println("babyId", babyId)
	if err != nil {
		baby.PersonId = copilot.UserNum()
		baby.GroupId = "personInfo"
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
	credential := common.NewCredential(
		global.GVA_CONFIG.TencentCOS.SecretID,
		global.GVA_CONFIG.TencentCOS.SecretKey,
	)

	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := sms.NewClient(credential, "ap-beijing", cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := sms.NewSendSmsRequest()
	request.PhoneNumberSet = common.StringPtrs([]string{phoneNumber})
	request.TemplateId = common.StringPtr("1757796")
	request.SmsSdkAppId = common.StringPtr("1400809494")
	request.SignName = common.StringPtr("全民互助公众号")
	//生成六位随机数
	smsCode := strconv.Itoa(rand.Intn(899999) + 100000)
	request.TemplateParamSet = common.StringPtrs([]string{smsCode})
	// 返回的resp是一个SendSmsResponse的实例，与请求对象对应
	fmt.Println("请求的数据", request.ToJsonString())
	response, err := client.SendSms(request)
	if _, ok := err.(*tencentErrors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return errors.New("发送验证码错误")
	}
	if err != nil {
		return errors.New("发送验证码错误")
	}
	fmt.Printf("输出： %s", response.ToJsonString())
	code := *response.Response.SendStatusSet[0].Code
	fmt.Println(code, "code")
	// 输出json格式的字符串回包
	if code != "Ok" {
		fmt.Println(*response.Response.SendStatusSet[0].Code)
		return errors.New("发送验证码错误")
	}

	global.GVA_REDIS.Set(ctx, phoneNumber, smsCode, time.Second*60*15)
	return err
}

// 添加验证手机号是否正确服务
func (babyService *BabyService) VerifyPhone(phoneNumber, phoneCode string) (err error) {
	err = VerifyPhoneCode(phoneNumber, phoneCode)
	return err
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
