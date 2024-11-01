package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/assetm"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/eventm"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/vulnerM"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	VulnerMApiGroup vulnerM.ApiGroup
	AssetmApiGroup  assetm.ApiGroup
	EventmApiGroup  eventm.ApiGroup
}
