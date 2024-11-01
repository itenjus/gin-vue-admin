// 自动生成模板Vulner
package vulnerM
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
	"gorm.io/datatypes"
)

// 脆弱性管理 结构体  Vulner
type Vulner struct {
    global.GVA_MODEL
    Title  string `json:"title" form:"title" gorm:"column:title;comment:脆弱性名称;" binding:"required"`  //脆弱性名称 
    Desc  string `json:"desc" form:"desc" gorm:"column:desc;comment:脆弱性描述;"`  //脆弱性描述 
    Datereported  *time.Time `json:"datereported" form:"datereported" gorm:"column:datereported;comment:脆弱性上报日期;" binding:"required"`  //脆弱性上报日期 
    Finishdate  *time.Time `json:"finishdate" form:"finishdate" gorm:"column:finishdate;comment:脆弱性处理的结束时间;"`  //脆弱性处理的结束时间 
    Risklevel  string `json:"risklevel" form:"risklevel" gorm:"column:risklevel;comment:风险等级;" binding:"required"`  //风险等级 
    Vulnerstatus  string `json:"vulnerstatus" form:"vulnerstatus" gorm:"column:vulnerstatus;comment:脆弱性状态;" binding:"required"`  //脆弱性状态 
    Pic  datatypes.JSON `json:"pic" form:"pic" gorm:"column:pic;comment:脆弱性的图片;"swaggertype:"array,object"`  //脆弱性的图片 
    Source  string `json:"source" form:"source" gorm:"column:source;comment:脆弱性来源;"`  //脆弱性来源 
    Solution  string `json:"solution" form:"solution" gorm:"column:solution;comment:;"`  //解决方案 
    Processrecord  string `json:"processrecord" form:"processrecord" gorm:"column:processrecord;comment:;"`  //处置过程 
    Ops  string `json:"ops" form:"ops" gorm:"column:ops;comment:ISO27001安全域;"`  //安全域 
    Org  string `json:"org" form:"org" gorm:"column:org;comment:;"`  //org 
    Assignto  string `json:"assignto" form:"assignto" gorm:"column:assignto;comment:分配给;"`  //分配给 
    Assets  string `json:"assets" form:"assets" gorm:"column:assets;comment:;" binding:"required"`  //资产 
    EnTitle  string `json:"enTitle" form:"enTitle" gorm:"column:en_title;comment:;"`  //enTitle 
    EnDescription  string `json:"enDescription" form:"enDescription" gorm:"column:en_description;comment:;"`  //enDescription 
    EnSolution  string `json:"enSolution" form:"enSolution" gorm:"column:en_solution;comment:;"`  //enSolution 
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 脆弱性管理 Vulner自定义表名 vulner
func (Vulner) TableName() string {
    return "vulner"
}

