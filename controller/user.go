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

// Register 注册 目前的注册是本地的注册 尽可以用于用户的区分
// 以后可能需要开发服务器来提供非本地用户之间的交流通信
func (a *App) Register(userinfo User) *utils.Resp {
	userDao := Dao.UserDao{
		Username:    userinfo.Username,
		Password:    userinfo.Password,
		Email:       userinfo.Email,
		Avatar:      userinfo.Avatar,
		Description: userinfo.Description,
	}
	ok := userDao.Register(DB)
	if ok {
		return utils.Success(nil)
	}
	return utils.Error("注册失败")
}
