package migrations

import (
	"pxj/CloudTravelShopApi/go/databases"
	"pxj/CloudTravelShopApi/go/models"
)

func init() {
	databases.Db.AutoMigrate(
		&models.User{},
	)
}
