package assetm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/assetm"
    assetmReq "github.com/flipped-aurora/gin-vue-admin/server/model/assetm/request"
)

type AssettypeService struct {}
// CreateAssettype 创建资产类型记录
// Author [yourname](https://github.com/yourname)
func (ATService *AssettypeService) CreateAssettype(AT *assetm.Assettype) (err error) {
	err = global.GVA_DB.Create(AT).Error
	return err
}

// DeleteAssettype 删除资产类型记录
// Author [yourname](https://github.com/yourname)
func (ATService *AssettypeService)DeleteAssettype(ID string) (err error) {
	err = global.GVA_DB.Delete(&assetm.Assettype{},"id = ?",ID).Error
	return err
}

// DeleteAssettypeByIds 批量删除资产类型记录
// Author [yourname](https://github.com/yourname)
func (ATService *AssettypeService)DeleteAssettypeByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]assetm.Assettype{},"id in ?",IDs).Error
	return err
}

// UpdateAssettype 更新资产类型记录
// Author [yourname](https://github.com/yourname)
func (ATService *AssettypeService)UpdateAssettype(AT assetm.Assettype) (err error) {
	err = global.GVA_DB.Model(&assetm.Assettype{}).Where("id = ?",AT.ID).Updates(&AT).Error
	return err
}

// GetAssettype 根据ID获取资产类型记录
// Author [yourname](https://github.com/yourname)
func (ATService *AssettypeService)GetAssettype(ID string) (AT assetm.Assettype, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&AT).Error
	return
}

// GetAssettypeInfoList 分页获取资产类型记录
// Author [yourname](https://github.com/yourname)
func (ATService *AssettypeService)GetAssettypeInfoList(info assetmReq.AssettypeSearch) (list []assetm.Assettype, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&assetm.Assettype{})
    var ATs []assetm.Assettype
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["type_name"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&ATs).Error
	return  ATs, total, err
}
func (ATService *AssettypeService)GetAssettypePublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
