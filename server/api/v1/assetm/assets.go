package assetm

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/assetm"
    assetmReq "github.com/flipped-aurora/gin-vue-admin/server/model/assetm/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type AssetsApi struct {}



// CreateAssets 创建资产管理
// @Tags Assets
// @Summary 创建资产管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body assetm.Assets true "创建资产管理"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /AM/createAssets [post]
func (AMApi *AssetsApi) CreateAssets(c *gin.Context) {
	var AM assetm.Assets
	err := c.ShouldBindJSON(&AM)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    AM.CreatedBy = utils.GetUserID(c)
	err = AMService.CreateAssets(&AM)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteAssets 删除资产管理
// @Tags Assets
// @Summary 删除资产管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body assetm.Assets true "删除资产管理"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /AM/deleteAssets [delete]
func (AMApi *AssetsApi) DeleteAssets(c *gin.Context) {
	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := AMService.DeleteAssets(ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteAssetsByIds 批量删除资产管理
// @Tags Assets
// @Summary 批量删除资产管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /AM/deleteAssetsByIds [delete]
func (AMApi *AssetsApi) DeleteAssetsByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := AMService.DeleteAssetsByIds(IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateAssets 更新资产管理
// @Tags Assets
// @Summary 更新资产管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body assetm.Assets true "更新资产管理"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /AM/updateAssets [put]
func (AMApi *AssetsApi) UpdateAssets(c *gin.Context) {
	var AM assetm.Assets
	err := c.ShouldBindJSON(&AM)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    AM.UpdatedBy = utils.GetUserID(c)
	err = AMService.UpdateAssets(AM)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindAssets 用id查询资产管理
// @Tags Assets
// @Summary 用id查询资产管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query assetm.Assets true "用id查询资产管理"
// @Success 200 {object} response.Response{data=assetm.Assets,msg=string} "查询成功"
// @Router /AM/findAssets [get]
func (AMApi *AssetsApi) FindAssets(c *gin.Context) {
	ID := c.Query("ID")
	reAM, err := AMService.GetAssets(ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reAM, c)
}

// GetAssetsList 分页获取资产管理列表
// @Tags Assets
// @Summary 分页获取资产管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query assetmReq.AssetsSearch true "分页获取资产管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /AM/getAssetsList [get]
func (AMApi *AssetsApi) GetAssetsList(c *gin.Context) {
	var pageInfo assetmReq.AssetsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := AMService.GetAssetsInfoList(pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetAssetsPublic 不需要鉴权的资产管理接口
// @Tags Assets
// @Summary 不需要鉴权的资产管理接口
// @accept application/json
// @Produce application/json
// @Param data query assetmReq.AssetsSearch true "分页获取资产管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /AM/getAssetsPublic [get]
func (AMApi *AssetsApi) GetAssetsPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    AMService.GetAssetsPublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的资产管理接口信息",
    }, "获取成功", c)
}
