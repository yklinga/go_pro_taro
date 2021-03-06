package utils

import (
	"go_pro_taro/model"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

// 生成随机 username 注册使用
func RandomUserName(n int) string{
	var letters = []byte("asdfghjklzxcvbnmqwertyuiopASDFGHJKLZXCVBNMQWERTYUIOP")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		if (i == 0) {
			result[i] = 'U'
		} else {
			result[i] = letters[rand.Intn(len(letters))]
		}
	}
	return string(result)
}

// 判断是否存在手机号码 注册使用
func IsTelephoneExist(db * gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
