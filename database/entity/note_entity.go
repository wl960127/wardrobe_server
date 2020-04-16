package entity

// Note 每日穿搭.
type Note struct {
	AutoIncrementEntity
	UserID        int
	Experience string // 心得 备注
	pic0       string
	pic1       string
	pic2       string
	pic3       string
	pic4       string
}
