package models

type Autor struct {
	ID     uint    `gorm:"primaryKey;autoIncrement"`
	Name   string  `gorm:"not null"`
	Email  string  `gorm:"not null;unique"`
	Libros []Libro `gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
