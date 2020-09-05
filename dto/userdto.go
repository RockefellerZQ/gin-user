package dto

import "gin-user/model"

type UserDTO struct {
	Email string `json:"email"`
}

func UserDTOFromUserModel(user model.User) UserDTO {
	return UserDTO{Email: user.Email}
}
