package dal

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id            int64  `gorm:"primarykey"`
	Name          string `gorm:"not null"`
	Password      string `json:"password"`
	FollowCount   int64  `gorm:"not null;default:0"`
	FollowerCount int64  `gorm:"not null;default:0"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func NewGorm(dsn string) *gorm.DB {
	DB, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	if err := DB.AutoMigrate(&User{}); err != nil {
		panic(err)
	}
	return DB
}
