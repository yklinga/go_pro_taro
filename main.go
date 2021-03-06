package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go_pro_taro/route"
	"go_pro_taro/utils"
	"os"
)



func main() {
	InitConfig()
	utils.BaseDB()
	//defer db.Close() todo

	r := gin.Default()

	r = route.RouterCollect(r)

	port := viper.GetString("server.port")

	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("err")
	}
}

