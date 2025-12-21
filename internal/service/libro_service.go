package service

import (
	"errors"
	"prueba-tecnica-back/internal/models"

	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) ObtenTodosLosLibros() ([]*models.Libro, error) {
	var libros []*models.Libro
	err := s.db.Preload("Autor").Find(&libros).Error
	return libros, err
}

func (s *Service) ObtenerLibroPorID(id uint) (*models.Libro, error) {
	var libro models.Libro
	err := s.db.Preload("Autor").First(&libro, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("libro no encontrado")
		}
		return nil, err
	}
	return &libro, nil
}

func (s *Service) CrearLibro(libro *models.Libro) (*models.Libro, error) {
	if libro.Title == "" {
		return nil, errors.New("necesitamos el título")
	}

	err := s.db.Create(libro).Error
	if err != nil {
		return nil, err
	}

	var created models.Libro
	err = s.db.Preload("Autor").First(&created, libro.ID).Error
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func (s *Service) ActualizarLibro(id uint, libro *models.Libro) (*models.Libro, error) {
	if libro.Title == "" {
		return nil, errors.New("necesitamos el título")
	}

	if err := s.db.First(&models.Libro{}, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("libro no encontrado")
		}
		return nil, err
	}

	libro.ID = id
	if err := s.db.Save(libro).Error; err != nil {
		return nil, err
	}

	var updated models.Libro
	if err := s.db.Preload("Autor").First(&updated, id).Error; err != nil {
		return nil, err
	}

	return &updated, nil
}

func (s *Service) EliminarLibro(id uint) error {
	var libro models.Libro
	err := s.db.First(&libro, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("libro no encontrado")
		}
		return err
	}
	return s.db.Delete(&libro).Error
}
