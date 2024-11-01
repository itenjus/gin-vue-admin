package vulnerM

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ VulnerRouter }

var vulnerApi = api.ApiGroupApp.VulnerMApiGroup.VulnerApi
