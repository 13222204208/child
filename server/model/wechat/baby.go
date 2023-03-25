// 自动生成模板Baby
package wechat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Baby 结构体
type Baby struct {
	global.GVA_MODEL
	Uid                     uint     `json:"uid" form:"uid" gorm:"column:uid;comment:家长的id;"`
	Name                    string   `json:"name" form:"name" gorm:"column:name;comment:;"`
	Gender                  *int     `json:"gender" form:"gender" gorm:"column:gender;comment:;"`
	Birthday                string   `json:"birthday" form:"birthday" gorm:"column:birthday;comment:;"`
	Height                  *float64 `json:"height" form:"height" gorm:"column:height;comment:;"`
	Weight                  *float64 `json:"weight" form:"weight" gorm:"column:weight;comment:;"`
	Appearance              string   `json:"appearance" form:"appearance" gorm:"column:appearance;comment:体貌特征;"`
	BloodType               string   `gorm:"column:blood_type;type:varchar(10);comment:血型" json:"bloodType" form:"bloodType"`                    // 血型，字符串类型，最大长度为10
	DrugAllergy             string   `gorm:"column:drug_allergy;type:text;comment:药物过敏" json:"drugAllergy" form:"drugAllergy"`                   // 药物过敏，字符串类型，文本格式
	CongenitalDisease       string   `gorm:"column:congenital_disease;type:text;comment:先天疾病" json:"congenitalDisease" form:"congenitalDisease"` // 先天疾病，字符串类型，文本格式
	OtherMedicalInformation string   `gorm:"column:other_medical_information;type:text;comment:其他医疗信息" json:"otherMedicalInformation"`           // 其它医疗信息，字符串类型，文本格式
}

// TableName Baby 表名
func (Baby) TableName() string {
	return "wechat_baby"
}
