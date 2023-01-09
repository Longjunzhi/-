package test

import (
	"pxj/CloudTravelShopApi/go/models"
	"testing"
)

func TempGorm(t *testing.T) {
	adminUser, err := models.GetAdminUserById(1)
	if err != nil {
		t.Errorf("err: %+v", err)
	}
	t.Fatal(adminUser)
}
