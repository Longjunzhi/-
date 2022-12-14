package models

import "github.com/jinzhu/gorm"

const AdminUserTableName = "admin_users"

type AdminUser struct {
	gorm.Model
	ID       int64
	Name     string `gorm:"default:''"`
	Mobile   string `gorm:"default:''"`
	Password string `gorm:"default:''"`
	ParentId int64  `gorm:"default:0"`
}

func NewAdminUser() (au *AdminUser) {
	return &AdminUser{}
}

func (au *AdminUser) TableName() string {
	return AdminUserTableName
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
