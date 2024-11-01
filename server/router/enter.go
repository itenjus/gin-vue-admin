package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/assetm"
	"github.com/flipped-aurora/gin-vue-admin/server/router/eventm"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/vulnerM"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	VulnerM vulnerM.RouterGroup
	Assetm  assetm.RouterGroup
	Eventm  eventm.RouterGroup
}
