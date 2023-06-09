package wechat

import (
	"fmt"
	"image/jpeg"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	wechatReq "github.com/flipped-aurora/gin-vue-admin/server/model/wechat/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
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
	var userInfo wechatReq.WechatUserInfo
	err := c.ShouldBindJSON(&userInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if userInfo.Code == "" {
		response.FailWithMessage("code不能为空", c)
		return
	}

	if token, exist, err := wechatUserService.Login(userInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(gin.H{"token": token, "exist": exist}, c)
	}

}

// save user api
func (wechatUserApi *WechatUserApi) SaveUser(c *gin.Context) {
	var userInfo wechat.WechatUser
	err := c.ShouldBindJSON(&userInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if userInfo.Phone == "" {
		response.FailWithMessage("手机号不能为空", c)
		return
	}
	userInfo.ID = c.MustGet("id").(uint)
	if err = wechatUserService.SaveWechatUser(userInfo); err != nil {
		global.GVA_LOG.Error("保存失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("保存成功", c)
	}
}

func (wechatUserApi *WechatUserApi) UploadFile(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 根据日期创建目录
	dateDir := time.Now().Format("20060102")
	targetDir := filepath.Join("uploads/file/pic", dateDir)
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

// 压缩图片
func (wechatUserApi *WechatUserApi) CompressImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 打开上传的文件
	f, err := file.Open()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	defer f.Close()
	// 解码文件
	img, err := jpeg.Decode(f)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 调整大小
	resized := resize.Resize(600, 0, img, resize.Lanczos3)

	// 根据日期创建目录
	dateDir := time.Now().Format("20060102")
	targetDir := filepath.Join("uploads/file/pic", dateDir)
	fmt.Println("图片路径", targetDir)
	// 创建目录
	if err := os.MkdirAll(targetDir, os.ModePerm); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 生成文件名
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

	targetPath := filepath.Join(targetDir, filename)
	// 创建一个新的文件
	outFile, err := os.Create(targetPath)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	defer outFile.Close()

	// 将压缩后的图片写入新的文件中
	if err := jpeg.Encode(outFile, resized, nil); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(gin.H{"path": targetPath}, c)
		return
	}
}
