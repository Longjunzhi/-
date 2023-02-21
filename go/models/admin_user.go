package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type AdminUser struct {
	ID            uint      `json:"id"`
	Name          string    `gorm:"default:''" json:"name"`
	Mobile        string    `gorm:"default:''" json:"mobile"`
	Password      string    `gorm:"default:''" json:"password"`
	ParentId      uint8     `gorm:"default:0" json:"parent_id"`
	LastLoginTime time.Time `json:"last_login_time"`
	LoginIp       string    `json:"login_ip"`
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
	err = Db.Debug().First(&adminUser, "id=?", id).Error
	return adminUser, err
}
