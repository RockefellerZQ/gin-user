package vo

import (
	"crypto/md5"
	"fmt"
	"gin-user/model"
)

type UserRequest struct {
	Email string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required,min=6"`
}

func (u *UserRequest) ToUserModel() *model.User {
	return &model.User{
		Email: u.Email,
		Password: u.MD5Password(),
	}
}

func (u *UserRequest) MD5Password() string {
	bytes := md5.Sum([]byte(u.Password))
	return fmt.Sprintf("%x", bytes)
}

