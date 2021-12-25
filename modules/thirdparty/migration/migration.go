package migration

import (
	"go-hexagonal-auth/modules/billing"
	"go-hexagonal-auth/modules/log"
	"go-hexagonal-auth/modules/user"
	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(user.User{}, billing.Billing{}, log.Log{})
}
