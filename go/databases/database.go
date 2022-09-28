package databases

import "github.com/jinzhu/gorm"

var (
	Db *gorm.DB
)

func init() {
	Db, _ = gorm.Open("mysql", "")
}

func getDb() *gorm.DB {
	return Db
}
