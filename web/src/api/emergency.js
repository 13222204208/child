import service from '@/utils/request'

// @Tags EmergencyAlert
// @Summary 创建EmergencyAlert
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EmergencyAlert true "创建EmergencyAlert"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /emergencyAlert/createEmergencyAlert [post]
export const createEmergencyAlert = (data) => {
  return service({
    url: '/emergencyAlert/createEmergencyAlert',
    method: 'post',
    data
  })
}

// @Tags EmergencyAlert
// @Summary 删除EmergencyAlert
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EmergencyAlert true "删除EmergencyAlert"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /emergencyAlert/deleteEmergencyAlert [delete]
export const deleteEmergencyAlert = (data) => {
  return service({
    url: '/emergencyAlert/deleteEmergencyAlert',
    method: 'delete',
    data
  })
}

// @Tags EmergencyAlert
// @Summary 删除EmergencyAlert
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EmergencyAlert"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /emergencyAlert/deleteEmergencyAlert [delete]
export const deleteEmergencyAlertByIds = (data) => {
  return service({
    url: '/emergencyAlert/deleteEmergencyAlertByIds',
    method: 'delete',
    data
  })
}

// @Tags EmergencyAlert
// @Summary 更新EmergencyAlert
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EmergencyAlert true "更新EmergencyAlert"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /emergencyAlert/updateEmergencyAlert [put]
export const updateEmergencyAlert = (data) => {
  return service({
    url: '/emergencyAlert/updateEmergencyAlert',
    method: 'put',
    data
  })
}

// @Tags EmergencyAlert
// @Summary 用id查询EmergencyAlert
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.EmergencyAlert true "用id查询EmergencyAlert"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /emergencyAlert/findEmergencyAlert [get]
export const findEmergencyAlert = (params) => {
  return service({
    url: '/emergencyAlert/findEmergencyAlert',
    method: 'get',
    params
  })
}

// @Tags EmergencyAlert
// @Summary 分页获取EmergencyAlert列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取EmergencyAlert列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /emergencyAlert/getEmergencyAlertList [get]
export const getEmergencyAlertList = (params) => {
  return service({
    url: '/emergencyAlert/getEmergencyAlertList',
    method: 'get',
    params
  })
}
