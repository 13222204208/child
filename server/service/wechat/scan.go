package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"

	"github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	wechatReq "github.com/flipped-aurora/gin-vue-admin/server/model/wechat/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wechat/response"
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

// 获取扫描记录
func (scanService *ScanService) GetScanList(babyId uint64) (list []wechat.Scan, err error) {
	err = global.GVA_DB.Preload("WechatUser").Order("id DESC").Where("baby_id = ?", babyId).Find(&list).Error
	return
}

// Contrast 对比 两张图片
func (scanService *ScanService) Contrast(pic string, lng, lat float64, uid uint) (info wechat.Baby, eid uint, err error) {
	// res, err := GetEmergencyBaby()
	// if err != nil {
	// 	return
	// }
	//lat和lng string转换成float64
	//字符串转float64

	address, err := GetAddressByLatLng(lat, lng)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Address:", address)
	personId, err := response.SearchFace(pic, "personInfo")
	if err != nil {
		return
	}
	if personId == "" {
		err = errors.New("未检测到人脸")
		return
	} else {
		err = global.GVA_DB.Where("person_id = ?", personId).First(&info).Error
		if err != nil {
			return
		}
		err = SaveScan(uid, info.ID, address, pic)
		if err != nil {
			return
		}
		//查询出预警信息的id
		var e wechat.EmergencyAlert
		err = global.GVA_DB.Order("id DESC").Where("baby_id = ?", info.ID).First(&e).Error

		fmt.Println("检测到的人员信息", info)
		return info, e.ID, err
	}

	// var scoreNum float64
	// for _, v := range res {
	// 	if len(v.Feature) < 10 {
	// 		continue
	// 	}
	// 	img, err := ParseJsonArray(v.Feature)
	// 	if err != nil {
	// 		return info, err
	// 	}
	// 	fmt.Println("对比的图片:", img)
	// 	fmt.Println("扫描的图片:", pic)
	// 	picBase64, _ := copilot.PrefixImgBase64(pic)
	// 	imgBase64, _ := copilot.PrefixImgBase64(img)

	// 	credential := common.NewCredential(
	// 		global.GVA_CONFIG.TencentCOS.SecretID,
	// 		global.GVA_CONFIG.TencentCOS.SecretKey,
	// 	)
	// 	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	// 	cpf := profile.NewClientProfile()
	// 	cpf.HttpProfile.Endpoint = "iai.tencentcloudapi.com"
	// 	// 实例化要请求产品的client对象,clientProfile是可选的
	// 	client, _ := iai.NewClient(credential, "ap-beijing", cpf)

	// 	// 实例化一个请求对象,每个接口都会对应一个request对象
	// 	request := iai.NewCompareFaceRequest()
	// 	request.ImageA = &picBase64
	// 	request.ImageB = &imgBase64
	// 	// 返回的resp是一个CompareFaceResponse的实例，与请求对象对应
	// 	response, err := client.CompareFace(request)
	// 	if _, ok := err.(*tencentErrors.TencentCloudSDKError); ok {
	// 		fmt.Printf("An API error has returned: %s", err)
	// 		continue
	// 	}
	// 	if err != nil {
	// 		// panic(err)
	// 		fmt.Println("err错误", err)
	// 		continue
	// 	} else {
	// 		// 输出json格式的字符串回包
	// 		score := response.Response.Score

	// 		fmt.Println("resMap返回的相似度", *score, "地址", address, "图片", pic)
	// 		scoreNum = *score
	// 		if *score > 80 {
	// 			err = SaveScan(uid, v.ID, address, pic)
	// 			if err != nil {
	// 				fmt.Println("保存错误", err)
	// 				return info, err
	// 			} else {
	// 				return v, nil
	// 			}
	// 		}
	// 	}

	// }
	// if scoreNum < 80 {
	// 	return info, errors.New("没有找到对应的信息")
	// }

}

const (
	TencentMapAPI = "https://apis.map.qq.com/ws/geocoder/v1/"
	TencentMapKey = "PSXBZ-Y2JHU-IQMVL-GSWQX-7QLI6-5UBJW" // 请替换为你自己的API密钥
)

type Response struct {
	Status    int    `json:"status"`
	Message   string `json:"message"`
	Result    Result `json:"result"`
	RequestID string `json:"request_id"`
}

type Result struct {
	Address          string           `json:"address"`
	AddressComponent AddressComponent `json:"address_component"`
}

type AddressComponent struct {
	Nation       string `json:"nation"`
	Province     string `json:"province"`
	City         string `json:"city"`
	District     string `json:"district"`
	Street       string `json:"street"`
	StreetNumber string `json:"street_number"`
}

func GetAddressByLatLng(lat, lng float64) (string, error) {
	params := url.Values{}
	params.Add("key", TencentMapKey)
	params.Add("location", fmt.Sprintf("%f,%f", lat, lng))

	apiURL := TencentMapAPI + "?" + params.Encode()
	resp, err := http.Get(apiURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	if response.Status != 0 {
		return "", fmt.Errorf("error: %s", response.Message)
	}

	return response.Result.Address, nil
}

// 查询出所有发布预警的信息
func GetEmergencyBaby() (baby []wechat.Baby, err error) {
	//查询出所有EmergencyAlert 中的 baby_id 并去重
	// var emergencyAlert []uint
	// err = global.GVA_DB.Model(&wechat.EmergencyAlert{}).Pluck("baby_id", &emergencyAlert).Error
	// if err != nil {
	// 	return
	// }
	// fmt.Println("查询出的baby_id", emergencyAlert, "长度", len(emergencyAlert))
	// //根据 baby_id 查询出所有的baby
	// err = global.GVA_DB.Model(&wechat.Baby{}).Where("id in ?", emergencyAlert).Find(&baby).Error
	//查出所有的babyId 且 feature 不为空的
	err = global.GVA_DB.Model(&wechat.Baby{}).Where("feature != ?", "").Find(&baby).Error
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
func SaveScan(uid, babyId uint, address, pic string) (err error) {
	var scan wechat.Scan
	scan.Uid = &uid
	scan.BabyId = &babyId
	scan.Address = address
	scan.Pic = pic
	err = global.GVA_DB.Create(&scan).Error
	if err != nil {
		return err
	}
	//发送消息给家长
	err = SendScanMsgToParent(uid, babyId, address)
	return err
}

// uid 扫描人的id
// 扫描的宝贝id
func SendScanMsgToParent(uid, babyId uint, address string) (err error) {

	appid := global.GVA_CONFIG.Wechat.Appid
	secret := global.GVA_CONFIG.Wechat.Secret
	if appid == "" || secret == "" {
		return errors.New("appid或secret为空")
	}
	accessToken, err := GetWechatAccessToken(appid, secret)
	if err != nil {
		return errors.New(err.Error())
	}

	//获取宝贝信息
	babyService := BabyService{}
	baby, err := babyService.GetBaby(babyId)
	if err != nil {
		return err
	}
	//获取宝贝的家长信息
	parentService := WechatUserService{}
	parents, err := parentService.GetWechatUser(baby.Uid)
	if err != nil {
		return err
	}

	scanUser, err := parentService.GetWechatUser(uid)
	if err != nil {
		return err
	}

	//babyId 转换为string
	// var babyIdStr string
	babyIdStr := strconv.Itoa(int(babyId))
	var t TemplateList
	t.Appid = appid
	t.Token = accessToken
	t.Secret = secret
	t.Openid = parents.Openid
	t.Phone = scanUser.Phone
	t.BabyId = babyIdStr
	t.Remark = scanUser.Name + "在" + address + ",扫描了" + baby.Name
	t.Url = "https://huzhu.cnecip.com/wechat/#/pages/zhoubian/saomiao"
	//当前时间
	t.LostTime = time.Now().Format("2006-01-02 15:04:05")
	err = SendTextMessageToUser(&t)
	return err
}
