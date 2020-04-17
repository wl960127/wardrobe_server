package entity

// Note 每日穿搭.
type Note struct {
	AutoIncrementEntity
	UserID        int
	Experience string // 心得 备注
	PicWhole       string //整体
	PicCoat      string   //上衣
	PicSkirt       string  //裙子
	PicPants       string  //裤子
	PicShoes       string  //鞋子
}
