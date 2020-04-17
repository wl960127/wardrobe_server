package noteservice

import (
	"wardrobe_server/database/operating"
)

// Note 接口使用参数封装
type Note struct {
	ID         int
	UserID     int
	Experience string // 心得 备注
	PicWhole   string //整体
	PicCoat    string //上衣
	PicSkirt   string //裙子
	PicPants   string //裤子
	PicShoes   string //鞋子
}

// AddNote 作为操作数据库的中转站
func (note *Note) AddNote() error {
	data := map[string]interface{}{
		"userid":     note.UserID,
		"experience": note.Experience,
		"whole":      note.PicWhole,
		"coat":       note.PicCoat,
		"skirt":      note.PicSkirt,
		"pants":      note.PicPants,
		"shoes":      note.PicShoes,
	}
	return operating.AddNote(data)
}
