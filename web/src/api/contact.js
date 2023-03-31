import service from '@/utils/request'

// @Tags Contact
// @Summary 创建Contact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Contact true "创建Contact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /contact/createContact [post]
export const createContact = (data) => {
  return service({
    url: '/contact/createContact',
    method: 'post',
    data
  })
}

// @Tags Contact
// @Summary 删除Contact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Contact true "删除Contact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /contact/deleteContact [delete]
export const deleteContact = (data) => {
  return service({
    url: '/contact/deleteContact',
    method: 'delete',
    data
  })
}

// @Tags Contact
// @Summary 删除Contact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Contact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /contact/deleteContact [delete]
export const deleteContactByIds = (data) => {
  return service({
    url: '/contact/deleteContactByIds',
    method: 'delete',
    data
  })
}

// @Tags Contact
// @Summary 更新Contact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Contact true "更新Contact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /contact/updateContact [put]
export const updateContact = (data) => {
  return service({
    url: '/contact/updateContact',
    method: 'put',
    data
  })
}

// @Tags Contact
// @Summary 用id查询Contact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Contact true "用id查询Contact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /contact/findContact [get]
export const findContact = (params) => {
  return service({
    url: '/contact/findContact',
    method: 'get',
    params
  })
}

// @Tags Contact
// @Summary 分页获取Contact列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Contact列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /contact/getContactList [get]
export const getContactList = (params) => {
  return service({
    url: '/contact/getContactList',
    method: 'get',
    params
  })
}
