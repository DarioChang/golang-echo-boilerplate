package post

import (
	"echo-demo-project/models"
	"echo-demo-project/requests"

	"github.com/jinzhu/gorm"
)

type ServiceWrapper interface {
	Create(post *models.Post)
	Delete(post *models.Post)
	Update(post *models.Post, updatePostRequest *requests.UpdatePostRequest)
}

type Service struct {
	Db *gorm.DB
}

func NewPostService(db *gorm.DB) ServiceWrapper {
	return &Service{Db: db}
}
