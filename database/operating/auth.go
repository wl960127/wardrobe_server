package operating

import (
	"wardrobe_server/database"
)




// Auth .
type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

// CheckAuth .
func CheckAuth(mobile, password string) bool {
	var auth Auth
	database.GetDb().Select("id").Where(Auth{Mobile: mobile, Password: password}).First(&auth)
	if auth.ID > 0 {
		return true
	}

	return false
}
