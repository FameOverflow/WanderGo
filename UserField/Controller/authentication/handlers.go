package Authentication

import (
	con "SparkForge/configs"
	util "SparkForge/util"
)

func AccountConflictVerification(a string) error { //有错说明不冲突
	var tUser con.User
	err := con.GLOBAL_DB.Model(&con.User{}).Where("user_account = ?", a).First(&tUser).Error
	return err
}
func UserLoginVerification(u con.User) (int, error) {
	var tUser con.User
	err := con.GLOBAL_DB.Model(&con.User{}).Where(con.User{UserAccount: u.UserAccount}, "user_account").
		First(&tUser).Error
	if err != nil {
		// 输入账号不存在
		return 1, err
	} else {
		// 若账号存在，检测密码是否正确
		err = con.GLOBAL_DB.Model(&con.User{}).Where(con.User{UserAccount: u.UserAccount, UserPassword: util.EncryptMd5(u.UserPassword)}, "user_account", "user_password").
			First(&tUser).Error
		if err != nil {
			// 密码不正确
			return 2, err
		}
		// 登录成功，没有错误
		return 0, nil
	}
}
