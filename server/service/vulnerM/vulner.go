package vulnerM

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/vulnerM"
    vulnerMReq "github.com/flipped-aurora/gin-vue-admin/server/model/vulnerM/request"
    "gorm.io/gorm"
)

type VulnerService struct {}
// CreateVulner 创建脆弱性管理记录
// Author [yourname](https://github.com/yourname)
func (vulnerService *VulnerService) CreateVulner(vulner *vulnerM.Vulner) (err error) {
	err = global.GVA_DB.Create(vulner).Error
	return err
}

// DeleteVulner 删除脆弱性管理记录
// Author [yourname](https://github.com/yourname)
func (vulnerService *VulnerService)DeleteVulner(ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&vulnerM.Vulner{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&vulnerM.Vulner{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteVulnerByIds 批量删除脆弱性管理记录
// Author [yourname](https://github.com/yourname)
func (vulnerService *VulnerService)DeleteVulnerByIds(IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&vulnerM.Vulner{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&vulnerM.Vulner{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateVulner 更新脆弱性管理记录
// Author [yourname](https://github.com/yourname)
func (vulnerService *VulnerService)UpdateVulner(vulner vulnerM.Vulner) (err error) {
	err = global.GVA_DB.Model(&vulnerM.Vulner{}).Where("id = ?",vulner.ID).Updates(&vulner).Error
	return err
}

// GetVulner 根据ID获取脆弱性管理记录
// Author [yourname](https://github.com/yourname)
func (vulnerService *VulnerService)GetVulner(ID string) (vulner vulnerM.Vulner, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&vulner).Error
	return
}

// GetVulnerInfoList 分页获取脆弱性管理记录
// Author [yourname](https://github.com/yourname)
func (vulnerService *VulnerService)GetVulnerInfoList(info vulnerMReq.VulnerSearch) (list []vulnerM.Vulner, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&vulnerM.Vulner{})
    var vulners []vulnerM.Vulner
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Title != "" {
        db = db.Where("title LIKE ?","%"+ info.Title+"%")
    }
        if info.StartDatereported != nil && info.EndDatereported != nil {
            db = db.Where("datereported BETWEEN ? AND ? ",info.StartDatereported,info.EndDatereported)
        }
        if info.StartFinishdate != nil && info.EndFinishdate != nil {
            db = db.Where("finishdate BETWEEN ? AND ? ",info.StartFinishdate,info.EndFinishdate)
        }
    if info.Risklevel != "" {
        db = db.Where("risklevel = ?",info.Risklevel)
    }
    if info.Vulnerstatus != "" {
        db = db.Where("vulnerstatus = ?",info.Vulnerstatus)
    }
    if info.Source != "" {
        db = db.Where("source LIKE ?","%"+ info.Source+"%")
    }
    if info.Ops != "" {
        db = db.Where("ops LIKE ?","%"+ info.Ops+"%")
    }
    if info.Org != "" {
        db = db.Where("org LIKE ?","%"+ info.Org+"%")
    }
    if info.Assignto != "" {
        db = db.Where("assignto LIKE ?","%"+ info.Assignto+"%")
    }
    if info.Assets != "" {
        db = db.Where("assets LIKE ?","%"+ info.Assets+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&vulners).Error
	return  vulners, total, err
}
func (vulnerService *VulnerService)GetVulnerPublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
