package operating



import (
	"wardrobe_server/database/entity"
	"wardrobe_server/database"
	"fmt"

	"github.com/jinzhu/gorm"
)

// QueryUser 查询数据 .
func QueryUser(mobile, password string) (bool,map[string]interface{}, error) {
	var user entity.User
	// User{Mobile: mobile, Password: password}
	// db.Select("id").Where(&User{Mobile: mobile, Password: password})
	if err := database.GetDb().Where("Mobile=? and Password=?", mobile, password).First(&user).Error; err != nil {
		return false,nil, nil
	}

	fmt.Printf(" 数据库查寻 %d %s %s ", user.UserID, user.Mobile, user.Password)

	if user.UserID > 0 {
		return true,map[string]interface{}{
			"userId":user.UserID,
		}, nil
	}

	return false,nil, nil
}

// GetUser .
func GetUser(claimsID int) (*entity.User, error) {
	var user entity.User
	err := database.GetDb().Where(&entity.User{UserID: claimsID}).First(&user).Error
	if err != nil {
		return nil, err
	}
	fmt.Printf(" 数据库 传过来的参数  %d  user数据 %d %s  &user的数据   %d ",claimsID,user.UserID,user.Mobile,&user.UserID)
	return &user, nil
}

// AddUser 添加 .
func AddUser(data map[string]interface{}) error {
	err := database.GetDb().Create(&entity.User{
		Mobile:    data["mobile"].(string),
		Password:  data["password"].(string),
		Username:  data["username"].(string),
	}).Error

	if nil != err {
		return err
	}
	return nil

	// isNull := db.NewRecord()
	// return !isNull

}

// CheckUser .
func CheckUser(mobile, username string) (bool, error) {
	var user entity.User
	err := database.GetDb().Where("mobile = ? ", mobile).First(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.UserID > 0 {
		return true, nil
	}
	return false, nil
}

// UpdateUser .
func UpdateUser() {

}

