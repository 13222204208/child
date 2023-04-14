package response

import (
	"encoding/json"
	"fmt"

	"github.com/13222204208/copilot"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	iai "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iai/v20200303"
)

// 创建人员库
func CreateGroup() (groupId string) {
	client := ConfigGroup()
	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := iai.NewCreateGroupRequest()
	str := copilot.UserNum()
	request.GroupName = &str
	request.GroupId = &str
	// 返回的resp是一个CreateGroupResponse的实例，与请求对象对应
	response, err := client.CreateGroup(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Println("err错误创建人员库", err)
		return
	}
	if err != nil {
		fmt.Println("err错误创建人员库", err)
		return
	} else {
		if *response.Response.RequestId != "" {
			groupId = str
			return
		}
	}
	return
}

type Person struct {
	PersonId   string `json:"person_id"`
	PersonName string `json:"person_name"`
	GroupId    string `json:"group_id"`
	Pic        string `json:"pic"`
}

// 创建人员
func CreatePerson(p *Person) (faceId string) {
	client := ConfigGroup()
	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := iai.NewCreatePersonRequest()
	request.GroupId = &p.GroupId
	request.PersonName = &p.PersonName
	request.PersonId = &p.PersonId
	pic := p.Pic
	picBase64, _ := copilot.PrefixImgBase64(pic)
	request.Image = &picBase64
	// 返回的resp是一个CreatePersonResponse的实例，与请求对象对应
	response, err := client.CreatePerson(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Println("err错误 创建人员", err)
		return
	}
	if err != nil {
		fmt.Println("err错误 创建人员", err)
	}
	if *response.Response.FaceId != "" {
		faceId = *response.Response.FaceId
		return
	}
	return
}

// 配置人员库
func ConfigGroup() *iai.Client {
	credential := common.NewCredential(
		global.GVA_CONFIG.TencentCOS.SecretID,
		global.GVA_CONFIG.TencentCOS.SecretKey,
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "iai.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := iai.NewClient(credential, "ap-beijing", cpf)
	return client
}

// 解析应及影像内容
type image struct {
	Img  string `json:"img"`
	Tips string `json:"tips"`
	Time string `json:"time"`
}

// 解析一个json 数组字符串
func ParseJsonArray(jsonStr string) (img string, err error) {
	var res []image
	err = json.Unmarshal([]byte(jsonStr), &res)
	img = res[0].Img
	return
}

// 保存faceId到baby表
func SaveFaceId(image string, babyId uint) (faceId string, err error) {
	pic := ""
	pic, err = ParseJsonArray(image)
	if err != nil {
		return
	}
	if pic != "" {
		var baby wechat.Baby
		err = global.GVA_DB.Where("id = ?", babyId).First(&baby).Error
		if err != nil {
			return
		}

		if baby.FaceId != "" {
			// err = DeleteFace(baby.FaceId, baby.PersonId)
			// fmt.Println("删除人脸的状态:", err)
			//获取添加的人脸id
			var newFaceId string
			if newFaceId, err = AddFace(pic, baby.PersonId); err != nil {
				return
			}
			baby.FaceId = newFaceId
			err = global.GVA_DB.Save(&baby).Error
			fmt.Println("保存人脸的状态:", err)
		} else {
			var p Person
			p.GroupId = baby.GroupId // 人员库id
			p.PersonName = baby.Name
			p.PersonId = baby.PersonId
			p.Pic = pic
			faceId = CreatePerson(&p)
		}
	}
	return
}

// 删除人脸
func DeleteFace(faceId, personId string) (err error) {
	client := ConfigGroup()
	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := iai.NewDeleteFaceRequest()
	request.FaceIds = []*string{&faceId}
	request.PersonId = &personId
	// 返回的resp是一个DeleteFaceResponse的实例，与请求对象对应
	_, err = client.DeleteFace(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("删除人脸错误: %s", err)
		return
	}
	if err != nil {
		return
	}
	return
}

// 人脸搜索
func SearchFace(image, person string) (personId string, err error) {
	client := ConfigGroup()
	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := iai.NewSearchFacesRequest()
	picBase64, _ := copilot.PrefixImgBase64(image)
	request.Image = &picBase64
	request.GroupIds = []*string{&person}
	// 返回的resp是一个SearchFacesResponse的实例，与请求对象对应
	response, err := client.SearchFaces(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("人脸搜索 : %s", err)
		return
	}
	if err != nil {
		return
	}
	result := response.Response.Results

	score := *result[0].Candidates[0].Score
	fmt.Println("人脸搜索结果:", score, *result[0].Candidates[0].PersonId)
	if score > 80 {
		// return *result[0].Candidates[0].FaceId, nil
		//返回personId
		return *result[0].Candidates[0].PersonId, nil
	}
	return
}

// 增加人脸
func AddFace(image, personId string) (faceId string, err error) {
	client := ConfigGroup()
	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := iai.NewCreateFaceRequest()
	base64Image, _ := copilot.PrefixImgBase64(image)
	request.PersonId = &personId
	request.Images = []*string{&base64Image}
	// 返回的resp是一个CreateFaceResponse的实例，与请求对象对应
	response, err := client.CreateFace(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("增加人脸错误: %s", err)
		return
	}
	if err != nil {
		return
	}
	faceId = *response.Response.SucFaceIds[0]
	fmt.Println("增加人脸的faceId:", faceId)
	return
}
