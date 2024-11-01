package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
)

func holder(routers ...*gin.RouterGroup) {
	_ = routers
	_ = router.RouterGroupApp
}
func initBizRouter(routers ...*gin.RouterGroup) {
	privateGroup := routers[0]
	publicGroup := routers[1]
	holder(publicGroup, privateGroup)
	{
		vulnerMRouter := router.RouterGroupApp.VulnerM
		vulnerMRouter.InitVulnerRouter(privateGroup, publicGroup)
	}
	{
		assetmRouter := router.RouterGroupApp.Assetm
		assetmRouter.InitAssetsRouter(privateGroup, publicGroup)
		assetmRouter.InitAssettypeRouter(privateGroup, publicGroup)
	}
	{
		eventmRouter := router.RouterGroupApp.Eventm
		// 占位方法，保证文件可以正确加载，避免go空变量检测报错，请勿删除。
		eventmRouter.InitEventmtRouter(privateGroup, publicGroup)
	}
}
