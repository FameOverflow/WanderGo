package Database

import (
	"gorm.io/gorm"
)

type Address struct {
	X float64
	Y float64
}
type User struct {
	gorm.Model
	UserName     string `gorm:"not null" json:"user_name"`
	UserAccount  string `gorm:"not null;index" json:"user_account" binding:"required`
	UserPassword string `gorm:"not null" json:"user_password" binding:"required`
	UserCaptcha  int    `gorm:"-" json:"user_captcha" binding:"required`
	Comments     []Comment
}
type Photo struct {
	gorm.Model
	UserAccount string `gorm:"not null;index" json:"user_account" binding:"required`
	PhotoData   []byte `gorm:"not null" json:"photo_data"`
	PhotoID     string `grom:"not null;index json:"photo_id binding:"required"`
}
type Avatar struct {
	gorm.Model
	UserAccount string `gorm:"not null;index" json:"user_account" binding:"required`
	AvatarData  []byte `gorm:"not null" json:"photo_data"`
	AvatarID    string `grom:"not null;index json:"photo_id binding:"required"`
}
type Comment struct {
	gorm.Model
	UserAccount string  `gorm:"not null;index" json:"user_account" binding:"required`
	PhotoID     string  `json:"photo_id`
	PhotoData   []byte  `json:"photo_data"`
	Sentence    string  `json:"sentence"`
	CommentID   string  `gorm:"not null" json:"comment_id"`
	Position    Address `gorm:"not null" json:"position" binding:"required"`
}
type Place struct {
	gorm.Model
	PlaceName        string  ``
	PlaceID          int     ``
	TopLeftPoint     Address ``
	BottomRightPoint Address ``
	CenterPoint      Address ``
	Comments         []Comment
}
