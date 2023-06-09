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
	Feature                 string   `json:"feature" form:"feature" gorm:"column:feature;type:text;comment:生物特征;"`
	FaceId                  string   `json:"faceId" form:"faceId" gorm:"column:face_id;comment:人脸id;"`       // 人脸id
	PersonId                string   `json:"personId" form:"personId" gorm:"column:person_id;comment:人员id;"` // 人员id
	Image                   string   `json:"image" form:"image" gorm:"column:image;type:text;comment:应及影像;"`
	GroupId                 string   `json:"groupId" form:"groupId" gorm:"column:group_id;comment:人员库id;"`
	ContactPerson           string   `json:"contactPerson" form:"contactPerson" gorm:"column:contact_person;type:text;comment:紧急联系人;"`
	BloodType               string   `gorm:"column:blood_type;type:varchar(10);comment:血型" json:"bloodType" form:"bloodType"`                    // 血型，字符串类型，最大长度为10
	DrugAllergy             string   `gorm:"column:drug_allergy;type:text;comment:药物过敏" json:"drugAllergy" form:"drugAllergy"`                   // 药物过敏，字符串类型，文本格式
	CongenitalDisease       string   `gorm:"column:congenital_disease;type:text;comment:先天疾病" json:"congenitalDisease" form:"congenitalDisease"` // 先天疾病，字符串类型，文本格式
	OtherMedicalInformation string   `gorm:"column:other_medical_information;type:text;comment:其他医疗信息" json:"otherMedicalInformation"`           // 其它医疗信息，字符串类型，文本格式
	Status                  int      `gorm:"column:status;type:int;default:1;comment:状态 1安全，2危险" json:"status" form:"status"`                    // 状态，整型
}

// TableName Baby 表名
func (Baby) TableName() string {
	return "wechat_baby"
}
