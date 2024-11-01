package vulnerM

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type VulnerRouter struct {}

// InitVulnerRouter 初始化 脆弱性管理 路由信息
func (s *VulnerRouter) InitVulnerRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	vulnerRouter := Router.Group("vulner").Use(middleware.OperationRecord())
	vulnerRouterWithoutRecord := Router.Group("vulner")
	vulnerRouterWithoutAuth := PublicRouter.Group("vulner")
	{
		vulnerRouter.POST("createVulner", vulnerApi.CreateVulner)   // 新建脆弱性管理
		vulnerRouter.DELETE("deleteVulner", vulnerApi.DeleteVulner) // 删除脆弱性管理
		vulnerRouter.DELETE("deleteVulnerByIds", vulnerApi.DeleteVulnerByIds) // 批量删除脆弱性管理
		vulnerRouter.PUT("updateVulner", vulnerApi.UpdateVulner)    // 更新脆弱性管理
	}
	{
		vulnerRouterWithoutRecord.GET("findVulner", vulnerApi.FindVulner)        // 根据ID获取脆弱性管理
		vulnerRouterWithoutRecord.GET("getVulnerList", vulnerApi.GetVulnerList)  // 获取脆弱性管理列表
	}
	{
	    vulnerRouterWithoutAuth.GET("getVulnerPublic", vulnerApi.GetVulnerPublic)  // 脆弱性管理开放接口
	}
}
