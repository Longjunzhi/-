package test

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"pxj/CloudTravelShopApi/go/models"
	"testing"
)

func TestGorm(t *testing.T) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		"root",
		"119129",
		"175.178.218.166",
		3306,
		"pxjCourseSystem",
	)
	fmt.Printf("dsn: %v", dsn)
	db, _ := gorm.Open("mysql", dsn)

	db.AutoMigrate(
		&models.User{},
		&models.AdminUser{},
	)
	// Create
	db.Create(&models.User{Name: "D42", Password: "100"})

	// Read
	var user models.User
	db.First(&user, 1)                 // 根据整型主键查找
	db.First(&user, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	db.Model(&user).Update("Password", 200)
	// Update - 更新多个字段
	db.Model(&user).Updates(models.User{Name: "200", Password: "F42"}) // 仅更新非零值字段
	db.Model(&user).Updates(map[string]interface{}{"Name": 200, "Password": "F42"})

	// Delete - 删除 product
	db.Delete(&user, 1)
}
