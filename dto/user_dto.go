package dto

import "go_pro_taro/model"

type UserDto struct {
	Username string `json: "username"`
	Telephone string `json: "telephone"`
}

func ToUserDto(user model.User) UserDto  {
	return UserDto{
		Username: user.Username,
		Telephone: user.Telephone,
	}
}