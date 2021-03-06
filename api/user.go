package api

import (
	"github.com/gin-gonic/gin"
	"go_pro_taro/dto"
	"go_pro_taro/model"
	"go_pro_taro/response"
	"go_pro_taro/utils"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Register (ctx * gin.Context) {
	DB := utils.GetDB()
	var reqUser = model.User{}

	ctx.Bind(&reqUser)
	// 获取参数 username telephone password
	username := reqUser.Username
	telephone := reqUser.Telephone
	password := reqUser.Password
	// 数据验证
	if len(telephone) != 11 {
		response.ResCommon(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.ResCommon(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	//如果没有传 生成一个随机10位用户名
	if len(username) == 0 {
		username = utils.RandomUserName(10)
	}
	log.Println(username, telephone, password)

	// 判断手机号是否存在

	if utils.IsTelephoneExist(DB, telephone) {
		response.ResCommon(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已经存在")
		return
	}

	// 创建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.ResCommon(ctx, http.StatusInternalServerError, 500, nil, "用户已经存在")
		return
	}

	newUser := model.User{
		Username: username,
		Password: string(hasedPassword),
		Telephone: telephone,
	}
	DB.Create(&newUser)
	// 发放token
	token, err := utils.ReleaseToken(newUser)
	if err != nil {
		response.ResCommon(ctx, http.StatusInternalServerError, 500, nil, "系统异常")

		log.Panicf("token generate error: %v", err)
		return
	}
	//返回结果
	response.Success(ctx, gin.H{
		"token": token,
	}, "注册成功")
}

func Login (ctx * gin.Context) {
	db := utils.GetDB()
	var reqUser = model.User{}

	ctx.Bind(&reqUser)
	// 获取参数
	telephone := reqUser.Telephone
	password := reqUser.Password
	// 数据验证
	if len(telephone) != 11 {
		response.ResCommon(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.ResCommon(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	// 手机号是否存在
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.ResCommon(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")

		return
	}
	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.ResCommon(ctx, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}
	// 发放token
	token, err := utils.ReleaseToken(user)
	if err != nil {
		response.ResCommon(ctx, http.StatusInternalServerError, 500, nil, "系统异常")

		log.Panicf("token generate error: %v", err)
		return
	}

	// 返回结果
	response.Success(ctx, gin.H{
		"token": token,
	}, "登录成功")
}

func Userinfo(ctx * gin.Context)  {
	user, _ := ctx.Get("user")

	response.Success(ctx, gin.H{"user": dto.ToUserDto(user.(model.User))}, "")
}