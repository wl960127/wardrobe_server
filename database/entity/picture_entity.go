package entity

import (
	"encoding/json"
)

// Picture .
type Picture struct {
	// MD5    string `gorm:"not null;unique"`
	// URL    string `gorm:"not null;unique"`
	// AbsolutePath    string `gorm:"not null;unique"`
	AutoIncrementEntity
	UserID       int    //`json:"-"`
	MD5          string `gorm:"not nul;" json:"-"`
	URL          string `gorm:"not nul;" json:"-"`
	AbsolutePath string `gorm:"not nul;" json:"-"`
	BRAND        string `json:"brand"`  // 品牌
	COLOR        string `json:"color"`  //颜色
	LABLE        string `json:"lable"`  // 备注
	TYPE         int    `json:"type"`   // 上衣之类
	SEASON       int    `json:"season"` // 季节  0 默认
	Count        int    `json:"count"`  // 调用次数 每次使用就 +1
	Size         int64  `json:"-"`      // 图片大小
}

// MarshalJSON .
func (p *Picture) MarshalJSON() ([]byte, error) {
	type picture Picture
	pic := struct {
		UpdatedTime int64  `json:"time"`
		PicURL      string `json:"picUrl"`
		*picture
	}{
		p.UpdatedAt.Unix(),
		"http://192.168.5.28:8000/" + p.URL,
		(*picture)(p),
	}
	return json.Marshal(pic)
}
