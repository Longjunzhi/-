package test

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"pxj/CloudTravelShopApi/go/models"
	"testing"
)

func TestGorm(t *testing.T) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		"root",
		"119129",
		"175.178.218.166",
		3306,
		"cloudTravelShop",
	)
	fmt.Printf("dsn: %v", dsn)
	db, _ := gorm.Open("mysql", dsn)

	db.AutoMigrate(
		&models.User{},
		&models.AdminUser{},
	)
	// Create
	password, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.MinCost)
	if err != nil {
		return
	}
	db.Create(&models.AdminUser{Name: "root", Password: string(password)})

	// Read
	//var adminUser models.AdminUser
	//db.First(&adminUser, 1)                 // 根据整型主键查找
	//db.First(&adminUser, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 adminUser 的 Password 更新为 200
	//db.Model(&adminUser).Update("Password", 200)
	// Update - 更新多个字段
	//db.Model(&adminUser).Updates(models.User{Name: "200", Password: "F42"}) // 仅更新非零值字段
	//db.Model(&adminUser).Updates(map[string]interface{}{"Name": root, "Password": "123456"})

	// Delete - 删除 adminUser
	//db.Delete(&adminUser, 1)
}
