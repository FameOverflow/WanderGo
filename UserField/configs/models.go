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
func (p *Address) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), p)
}
func (c Photo) Value() (driver.Value, error) {
	data, err := json.Marshal(c)
	return string(data), err
}
func (c *Photo) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

type User struct {
	gorm.Model
	UserName     string    `gorm:"not null" json:"user_name"`
	UserAccount  string    `gorm:"not null;index" json:"user_account"`
	UserPassword string    `gorm:"not null" json:"user_password"`
	UserCaptcha  int       `gorm:"-" json:"user_captcha"`
	Comments     []Comment `gorm:"foreignKey:UserUID" json:"comments"`
	Stars        []Star    `gorm:"foreignKey:UserUID"`
	Photos       []Photo   `gorm:"foreignKey:PhotoUID"`
}
type Photo struct {
	gorm.Model
	UserAccount string `gorm:"not null;index" json:"user_account"`
	PhotoData   []byte `gorm:"not null" json:"photo_data"`
	PhotoUID    string `gorm:"not null;index" json:"photo_id"`
	User        User   `gorm:"foreignKey:PhotoUID"`
}
type Avatar struct {
	gorm.Model
	UserAccount string `gorm:"not null;index" json:"user_account"`
	AvatarData  []byte `gorm:"not null" json:"avatar_data"`
	AvatarUID   string `gorm:"not null" json:"avatar_uid" binding:"required"`
}
type Comment struct {
	gorm.Model
	UserAccount string  `gorm:"not null;index" json:"user_account"`
	Date        string  `gorm:"not null" json:"date"`
	PhotoUID    string  `json:"photo_uid"`
	Text        string  `json:"text"`
	CommentUUID string  `gorm:"not null" json:"comment_uuid"`
	UserUID     uint    `gorm:"not null" json:"user_uid"`
	PlaceUID    uint    `gorm:"not null" json:"place_uid"`
	Position    Address `gorm:"-" json:"position"`
	User        User    `gorm:"foreignKey:UserUID"`
	Place       Place   `gorm:"foreignKey:PlaceUID"`
	StarCnt     int     `gorm:"not null" json:"star_cnt"`
	Stars       []Star  `gorm:"foreignKey:CommentUID"`
	PhotoData   []byte  `json:"photo_data"`
}
type Place struct {
	gorm.Model
	PlaceName        string    `gorm:"not null" json:"place_name"`
	TopLeftPoint     Address   `gorm:"TYPE:json" json:"top_left_point"`
	BottomRightPoint Address   `gorm:"TYPE:json" json:"bottom_right_point"`
	CenterPoint      Address   `gorm:"TYPE:json" json:"center_point"`
	Comments         []Comment `gorm:"foreignKey:PlaceUID" json:"comments"`
	IsMarked         bool      `json:"is_marked"`
}

// 点赞
type Star struct {
	gorm.Model
	UserAccount string  `json:"user_account"`
	CommentUUID string  `gorm:"not null" json:"comment_uuid"`
	UserUID     uint    `gorm:"not null" json:"user_uid"`
	CommentUID  uint    `gorm:"not null" json:"comment_uid"`
	Comment     Comment `gorm:"foreignKey:CommentUID"`
	User        User    `gorm:"foreignKey:UserUID"`
}
