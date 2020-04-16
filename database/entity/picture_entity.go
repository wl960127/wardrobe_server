package entity

// Picture .
type Picture struct {
	// MD5    string `gorm:"not null;unique"`
	// URL    string `gorm:"not null;unique"`
	// AbsolutePath    string `gorm:"not null;unique"`
	AutoIncrementEntity
	UserID        int
	MD5          string `gorm:"unique"`
	URL          string `gorm:"unique"`
	AbsolutePath string `gorm:"unique"`
	BRAND        string // 品牌
	COLOR        string //颜色
	LABLE        string // 备注
	TYPE         int    // 上衣之类
	SEASON       int    // 季节  0 默认
	Count        int    // 调用次数 每次使用就 +1
	Size         int64  // 图片大小
}
