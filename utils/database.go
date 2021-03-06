package utils

import (
	"github.com/spf13/viper"
	"go_pro_taro/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func BaseDB() *gorm.DB  {
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	charset := viper.GetString("datasource.charset")
	parseTime := viper.GetString("datasource.parseTime")

	dsn := username + ":" + password + "@tcp("+ host+":"+port+")/" + database + "?charset=" + charset + "&parseTime=" + parseTime
	//"root:Root.123@tcp(103.46.128.49:52706)/gin_database?charset=utf8&parseTime=True"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		//panic("dsn:"+ dsn)
		panic("failed to connect database, err" + err.Error())
	}
	db.AutoMigrate(&model.User{})

	DB = db
	return db
}

func GetDB() * gorm.DB  {
	return DB
}