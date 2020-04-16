package entity

// User id自增  手机号要唯一 .
type User struct {
	BaseModel
	UserID   int `gorm:"primary_key;AUTO_INCREMENT"`
	Username string
	Mobile   string `gorm:"not null;unique" json:"mobile"`
	Password string `gorm:"not null" json:"password"`
	Sex      int
	Picture  []Picture
	Note     []Note
}