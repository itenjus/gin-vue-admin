package vulnerM

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ VulnerApi }

var vulnerService = service.ServiceGroupApp.VulnerMServiceGroup.VulnerService
