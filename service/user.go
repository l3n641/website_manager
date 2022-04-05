package service

import (
	"errors"
	"gorm.io/gorm"
	"website_manager/database/sql"
	"website_manager/pkg/e"
	"website_manager/pkg/util"
	"website_manager/serializer"
)

type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=30"`
}

func (u *UserLoginService) UserLogin() serializer.Response {
	var user sql.User
	code := e.SUCCESS

	if result := sql.DB.Where("user_name = ?", u.UserName).First(&user); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			code = e.ERROR_NOT_EXIST_USER
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
	}

	state := user.CheckPassword(u.Password)
	if !state {
		code = e.ERROR_NOT_COMPARE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	token, err := util.GenerateToken(u.UserName, u.Password, 0)
	if err != nil {
		code = e.ERROR_AUTH_TOKEN
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Data:   serializer.TokenData{User: serializer.User{ID: user.ID, UserName: user.UserName}, Token: token},
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
