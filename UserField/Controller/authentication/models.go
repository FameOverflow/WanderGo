package Authentication

type PwdToBeChanged struct {
	OldPwd     string `json:"old_pwd" binding:"required"`
	CurrentPwd string `json:"current_pwd" binding:"required"`
}
type TempUser struct {
	UserAccount  string `json:"user_account" binding:"required"`
	UserPassword string `json:"user_password" binding:"required"`
}
type NameToBeChanged struct {
	UserName string `json:"user_name" binding:"required"`
}
type UserForgottenPre struct {
	UserAccount string `json:"user_account" binding:"required"`
}
type UserForgotten struct {
	UserAccount string `json:"user_account" binding:"required"`
	NewPwd      string `json:"new_pwd" binding:"required"`
	UserCaptcha int    `json:"user_captcha" binding:"required"`
}
type AccctStatus struct {
	Account            string
	TimeOfChangingName int64
}
type CommentsPayload struct {
	UserAccount string `gorm:"not null;index" json:"user_account"`
	Date        string `gorm:"not null" json:"date"`
	Text        string `json:"text"`
	CommentUUID string `gorm:"not null" json:"comment_uuid"`
	PlaceUID    uint   `gorm:"not null" json:"place_uid"`
	StarCnt     int    `gorm:"not null" json:"star_cnt"`
}

type UserCaptcha struct {
	Captcha int `json:"captcha"`
	UserAccount string `json:"user_account"`
	ExpireTime int64 `json:"expire_time"`
}