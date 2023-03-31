package wechat

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

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
	if *emergencyAlert.Status == 1 {
		// 发送模板消息
		SendTextMessage(emergencyAlert.ID)
	}
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
	return err
}

// GetEmergencyAlertList
func (emergencyAlertService *EmergencyAlertService) GetEmergencyAlertList() (emergencyAlerts []wechat.EmergencyAlert, err error) {
	err = global.GVA_DB.Find(&emergencyAlerts).Error
	return
}

// 根据id 查询预警通知 并关联baby信息
func (emergencyAlertService *EmergencyAlertService) GetEmergencyAlertById(id int) (emergencyAlert wechat.EmergencyAlert, err error) {
	err = global.GVA_DB.Preload("Baby").Where("id = ?", id).First(&emergencyAlert).Error
	return
}

// 给微信公众号用户发送文本消息
func SendTextMessage(emergencyID uint) (err error) {
	appid := global.GVA_CONFIG.Wechat.Appid
	secret := global.GVA_CONFIG.Wechat.Secret
	if appid == "" || secret == "" {
		return errors.New("appid或secret为空")
	}
	accessToken, err := GetWechatAccessToken(appid, secret)
	if err != nil {
		return errors.New(err.Error())
	}
	fmt.Println("token", accessToken)
	openid := "omzzn5tfhfJ2MTKjXK0oVYTuRmUs"
	content := "消息内容"
	res, err := PushTextMessage(accessToken, openid, content)
	if err != nil {
		return errors.New(err.Error())
	}
	fmt.Println("返回的", res)
	return nil
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

func PushTextMessage(accessToken string, toUser string, content string) (res interface{}, err error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s", accessToken)

	message := TextMessage{
		ToUser:  toUser,
		MsgType: "text",
		Text: struct {
			Content string `json:"content"`
		}{
			Content: content,
		},
	}

	body, err := json.Marshal(message)
	if err != nil {
		return res, err
	}
	res, err = PostJson(url, body)
	return res, err
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
