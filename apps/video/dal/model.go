package dal

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Video struct {
	Id            int64     `gorm:"primarykey"`
	Uid           int64     `gorm:"not null"`
	PlayUrl       string    `gorm:"not null"`
	CoverUrl      string    `gorm:"not null"`
	FavoriteCount int64     `gorm:"not null;default:0"`
	CommentCount  int64     `gorm:"not null;default:0"`
	Title         string    `gorm:"not null"`
	CreatedAt     time.Time `gorm:"index,sort:desc"`
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func NewGorm(dsn string) *gorm.DB {
	DB, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	if err := DB.AutoMigrate(&Video{}); err != nil {
		panic(err)
	}
	return DB
}
