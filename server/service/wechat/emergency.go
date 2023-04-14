package wechat

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	wechatReq "github.com/flipped-aurora/gin-vue-admin/server/model/wechat/request"
)

type EmergencyAlertService struct {
}

// CreateEmergencyAlert 创建EmergencyAlert记录
// Author [piexlmax](https://github.com/piexlmax)
func (emergencyAlertService *EmergencyAlertService) CreateEmergencyAlert(emergencyAlert wechat.EmergencyAlert) (err error) {
	err = global.GVA_DB.Create(&emergencyAlert).Error
	return err
}

// DeleteEmergencyAlert 删除EmergencyAlert记录
// Author [piexlmax](https://github.com/piexlmax)
func (emergencyAlertService *EmergencyAlertService) DeleteEmergencyAlert(emergencyAlert wechat.EmergencyAlert) (err error) {
	err = global.GVA_DB.Delete(&emergencyAlert).Error
	return err
}

// DeleteEmergencyAlertByIds 批量删除EmergencyAlert记录
// Author [piexlmax](https://github.com/piexlmax)
func (emergencyAlertService *EmergencyAlertService) DeleteEmergencyAlertByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]wechat.EmergencyAlert{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateEmergencyAlert 更新EmergencyAlert记录
// Author [piexlmax](https://github.com/piexlmax)
func (emergencyAlertService *EmergencyAlertService) UpdateEmergencyAlert(emergencyAlert wechat.EmergencyAlert) (err error) {
	// if *emergencyAlert.Status == 1 {
	// 	// 发送模板消息
	// 	SendTextMessage(emergencyAlert)
	// }
	err = global.GVA_DB.Save(&emergencyAlert).Error
	return err
}

// GetEmergencyAlert 根据id获取EmergencyAlert记录
// Author [piexlmax](https://github.com/piexlmax)
func (emergencyAlertService *EmergencyAlertService) GetEmergencyAlert(id uint) (emergencyAlert wechat.EmergencyAlert, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&emergencyAlert).Error
	return
}

// GetEmergencyAlertInfoList 分页获取EmergencyAlert记录
// Author [piexlmax](https://github.com/piexlmax)
func (emergencyAlertService *EmergencyAlertService) GetEmergencyAlertInfoList(info wechatReq.EmergencyAlertSearch) (list []wechat.EmergencyAlert, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&wechat.EmergencyAlert{})
	var emergencyAlerts []wechat.EmergencyAlert
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&emergencyAlerts).Error
	return emergencyAlerts, total, err
}

// save a emergencyAlert record
func (emergencyAlertService *EmergencyAlertService) Issue(emergencyAlert wechat.EmergencyAlert) (err error) {

	err = global.GVA_DB.Save(&emergencyAlert).Error
	if err == nil {
		//修改babyId status 为2
		err = global.GVA_DB.Model(&wechat.Baby{}).Where("id = ?", emergencyAlert.BabyId).Update("status", 2).Error

		SendTextMessage(emergencyAlert)
	}
	return err
}

type TemplateList struct {
	Openid   string `json:"openid"`
	Appid    string `json:"appid"`
	Secret   string `json:"secret"`
	Token    string `json:"token"`
	Phone    string `json:"phone"`
	Remark   string `json:"remark"`
	BabyId   string `json:"babyId"`
	LostTime string `json:"lostTime"`
	Url      string `json:"url"`
}

// 发送订阅消息给用户
// func PushTextMessage(t *templateList) (err error) {
// 	wc := sWechat.NewWechat()
// 	memory := cache.NewMemory()
// 	cfg := &config.Config{
// 		AppID:     t.Appid,
// 		AppSecret: t.Secret,
// 		Token:     t.Token,
// 		Cache:     memory,
// 	}
// 	officialAccount := wc.GetOfficialAccount(cfg)
// 	s := officialAccount.GetSubscribe()
// 	var sm message.SubscribeMessage
// 	sm.ToUser = t.Openid
// 	sm.TemplateID = "bPXvnSLbyh9GBIrW22fdIJnSrlU4wBm_7eCExLKA8qM"
// 	sm.Page = "https://huzhu.cnecip.com/wechat/"
// 	sm.Data = map[string]*message.SubscribeDataItem{
// 		"thing2": {
// 			Value: "测试",
// 		},
// 		"date3": {
// 			Value: "2022-01-12",
// 		},
// 		"thing4": {
// 			Value: "测试",
// 		},
// 	}

// 	return s.Send(&sm)
// }

type SubscribeMessage struct {
	ToUser      string                        `json:"touser"`        // 必须, 接受者OpenID
	TemplateID  string                        `json:"template_id"`   // 必须, 模版ID
	Url         string                        `json:"url,omitempty"` // 可选, 跳转网页时填写
	Data        map[string]*SubscribeDataItem `json:"data"`          // 必须, 模板数据
	MiniProgram struct {
		AppID    string `json:"appid"`    // 所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系）
		PagePath string `json:"pagepath"` // 所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar）
	} `json:"miniprogram"` // 可选,跳转至小程序地址
}

// SubscribeDataItem 模版内某个 .DATA 的值
type SubscribeDataItem struct {
	Value string `json:"value"`
}

// 发送模板消息给用户
func SendTextMessageToUser(t *TemplateList) (err error) {
	var sm SubscribeMessage
	sm.ToUser = t.Openid
	sm.TemplateID = "1f46L-66R2WZAeqk-sH60nmP44-_vMKE1eAv27-EjFQ"
	sm.Url = t.Url
	//当前日期 年月日时分秒
	now := time.Now().Format("2006-01-02 15:04:05")
	sm.Data = map[string]*SubscribeDataItem{
		"first": {
			Value: "人员信息",
		},
		"keyword1": {
			Value: t.Phone,
		},
		"keyword2": {
			Value: now,
		},
		"remark": {
			Value: t.Remark,
		},
	}
	//发送post请求
	url := "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=" + t.Token
	jsonStr, _ := json.Marshal(sm)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return

}

// GetEmergencyAlertList
func (emergencyAlertService *EmergencyAlertService) GetEmergencyAlertList() (emergencyAlerts []wechat.EmergencyAlert, err error) {
	err = global.GVA_DB.Order("id DESC").Find(&emergencyAlerts).Error
	return
}

// 根据id 查询预警通知 并关联baby信息
func (emergencyAlertService *EmergencyAlertService) GetEmergencyAlertById(id int) (emergencyAlert wechat.EmergencyAlert, err error) {
	err = global.GVA_DB.Preload("Baby").Where("id = ?", id).First(&emergencyAlert).Error
	return
}

// 给微信公众号用户发送文本消息
func SendTextMessage(e wechat.EmergencyAlert) (err error) {
	appid := global.GVA_CONFIG.Wechat.Appid
	secret := global.GVA_CONFIG.Wechat.Secret
	if appid == "" || secret == "" {
		return errors.New("appid或secret为空")
	}
	accessToken, err := GetWechatAccessToken(appid, secret)
	if err != nil {
		return errors.New(err.Error())
	}
	//从e.ContactPerson  json字符中 中解析出phone
	type ContactPerson struct {
		Phone   string `json:"phone"`
		ShenFen string `json:"shenfen"`
		Code    string `json:"code"`
		Checked bool   `json:"checked"`
	}
	var contactPerson []ContactPerson
	_ = json.Unmarshal([]byte(e.ContactPerson), &contactPerson)
	fmt.Println("contactPerson联系人", contactPerson)
	var t TemplateList
	t.Appid = appid
	t.Token = accessToken
	t.Secret = secret
	t.LostTime = e.LostTime
	t.Remark = "走失地址：" + e.LostLocation
	t.Phone = contactPerson[0].Phone
	//babayId 转为字符串
	t.BabyId = strconv.Itoa(*e.BabyId)
	id := strconv.Itoa(int(e.ID))
	fmt.Println("发送预警的babyId:", t.BabyId)
	t.Url = "https://huzhu.cnecip.com/wechat/#/pages/zhoubian/detail?id=" + id
	//查询所有用户的openid
	var users []wechat.WechatUser
	err = global.GVA_DB.Find(&users).Error
	if err != nil {
		return errors.New(err.Error())
	}
	for _, v := range users {
		t.Openid = v.Openid
		err = SendTextMessageToUser(&t)
	}
	fmt.Println("发送状态", err)
	return
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

func GetWechatAccessToken(appID, appSecret string) (string, error) {
	// 获取 access_token
	tokenURL := "https://api.weixin.qq.com/cgi-bin/token"
	tokenURL += fmt.Sprintf("?grant_type=client_credential&appid=%s&secret=%s", appID, appSecret)
	tokenResp, err := http.Get(tokenURL)
	if err != nil {
		panic(err)
	}
	defer tokenResp.Body.Close()

	var token AccessTokenResponse
	err = json.NewDecoder(tokenResp.Body).Decode(&token)
	if err != nil {
		panic(err)
	}
	return token.AccessToken, nil
}

type TextMessage struct {
	ToUser  string `json:"touser"`
	MsgType string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

func PostJson(url string, jsonBytes []byte) (str string, err error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("body体数据错误", err.Error())
		return
	}

	if resp.StatusCode != http.StatusOK {
		// 处理请求失败的情况
		return "", errors.New("请求失败")
	}
	str = readString(respBytes)
	fmt.Println("返回的字符串数据", str)
	return str, nil
}

func readString(respBytes []byte) string {
	if len(respBytes) == 0 {
		return ""
	}
	return string(respBytes)
}
