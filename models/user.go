package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// QueryUser 查询数据 .
func QueryUser(mobile, password string) (bool, error) {
	var user User
	// User{Mobile: mobile, Password: password}
	// db.Select("id").Where(&User{Mobile: mobile, Password: password})
	if err := db.Where("Mobile=? and Password=?", mobile, password).First(&user).Error; err != nil {
		return false, nil
	}

	fmt.Printf(" 数据库查寻 %d %s %s ", user.UserID, user.Mobile, user.Password)

	if user.UserID > 0 {
		return true, nil
	}

	return false, nil
}

// GetUser .
func GetUser(claimsID int) (*User, error) {
	var user User
	err := db.Where(&User{UserID: claimsID}).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// AddUser 添加 .
func AddUser(data map[string]interface{}) error {
	err := db.Create(&User{
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
	var user User
	err := db.Where("mobile = ? ", mobile).First(&user).Error

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

// BeforeCreate .
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

// BeforeUpdate .
func (user *User) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
