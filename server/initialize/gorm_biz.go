package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/assetm"
	"github.com/flipped-aurora/gin-vue-admin/server/model/eventm"
	"github.com/flipped-aurora/gin-vue-admin/server/model/vulnerM"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(vulnerM.Vulner{}, assetm.Assettype{}, eventm.Eventmt{})
	if err != nil {
		return err
	}
	return nil
}
