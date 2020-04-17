package operating

import (
	"wardrobe_server/database"
	"wardrobe_server/database/entity"
)

// AddNote 直接数据库操作.
func AddNote(data map[string]interface{}) error {

	if err := database.GetDb().Create(&entity.Note{
		UserID:     data["userid"].(int),
		Experience: data["experience"].(string),
		PicWhole:   data["whole"].(string),
		PicCoat:    data["coat"].(string),
		PicSkirt:   data["skirt"].(string),
		PicPants:   data["pants"].(string),
		PicShoes:   data["shoes"].(string),
	}).Error; err != nil {
		return err
	}
	return nil

}
