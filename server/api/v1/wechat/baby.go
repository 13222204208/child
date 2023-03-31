package wechat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	wechatReq "github.com/flipped-aurora/gin-vue-admin/server/model/wechat/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BabyApi struct {
}

var babyService = service.ServiceGroupApp.WechatServiceGroup.BabyService

// CreateBaby 创建Baby
// @Tags Baby
// @Summary 创建Baby
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wechat.Baby true "创建Baby"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /baby/createBaby [post]
func (babyApi *BabyApi) CreateBaby(c *gin.Context) {
	var baby wechat.Baby
	err := c.ShouldBindJSON(&baby)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":   {utils.NotEmpty()},
		"Gender": {utils.NotEmpty()},
	}
	if err := utils.Verify(baby, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := babyService.CreateBaby(baby); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBaby 删除Baby
// @Tags Baby
// @Summary 删除Baby
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wechat.Baby true "删除Baby"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /baby/deleteBaby [delete]
func (babyApi *BabyApi) DeleteBaby(c *gin.Context) {
	var baby wechat.Baby
	err := c.ShouldBindJSON(&baby)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := babyService.DeleteBaby(baby); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBabyByIds 批量删除Baby
// @Tags Baby
// @Summary 批量删除Baby
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Baby"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /baby/deleteBabyByIds [delete]
func (babyApi *BabyApi) DeleteBabyByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := babyService.DeleteBabyByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBaby 更新Baby
// @Tags Baby
// @Summary 更新Baby
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wechat.Baby true "更新Baby"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /baby/updateBaby [put]
func (babyApi *BabyApi) UpdateBaby(c *gin.Context) {
	var baby wechat.Baby
	err := c.ShouldBindJSON(&baby)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":   {utils.NotEmpty()},
		"Gender": {utils.NotEmpty()},
	}
	if err := utils.Verify(baby, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := babyService.UpdateBaby(baby); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBaby 用id查询Baby
// @Tags Baby
// @Summary 用id查询Baby
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wechat.Baby true "用id查询Baby"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /baby/findBaby [get]
func (babyApi *BabyApi) FindBaby(c *gin.Context) {
	var baby wechat.Baby
	err := c.ShouldBindQuery(&baby)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rebaby, err := babyService.GetBaby(baby.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rebaby": rebaby}, c)
	}
}

// GetBabyList 分页获取Baby列表
// @Tags Baby
// @Summary 分页获取Baby列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wechatReq.BabySearch true "分页获取Baby列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /baby/getBabyList [get]
func (babyApi *BabyApi) GetBabyList(c *gin.Context) {
	var pageInfo wechatReq.BabySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := babyService.GetBabyInfoList(pageInfo); err != nil {
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

// 保存baby的信息
func (babyApi *BabyApi) SaveBaby(c *gin.Context) {
	var baby wechat.Baby
	err := c.ShouldBindJSON(&baby)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	baby.Uid = c.MustGet("id").(uint)

	if babyId, err := babyService.SaveBabyInfo(baby); err != nil {
		global.GVA_LOG.Error("保存失败!", zap.Error(err))
		response.FailWithMessage("保存失败", c)
	} else {
		response.OkWithData(gin.H{"babyId": babyId}, c)
	}
}

func (babyApi *BabyApi) Info(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.FailWithMessage("id不能为空", c)
		return
	}
	if rebaby, err := babyService.Info(id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rebaby": rebaby}, c)
	}
}

// get babylist api
func (babyApi *BabyApi) List(c *gin.Context) {
	uid := c.MustGet("id").(uint)
	if list, err := babyService.GetBabyListByUid(uid); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"list": list}, c)
	}
}

// sendCode api
func (babyApi *BabyApi) SendCode(c *gin.Context) {
	var p wechatReq.SendCode
	err := c.ShouldBindJSON(&p)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := babyService.SendCode(p.Phone); err != nil {
		global.GVA_LOG.Error("发送失败!", zap.Error(err))
		response.FailWithMessage("发送失败", c)
	} else {
		response.OkWithMessage("发送成功", c)
	}
}
