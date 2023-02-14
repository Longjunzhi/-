package test

import (
	"github.com/sirupsen/logrus"
	"pxj/CloudTravelShopApi/go/database"
	"pxj/CloudTravelShopApi/go/models"
	"testing"
)

func TestTemp(t *testing.T) {
	database.DbInit()
	adminUsers := make([]models.AdminUser, 0)
	err := database.Db.Model(&models.AdminUser{}).Where("id = ?", 1).Find(&adminUsers).Error
	if err != nil {
		t.Errorf("is err: %+v", err)
	}
	logrus.Printf("adminUsers %+#v", adminUsers)
	t.Fatal()
}
