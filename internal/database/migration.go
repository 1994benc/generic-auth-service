package database

import (
	"1994benc/auth-service/internal/user"

	"gorm.io/gorm"
)

// Migrate our database and create bill table
func MigrateDB(db *gorm.DB) error {
	models := []interface{}{&user.UserModel{}}
	result := db.AutoMigrate(models...)
	return result
}
