package models

type Libro struct {
	ID          uint    `gorm:"primaryKey;autoIncrement"`
	Title       string  `gorm:"not null"`
	Description string  `gorm:"not null"`
	Price       float64 `gorm:"not null"`
	AuthorID    uint    `gorm:"not null"`
	Autor       Autor   `gorm:"foreignKey:AuthorID"`
}
