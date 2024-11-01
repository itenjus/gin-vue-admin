package vulnerM

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/vulnerM"
    vulnerMReq "github.com/flipped-aurora/gin-vue-admin/server/model/vulnerM/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type VulnerApi struct {}



// CreateVulner 创建脆弱性管理
// @Tags Vulner
// @Summary 创建脆弱性管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body vulnerM.Vulner true "创建脆弱性管理"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /vulner/createVulner [post]
func (vulnerApi *VulnerApi) CreateVulner(c *gin.Context) {
	var vulner vulnerM.Vulner
	err := c.ShouldBindJSON(&vulner)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    vulner.CreatedBy = utils.GetUserID(c)
	err = vulnerService.CreateVulner(&vulner)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteVulner 删除脆弱性管理
// @Tags Vulner
// @Summary 删除脆弱性管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body vulnerM.Vulner true "删除脆弱性管理"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /vulner/deleteVulner [delete]
func (vulnerApi *VulnerApi) DeleteVulner(c *gin.Context) {
	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := vulnerService.DeleteVulner(ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteVulnerByIds 批量删除脆弱性管理
// @Tags Vulner
// @Summary 批量删除脆弱性管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /vulner/deleteVulnerByIds [delete]
func (vulnerApi *VulnerApi) DeleteVulnerByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := vulnerService.DeleteVulnerByIds(IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateVulner 更新脆弱性管理
// @Tags Vulner
// @Summary 更新脆弱性管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body vulnerM.Vulner true "更新脆弱性管理"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /vulner/updateVulner [put]
func (vulnerApi *VulnerApi) UpdateVulner(c *gin.Context) {
	var vulner vulnerM.Vulner
	err := c.ShouldBindJSON(&vulner)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    vulner.UpdatedBy = utils.GetUserID(c)
	err = vulnerService.UpdateVulner(vulner)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindVulner 用id查询脆弱性管理
// @Tags Vulner
// @Summary 用id查询脆弱性管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query vulnerM.Vulner true "用id查询脆弱性管理"
// @Success 200 {object} response.Response{data=vulnerM.Vulner,msg=string} "查询成功"
// @Router /vulner/findVulner [get]
func (vulnerApi *VulnerApi) FindVulner(c *gin.Context) {
	ID := c.Query("ID")
	revulner, err := vulnerService.GetVulner(ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(revulner, c)
}

// GetVulnerList 分页获取脆弱性管理列表
// @Tags Vulner
// @Summary 分页获取脆弱性管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query vulnerMReq.VulnerSearch true "分页获取脆弱性管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /vulner/getVulnerList [get]
func (vulnerApi *VulnerApi) GetVulnerList(c *gin.Context) {
	var pageInfo vulnerMReq.VulnerSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := vulnerService.GetVulnerInfoList(pageInfo)
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

// GetVulnerPublic 不需要鉴权的脆弱性管理接口
// @Tags Vulner
// @Summary 不需要鉴权的脆弱性管理接口
// @accept application/json
// @Produce application/json
// @Param data query vulnerMReq.VulnerSearch true "分页获取脆弱性管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /vulner/getVulnerPublic [get]
func (vulnerApi *VulnerApi) GetVulnerPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    vulnerService.GetVulnerPublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的脆弱性管理接口信息",
    }, "获取成功", c)
}
