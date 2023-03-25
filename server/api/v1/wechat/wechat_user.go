package wechat

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	wechatReq "github.com/flipped-aurora/gin-vue-admin/server/model/wechat/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WechatUserApi struct {
}

var wechatUserService = service.ServiceGroupApp.WechatServiceGroup.WechatUserService

// CreateWechatUser 创建WechatUser
// @Tags WechatUser
// @Summary 创建WechatUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wechat.WechatUser true "创建WechatUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wechatUser/createWechatUser [post]
func (wechatUserApi *WechatUserApi) CreateWechatUser(c *gin.Context) {
	var wechatUser wechat.WechatUser
	err := c.ShouldBindJSON(&wechatUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := wechatUserService.CreateWechatUser(wechatUser); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWechatUser 删除WechatUser
// @Tags WechatUser
// @Summary 删除WechatUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wechat.WechatUser true "删除WechatUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wechatUser/deleteWechatUser [delete]
func (wechatUserApi *WechatUserApi) DeleteWechatUser(c *gin.Context) {
	var wechatUser wechat.WechatUser
	err := c.ShouldBindJSON(&wechatUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := wechatUserService.DeleteWechatUser(wechatUser); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWechatUserByIds 批量删除WechatUser
// @Tags WechatUser
// @Summary 批量删除WechatUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除WechatUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wechatUser/deleteWechatUserByIds [delete]
func (wechatUserApi *WechatUserApi) DeleteWechatUserByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := wechatUserService.DeleteWechatUserByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWechatUser 更新WechatUser
// @Tags WechatUser
// @Summary 更新WechatUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wechat.WechatUser true "更新WechatUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wechatUser/updateWechatUser [put]
func (wechatUserApi *WechatUserApi) UpdateWechatUser(c *gin.Context) {
	var wechatUser wechat.WechatUser
	err := c.ShouldBindJSON(&wechatUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := wechatUserService.UpdateWechatUser(wechatUser); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWechatUser 用id查询WechatUser
// @Tags WechatUser
// @Summary 用id查询WechatUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wechat.WechatUser true "用id查询WechatUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wechatUser/findWechatUser [get]
func (wechatUserApi *WechatUserApi) FindWechatUser(c *gin.Context) {
	var wechatUser wechat.WechatUser
	err := c.ShouldBindQuery(&wechatUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rewechatUser, err := wechatUserService.GetWechatUser(wechatUser.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewechatUser": rewechatUser}, c)
	}
}

// GetWechatUserList 分页获取WechatUser列表
// @Tags WechatUser
// @Summary 分页获取WechatUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wechatReq.WechatUserSearch true "分页获取WechatUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wechatUser/getWechatUserList [get]
func (wechatUserApi *WechatUserApi) GetWechatUserList(c *gin.Context) {
	var pageInfo wechatReq.WechatUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wechatUserService.GetWechatUserInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// 根据code获取微信公众号openid 并保存到wechatUser
func (wechatUserApi *WechatUserApi) Login(c *gin.Context) {
	var code wechatReq.WechatCode
	err := c.ShouldBindJSON(&code)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if code.Code == "" {
		response.FailWithMessage("code不能为空", c)
		return
	}

	if token, err := wechatUserService.Login(code.Code); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(token, c)
	}

}

// 根据id来保存用户信息
func (wechatUserApi *WechatUserApi) SaveUserInfo(c *gin.Context) {
	var userInfo wechatReq.WechatUserInfo
	err := c.ShouldBindJSON(&userInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	id := c.MustGet("id").(uint)
	if err := wechatUserService.SaveUserInfo(id, userInfo); err != nil {
		global.GVA_LOG.Error("保存失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("保存成功", c)
	}
}

func (wechatUserApi *WechatUserApi) UploadFile(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("error: %s", err.Error()))
		return
	}

	// 根据日期创建目录
	dateDir := time.Now().Format("20060102")
	targetDir := filepath.Join("uploads/pic", dateDir)
	fmt.Println("图片路径", targetDir)
	// 创建目录
	if err := os.MkdirAll(targetDir, os.ModePerm); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 生成文件名
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

	// 保存文件
	targetPath := filepath.Join(targetDir, filename)
	if err := c.SaveUploadedFile(file, targetPath); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(gin.H{"path": targetPath}, c)
	}
}
