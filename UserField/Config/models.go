package Config

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
)

type Address struct {
	X float64 `gorm:"not null" json:"x"`
	Y float64 `gorm:"not null" json:"y"`
}

func (p Address) Value() (driver.Value, error) {
	data, err := json.Marshal(p)
	return string(data), err
}
func (c *Address) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

type User struct {
	gorm.Model
	UserName     string    `gorm:"not null" json:"user_name"`
	UserAccount  string    `gorm:"not null;index" json:"user_account" binding:"required`
	UserPassword string    `gorm:"not null" json:"user_password" binding:"required`
	UserCaptcha  int       `gorm:"-" json:"user_captcha" binding:"required`
	Comments     []Comment `gorm:"foreignKey:CommentUID" json:"comments"`
	Stars        []Star    `gorm:"foreignKey:StarUID"`
}
type Photo struct {
	gorm.Model
	UserAccount string `gorm:"not null;index" json:"user_account" binding:"required`
	PhotoData   []byte `gorm:"not null" json:"photo_data"`
	PhotoID     string `gorm:"not null;index json:"photo_id binding:"required"`
}
type Avatar struct {
	gorm.Model
	UserAccount string `gorm:"not null;index" json:"user_account"`
	AvatarData  []byte `gorm:"not null" json:"photo_data"`
	AvatarID    string `gorm:"not null;index json:"photo_id binding:"required"`
}
type Comment struct {
	gorm.Model
	UserAccount string  `gorm:"not null;index" json:"user_account"`
	Data        string  `gorm:"not null" json:"data"`
	PhotoID     string  `json:"photo_id"`
	PhotoData   []byte  `json:"photo_data"`
	Sentence    string  `json:"sentence"`
	CommentUID  string  `gorm:"not null" json:"comment_id"`
	Position    Address `gorm:"-" json:"position"`
	User        User    `gorm:"foreignKey:CommentUID"`
	Place       Place   `gorm:"foreignKey:CommentUID"`
	StarCnt     int     `gorm:"default:0" json:"star_cnt"`
	Stars       []Star  `gorm:"foreignKey:StarUID"`
}
type Place struct {
	gorm.Model
	PlaceName        string    `gorm:"not null" json:"place_name"`
	PlaceUID         int       `gorm:"not null" json:"place_uid"`
	TopLeftPoint     Address   `gorm:"TYPE:json" json:"top_left_point"`
	BottomRightPoint Address   `gorm:"TYPE:json" json:"bottom_right_point"`
	CenterPoint      Address   `gorm:"TYPE:json" json:center_point"`
	Comments         []Comment `gorm:"foreignKey:CommentUID" json:"comments"`
}
type Star struct {
	gorm.Model
	UserAccount string  `json:"user_account"`
	CommentUID  string  `json:"comment_uid"`
	StarUID     string  `json:"star_uid"`
	Comment     Comment `gorm:"foreignKey:StarUID"`
	User        User    `gorm:"foreignKey:StarUID"`
}
