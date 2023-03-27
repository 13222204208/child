import service from '@/utils/request'

// @Tags Agency
// @Summary 创建Agency
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Agency true "创建Agency"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /agency/createAgency [post]
export const createAgency = (data) => {
  return service({
    url: '/agency/createAgency',
    method: 'post',
    data
  })
}

// @Tags Agency
// @Summary 删除Agency
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Agency true "删除Agency"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /agency/deleteAgency [delete]
export const deleteAgency = (data) => {
  return service({
    url: '/agency/deleteAgency',
    method: 'delete',
    data
  })
}

// @Tags Agency
// @Summary 删除Agency
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Agency"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /agency/deleteAgency [delete]
export const deleteAgencyByIds = (data) => {
  return service({
    url: '/agency/deleteAgencyByIds',
    method: 'delete',
    data
  })
}

// @Tags Agency
// @Summary 更新Agency
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Agency true "更新Agency"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /agency/updateAgency [put]
export const updateAgency = (data) => {
  return service({
    url: '/agency/updateAgency',
    method: 'put',
    data
  })
}

// @Tags Agency
// @Summary 用id查询Agency
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Agency true "用id查询Agency"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /agency/findAgency [get]
export const findAgency = (params) => {
  return service({
    url: '/agency/findAgency',
    method: 'get',
    params
  })
}

// @Tags Agency
// @Summary 分页获取Agency列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Agency列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /agency/getAgencyList [get]
export const getAgencyList = (params) => {
  return service({
    url: '/agency/getAgencyList',
    method: 'get',
    params
  })
}
