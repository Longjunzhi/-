package test

import (
	"fmt"
	"pxj/CloudTravelShopApi/go/models"
	"testing"
)

func TestTemp(t *testing.T) {
	fmt.Print("test temp")
	models.DbInit()
	adminUser, err := models.GetAdminUserById(1)
	if err != nil {
		t.Errorf("err: %+v", err)
	}
	t.Fatal(adminUser)
}
