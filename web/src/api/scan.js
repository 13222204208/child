import service from '@/utils/request'

// @Tags Scan
// @Summary 创建Scan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Scan true "创建Scan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /scan/createScan [post]
export const createScan = (data) => {
  return service({
    url: '/scan/createScan',
    method: 'post',
    data
  })
}

// @Tags Scan
// @Summary 删除Scan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Scan true "删除Scan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /scan/deleteScan [delete]
export const deleteScan = (data) => {
  return service({
    url: '/scan/deleteScan',
    method: 'delete',
    data
  })
}

// @Tags Scan
// @Summary 删除Scan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Scan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /scan/deleteScan [delete]
export const deleteScanByIds = (data) => {
  return service({
    url: '/scan/deleteScanByIds',
    method: 'delete',
    data
  })
}

// @Tags Scan
// @Summary 更新Scan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Scan true "更新Scan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /scan/updateScan [put]
export const updateScan = (data) => {
  return service({
    url: '/scan/updateScan',
    method: 'put',
    data
  })
}

// @Tags Scan
// @Summary 用id查询Scan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Scan true "用id查询Scan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /scan/findScan [get]
export const findScan = (params) => {
  return service({
    url: '/scan/findScan',
    method: 'get',
    params
  })
}

// @Tags Scan
// @Summary 分页获取Scan列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Scan列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /scan/getScanList [get]
export const getScanList = (params) => {
  return service({
    url: '/scan/getScanList',
    method: 'get',
    params
  })
}
