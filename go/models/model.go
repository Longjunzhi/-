package models

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"pxj/CloudTravelShopApi/go/config"
	"strings"
)

var (
	Db *gorm.DB
)

func DbInit() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.AppConf.MysqlConf.User,
		config.AppConf.MysqlConf.Password,
		config.AppConf.MysqlConf.Host,
		config.AppConf.MysqlConf.Port,
		config.AppConf.MysqlConf.DB,
	)
	Db, _ = gorm.Open("mysql", dsn)
}
func IsErrorRecordNotFound(err error) bool {
	if strings.Contains(err.Error(), "record not found") {
		return true
	}
	return false
}

func IsErrorRecordDuplicated(err error) bool {
	if err, ok := err.(*mysql.MySQLError); ok && err.Number == 1062 {
		return true
	}
	return false
}
