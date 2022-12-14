package migrations

import (
	"pxj/CloudTravelShopApi/go/models"
)

func init() {
	models.Db.AutoMigrate(
		&models.User{},
		&models.AdminUser{},
	)
}
