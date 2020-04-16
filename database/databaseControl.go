package database



import (
	parse_config "wardrobe_server/pkg/app/parseConfig"
	"wardrobe_server/database/entity"
	"fmt"
	"wardrobe_server/pkg/logging"

	"github.com/jinzhu/gorm"

	// _ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


var db *gorm.DB
var err error

// GetDb 获取数据库实例
func GetDb()  *gorm.DB{
	return db
}


func init() {
	var str string

	// if string.Compare(model, "debug") {
	str = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
	parse_config.MysqlDataSetting.User,
	parse_config.MysqlDataSetting.Password,
	parse_config.MysqlDataSetting.Host,
	parse_config.MysqlDataSetting.Name)
	db, err = gorm.Open(parse_config.MysqlDataSetting.Type, str)

	// } else {
	// str = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
	// 	parse_config.PostgresDataSetting.Host,
	// 	parse_config.PostgresDataSetting.HTTPPort,
	// 	parse_config.PostgresDataSetting.User,
	// 	parse_config.PostgresDataSetting.Name,
	// 	parse_config.PostgresDataSetting.Password,
	// 	parse_config.PostgresDataSetting.SSLMode,
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
	// db.LogMode(true)

	// 建表
	db.AutoMigrate(new(entity.User), new(entity.Picture), new(entity.Note))
}

// CloseDB .
func CloseDB() {
	defer db.Close()
}
