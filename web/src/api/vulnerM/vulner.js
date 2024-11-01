import service from '@/utils/request'
// @Tags Vulner
// @Summary 创建脆弱性管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Vulner true "创建脆弱性管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /vulner/createVulner [post]
export const createVulner = (data) => {
  return service({
    url: '/vulner/createVulner',
    method: 'post',
    data
  })
}

// @Tags Vulner
// @Summary 删除脆弱性管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Vulner true "删除脆弱性管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /vulner/deleteVulner [delete]
export const deleteVulner = (params) => {
  return service({
    url: '/vulner/deleteVulner',
    method: 'delete',
    params
  })
}

// @Tags Vulner
// @Summary 批量删除脆弱性管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除脆弱性管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /vulner/deleteVulner [delete]
export const deleteVulnerByIds = (params) => {
  return service({
    url: '/vulner/deleteVulnerByIds',
    method: 'delete',
    params
  })
}

// @Tags Vulner
// @Summary 更新脆弱性管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Vulner true "更新脆弱性管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /vulner/updateVulner [put]
export const updateVulner = (data) => {
  return service({
    url: '/vulner/updateVulner',
    method: 'put',
    data
  })
}

// @Tags Vulner
// @Summary 用id查询脆弱性管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Vulner true "用id查询脆弱性管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /vulner/findVulner [get]
export const findVulner = (params) => {
  return service({
    url: '/vulner/findVulner',
    method: 'get',
    params
  })
}

// @Tags Vulner
// @Summary 分页获取脆弱性管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取脆弱性管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /vulner/getVulnerList [get]
export const getVulnerList = (params) => {
  return service({
    url: '/vulner/getVulnerList',
    method: 'get',
    params
  })
}

// @Tags Vulner
// @Summary 不需要鉴权的脆弱性管理接口
// @accept application/json
// @Produce application/json
// @Param data query vulnerMReq.VulnerSearch true "分页获取脆弱性管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /vulner/getVulnerPublic [get]
export const getVulnerPublic = () => {
  return service({
    url: '/vulner/getVulnerPublic',
    method: 'get',
  })
}
