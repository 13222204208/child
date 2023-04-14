package wechat

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	wechatReq "github.com/flipped-aurora/gin-vue-admin/server/model/wechat/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
)

type ScanApi struct {
}

var scanService = service.ServiceGroupApp.WechatServiceGroup.ScanService

// CreateScan 创建Scan
// @Tags Scan
// @Summary 创建Scan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wechat.Scan true "创建Scan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /scan/createScan [post]
func (scanApi *ScanApi) CreateScan(c *gin.Context) {
	var scan wechat.Scan
	err := c.ShouldBindJSON(&scan)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := scanService.CreateScan(scan); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteScan 删除Scan
// @Tags Scan
// @Summary 删除Scan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wechat.Scan true "删除Scan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /scan/deleteScan [delete]
func (scanApi *ScanApi) DeleteScan(c *gin.Context) {
	var scan wechat.Scan
	err := c.ShouldBindJSON(&scan)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := scanService.DeleteScan(scan); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteScanByIds 批量删除Scan
// @Tags Scan
// @Summary 批量删除Scan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Scan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /scan/deleteScanByIds [delete]
func (scanApi *ScanApi) DeleteScanByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := scanService.DeleteScanByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateScan 更新Scan
// @Tags Scan
// @Summary 更新Scan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wechat.Scan true "更新Scan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /scan/updateScan [put]
func (scanApi *ScanApi) UpdateScan(c *gin.Context) {
	var scan wechat.Scan
	err := c.ShouldBindJSON(&scan)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := scanService.UpdateScan(scan); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindScan 用id查询Scan
// @Tags Scan
// @Summary 用id查询Scan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wechat.Scan true "用id查询Scan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /scan/findScan [get]
func (scanApi *ScanApi) FindScan(c *gin.Context) {
	var scan wechat.Scan
	err := c.ShouldBindQuery(&scan)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rescan, err := scanService.GetScan(scan.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rescan": rescan}, c)
	}
}

// GetScanList 分页获取Scan列表
// @Tags Scan
// @Summary 分页获取Scan列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wechatReq.ScanSearch true "分页获取Scan列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /scan/getScanList [get]
func (scanApi *ScanApi) GetScanList(c *gin.Context) {
	var pageInfo wechatReq.ScanSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := scanService.GetScanInfoList(pageInfo); err != nil {
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

// save scan contrast api 扫描对比
func (scanApi *ScanApi) Contrast(c *gin.Context) {
	var scan wechatReq.ScanCompare
	err := c.ShouldBindJSON(&scan)
	fmt.Println("请求的数据", scan)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if scan.Pic == "" {
		response.FailWithMessage("图片不能为空", c)
		return
	}

	if scan.Lng == 0 || scan.Lat == 0 {
		response.FailWithMessage("经纬度不能为空", c)
		return
	}

	uid := c.MustGet("id").(uint)
	if info, eid, err := scanService.Contrast(scan.Pic, scan.Lng, scan.Lat, uid); err != nil {
		global.GVA_LOG.Error("保存失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(gin.H{"info": info, "eid": eid}, c)
	}
}

// 扫描记录
func (scanApi *ScanApi) ScanRecord(c *gin.Context) {
	id := c.Param("id")
	//id 转为uint
	idUint, _ := strconv.ParseUint(id, 10, 64)
	if idUint == 0 {
		response.FailWithMessage("宝贝id不能为空", c)
		return
	}

	if list, err := scanService.GetScanList(idUint); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"list": list}, c)
	}
}
