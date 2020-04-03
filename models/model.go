package models

import (
	"fmt"
	"wardrobe_server/pkg/logging"
	"wardrobe_server/pkg/setting"

	"github.com/jinzhu/gorm"

	// _ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User id自增  手机号要唯一 .
type User struct {
	UserID   int `gorm:"primary_key;AUTO_INCREMENT"`
	Username string
	Mobile   string `gorm:"not null;unique" json:"mobile"`
	Password string `gorm:"not null" json:"password"`
	Sex      int
	Picture  []Picture
	Note     []Note
}
// Picture .
type Picture struct {
	gorm.Model
	// MD5    string `gorm:"not null;unique"`
	// URL    string `gorm:"not null;unique"`
	// AbsolutePath    string `gorm:"not null;unique"`
	UserID        int
	MD5          string `gorm:"unique"`
	URL          string `gorm:"unique"`
	AbsolutePath string `gorm:"unique"`
	BRAND        string // 品牌
	COLOR        string //颜色
	LABLE        string // 备注
	TYPE         int // 上衣之类
	SEASON       int    // 季节  0 默认
	Count        int    // 调用次数 每次使用就 +1
	Size         int64  // 图片大小
}

// Note 每日穿搭.
type Note struct {
	UserID        int
	gorm.Model
	Experience string // 心得 备注
	pic0       string
	pic1       string
	pic2       string
	pic3       string
	pic4       string
}

// type

var db *gorm.DB
var err error

func init() {
	var str string

	// if string.Compare(model, "debug") {
	str = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.MysqlDataSetting.User,
		setting.MysqlDataSetting.Password,
		setting.MysqlDataSetting.Host,
		setting.MysqlDataSetting.Name)
	db, err = gorm.Open(setting.MysqlDataSetting.Type, str)

	// } else {
	// str = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
	// 	setting.PostgresDataSetting.Host,
	// 	setting.PostgresDataSetting.HTTPPort,
	// 	setting.PostgresDataSetting.User,
	// 	setting.PostgresDataSetting.Name,
	// 	setting.PostgresDataSetting.Password,
	// 	setting.PostgresDataSetting.SSLMode,
	// )
	// db, err := gorm.Open(setting.PostgresDataSetting.Type, str)
	// }

	fmt.Println("配置 " + str)
	if err != nil {
		logging.Error(err)
	} else {
		db.SingularTable(true)
		db.LogMode(true)
		db.DB().SetMaxIdleConns(10)
		db.DB().SetMaxOpenConns(100)
		logging.Info("数据库初始化")
		fmt.Println("数据库初始化")

	}
	// 建表
	db.AutoMigrate(&User{}, &Picture{}, &Note{})
}

// CloseDB .
func CloseDB() {
	defer db.Close()
}
