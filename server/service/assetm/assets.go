package assetm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/assetm"
    assetmReq "github.com/flipped-aurora/gin-vue-admin/server/model/assetm/request"
    "gorm.io/gorm"
)

type AssetsService struct {}
// CreateAssets 创建资产管理记录
// Author [yourname](https://github.com/yourname)
func (AMService *AssetsService) CreateAssets(AM *assetm.Assets) (err error) {
	err = global.GVA_DB.Create(AM).Error
	return err
}

// DeleteAssets 删除资产管理记录
// Author [yourname](https://github.com/yourname)
func (AMService *AssetsService)DeleteAssets(ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&assetm.Assets{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&assetm.Assets{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteAssetsByIds 批量删除资产管理记录
// Author [yourname](https://github.com/yourname)
func (AMService *AssetsService)DeleteAssetsByIds(IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&assetm.Assets{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&assetm.Assets{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateAssets 更新资产管理记录
// Author [yourname](https://github.com/yourname)
func (AMService *AssetsService)UpdateAssets(AM assetm.Assets) (err error) {
	err = global.GVA_DB.Model(&assetm.Assets{}).Where("id = ?",AM.ID).Updates(&AM).Error
	return err
}

// GetAssets 根据ID获取资产管理记录
// Author [yourname](https://github.com/yourname)
func (AMService *AssetsService)GetAssets(ID string) (AM assetm.Assets, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&AM).Error
	return
}

// GetAssetsInfoList 分页获取资产管理记录
// Author [yourname](https://github.com/yourname)
func (AMService *AssetsService)GetAssetsInfoList(info assetmReq.AssetsSearch) (list []assetm.Assets, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&assetm.Assets{})
    var AMs []assetm.Assets
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&AMs).Error
	return  AMs, total, err
}
func (AMService *AssetsService)GetAssetsPublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
