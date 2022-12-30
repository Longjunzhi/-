package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type AdminUser struct {
	Name          string `gorm:"default:''"`
	Mobile        string `gorm:"default:''"`
	Password      string `gorm:"default:''"`
	ParentId      int64  `gorm:"default:0"`
	LastLoginTime time.Time
	LoginIp       string
	gorm.Model
}

func NewAdminUser() (au *AdminUser) {
	return &AdminUser{}
}

func (adminUser *AdminUser) FirstOrCreate() (err error) {
	err = Db.Debug().Where(AdminUser{
		Name: adminUser.Name,
	}).Attrs(adminUser).FirstOrCreate(&adminUser).Error
	return err
}

func (adminUser *AdminUser) Create() (err error) {
	err = Db.Debug().Create(&adminUser).Error
	return err
}

func GetAdminUserModel() (model *gorm.DB, err error) {
	model = Db.Debug().Model(AdminUser{})
	err = model.Error
	return model, err
}

func GetAdminUserById(id uint) (adminUser *AdminUser, err error) {
	adminUser = &AdminUser{}
	err = Db.Debug().Where(map[string]uint{
		"id": id,
	}).First(&adminUser).Error
	return adminUser, err
}
