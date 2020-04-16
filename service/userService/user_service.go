package userservice

import (
	"wardrobe_server/database/entity"
	"wardrobe_server/database/operating"
	"errors"
	"fmt"
	"wardrobe_server/pkg/utils"

)

// User .
type User struct {
	ID       int
	Username string
	Mobile   string
	Password string
	// UserID int

	CreatedBy  string
	ModifiedBy string
}

// Check .
func (a *User) Check() (bool, map[string]interface{},error) {
	fmt.Printf(" 准备数据库操作 %s  %s  ", a.Mobile, a.Password)
	return operating.QueryUser(a.Mobile, utils.EncodeMD5(a.Password))
}

// Get .
func (a *User) Get() (*entity.User, error) {
	user, err := operating.GetUser(a.ID)
	if err != nil {
		return nil, err
	}
	return user, nil

}

// Add 操作数据库.
func (a *User) Add() error {
	// var err error
	data := map[string]interface{}{
		"username": a.Username,
		"mobile":   a.Mobile,
		"password": utils.EncodeMD5(a.Password),
	}

	isHas, err := operating.CheckUser(a.Mobile, a.Username)

	if isHas {
		return errors.New("该账号已经存在")
	}
	if err := operating.AddUser(data); err == nil {
		return nil
	}
	return err

}
