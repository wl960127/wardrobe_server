package operating


import (
	"fmt"
	"wardrobe_server/database"
	"wardrobe_server/database/entity"
	"github.com/jinzhu/gorm"
)

// Pic 图片 MD5 计算唯一值  查看是否存在 不存在就保存.
// type Pic struct {
// 	MD5          string
// 	URL          string
// 	AbsolutePath string
// 	BRAND        string // 品牌
// 	COLOR        string //颜色
// 	LABLE        string // 备注
// 	TYPE         string // 上衣之类
// 	SEASON       int    // 季节  0 默认
// 	Count        int    // 调用次数 每次使用就 +1
// 	Size         int    // 图片大小
// }

//AddPic

// CheckImageMD5 检查是否文件存在 md5 可以验证唯一.
func CheckImageMD5(md5Value string) (isExit bool,url, absolutePath string) {
	var picFile entity.Picture
	if err := database.GetDb().Where("md5 = ?", md5Value).First(&picFile).Error; err != gorm.ErrRecordNotFound {
		fmt.Printf("\n picture.go  %s %s ",picFile.URL,picFile.AbsolutePath);	
		return true,picFile.URL,picFile.AbsolutePath
	}
	return false,"",""
}

// AddPic 添加图片
func AddPic(data map[string]interface{}) error {

	if err := database.GetDb().Create(&entity.Picture{
		UserID:       data["userid"].(int),
		MD5:          data["md5"].(string),
		URL:          data["url"].(string),
		AbsolutePath: data["absolutePath"].(string),
		BRAND:        data["brand"].(string),
		COLOR:        data["color"].(string),
		LABLE:        data["lable"].(string),
		TYPE:         data["type"].(int),
		SEASON:       data["season"].(int),
		Size:         data["size"].(int64),
	}).Error; err != nil {
		return err
	}

	return nil
}

// QueryPic 查询所有类别
func QueryPic(userID int) ([]entity.Picture, error) {
	// db.Model()

	var pic []entity.Picture
	var err error
	if err = database.GetDb().Where("user_id = ?", userID).Find(&pic).Error; err != nil {
		return nil, err
	}

	return pic, err
}

// QueryByType 按照类别查询
func QueryByType() {

}

// UpdatePic 图片信息修改
func UpdatePic() {

}
