package service

import (
	"errors"
	"prueba-tecnica-back/internal/models"

	"gorm.io/gorm"
)

type PostService struct {
	db *gorm.DB
}

func NewPostService(db *gorm.DB) *PostService {
	return &PostService{db: db}
}

func (s *PostService) Create(userID uint, title, content string) (*models.Post, error) {
	post := &models.Post{
		Title:   title,
		Content: content,
		UserID:  userID,
	}

	if err := s.db.Create(post).Error; err != nil {
		return nil, err
	}

	if err := s.db.Preload("User").First(&post, post.ID).Error; err != nil {
		return nil, err
	}

	return post, nil
}

func (s *PostService) GetPosts(page, limit int, userID *uint, sortOrder string) ([]models.Post, int64, error) {
	var posts []models.Post
	var total int64

	db := s.db.Model(&models.Post{}).Preload("User")

	if userID != nil {
		db = db.Where("user_id = ?", *userID)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if sortOrder == "asc" {
		db = db.Order("created_at ASC")
	} else {
		db = db.Order("created_at DESC")
	}

	offset := (page - 1) * limit
	if err := db.Offset(offset).Limit(limit).Find(&posts).Error; err != nil {
		return nil, 0, err
	}

	return posts, total, nil
}

func (s *PostService) GetByID(id uint) (*models.Post, error) {
	var post models.Post
	if err := s.db.Preload("User").First(&post, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("post no encontrado")
		}
		return nil, err
	}
	return &post, nil
}

func (s *PostService) Update(id, userID uint, title, content string) (*models.Post, error) {
	var post models.Post
	if err := s.db.First(&post, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("post no encontrado")
		}
		return nil, err
	}

	if post.UserID != userID {
		return nil, errors.New("no autorizado")
	}

	post.Title = title
	post.Content = content

	if err := s.db.Save(&post).Error; err != nil {
		return nil, err
	}

	if err := s.db.Preload("User").First(&post, id).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func (s *PostService) Delete(id, userID uint) error {
	var post models.Post
	if err := s.db.First(&post, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("post no encontrado")
		}
		return err
	}

	if post.UserID != userID {
		return errors.New("no autorizado")
	}

	return s.db.Delete(&post).Error
}
