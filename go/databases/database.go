package databases

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"pxj/CloudTravelShopApi/go/config"
)

var (
	Db *gorm.DB
)

func init() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.AppConf.MysqlConf.User,
		config.AppConf.MysqlConf.Password,
		config.AppConf.MysqlConf.Host,
		config.AppConf.MysqlConf.Port,
		config.AppConf.MysqlConf.DB,
	)
	Db, _ = gorm.Open("mysql", dsn)
}
