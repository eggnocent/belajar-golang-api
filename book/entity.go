package book

type BookRequest struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"` // Menandakan kolom ID sebagai primary key dan auto increment
	Title       string `gorm:"type:varchar(255)"`
	Description string `gorm:"type:varchar(255)"`
	Price       int    `gorm:"type:int(11)"`
	Discount    int    `gorm:"type:int(11)"`
	Rating      int    `gorm:"type:int(11)"`
}
