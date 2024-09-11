package database

import (
	"fmt"

	"github.com/amirmohammadkariimi/interview-task/internal/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ceate new gorm DB
func New(address, name, user, pass string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, address, name)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

// migrate models to DB
func Migrate(db *gorm.DB) {
	_ = db.AutoMigrate(&models.DNSQuery{})
	_ = db.AutoMigrate(&models.Address{})
}
