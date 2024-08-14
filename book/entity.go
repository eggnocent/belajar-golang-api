package book

import "time"

type Book struct {
	ID          int       `gorm:"type:int(11);primaryKey"`
	Title       string    `gorm:"type:varchar(255)"`
	Description string    `gorm:"type:varchar(255)"`
	Price       int       `gorm:"type:int(11)"`
	Rating      int       `gorm:"type:int(11)"`
	CreateAt    time.Time `gorm:"type:datetime"`
	UpdateAt    time.Time `gorm:"type:datetime"`
}
