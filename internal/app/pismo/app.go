package pismo

import "gorm.io/gorm"

// App holds shared dependencies
type App struct {
	DB *gorm.DB
}
