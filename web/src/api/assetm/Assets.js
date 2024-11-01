import service from '@/utils/request'
// @Tags Assets
// @Summary 创建资产管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Assets true "创建资产管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /AM/createAssets [post]
export const createAssets = (data) => {
  return service({
    url: '/AM/createAssets',
    method: 'post',
    data
  })
}

// @Tags Assets
// @Summary 删除资产管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Assets true "删除资产管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /AM/deleteAssets [delete]
export const deleteAssets = (params) => {
  return service({
    url: '/AM/deleteAssets',
    method: 'delete',
    params
  })
}

// @Tags Assets
// @Summary 批量删除资产管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除资产管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /AM/deleteAssets [delete]
export const deleteAssetsByIds = (params) => {
  return service({
    url: '/AM/deleteAssetsByIds',
    method: 'delete',
    params
  })
}

// @Tags Assets
// @Summary 更新资产管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Assets true "更新资产管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /AM/updateAssets [put]
export const updateAssets = (data) => {
  return service({
    url: '/AM/updateAssets',
    method: 'put',
    data
  })
}

// @Tags Assets
// @Summary 用id查询资产管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Assets true "用id查询资产管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /AM/findAssets [get]
export const findAssets = (params) => {
  return service({
    url: '/AM/findAssets',
    method: 'get',
    params
  })
}

// @Tags Assets
// @Summary 分页获取资产管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取资产管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /AM/getAssetsList [get]
export const getAssetsList = (params) => {
  return service({
    url: '/AM/getAssetsList',
    method: 'get',
    params
  })
}

// @Tags Assets
// @Summary 不需要鉴权的资产管理接口
// @accept application/json
// @Produce application/json
// @Param data query assetmReq.AssetsSearch true "分页获取资产管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /AM/getAssetsPublic [get]
export const getAssetsPublic = () => {
  return service({
    url: '/AM/getAssetsPublic',
    method: 'get',
  })
}
