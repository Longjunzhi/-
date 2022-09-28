package migrations

import (
	"pxj/CloudTravelShopApi/go/models"
	"pxj/courseSystem/api/databases"
)

func init() {
	databases.Db.AutoMigrate(
		&models.User{},
	)
}
