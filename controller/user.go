package controller

import (
	"changeme/Dao"
	"changeme/utils"
)

type User struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
	Id          int    `json:"id"`
	Description string `json:"description"`
}

func (a *App) CheckLogin(userinfo User) *utils.Resp {
	userDao := Dao.UserDao{
		Username: userinfo.Username,
		Password: userinfo.Password,
	}
	ok, user := userDao.CheckLogin(DB)
	if ok {
		return utils.Success(user)
	}
	return utils.Error("用户名或密码错误")
}
