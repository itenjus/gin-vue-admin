package assetm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AssettypeRouter struct {}

// InitAssettypeRouter 初始化 资产类型 路由信息
func (s *AssettypeRouter) InitAssettypeRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	ATRouter := Router.Group("AT").Use(middleware.OperationRecord())
	ATRouterWithoutRecord := Router.Group("AT")
	ATRouterWithoutAuth := PublicRouter.Group("AT")
	{
		ATRouter.POST("createAssettype", ATApi.CreateAssettype)   // 新建资产类型
		ATRouter.DELETE("deleteAssettype", ATApi.DeleteAssettype) // 删除资产类型
		ATRouter.DELETE("deleteAssettypeByIds", ATApi.DeleteAssettypeByIds) // 批量删除资产类型
		ATRouter.PUT("updateAssettype", ATApi.UpdateAssettype)    // 更新资产类型
	}
	{
		ATRouterWithoutRecord.GET("findAssettype", ATApi.FindAssettype)        // 根据ID获取资产类型
		ATRouterWithoutRecord.GET("getAssettypeList", ATApi.GetAssettypeList)  // 获取资产类型列表
	}
	{
	    ATRouterWithoutAuth.GET("getAssettypePublic", ATApi.GetAssettypePublic)  // 资产类型开放接口
	}
}
