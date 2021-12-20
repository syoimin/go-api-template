package repositories

import (
	"github.com/syoimin/go-api-template/sample/models"
	"gorm.io/gorm"
)

type ISamplesRepository interface {
	FindAll() (*[]models.Sample, error)
}

type samplesRepository struct {
	*gorm.DB
}

func NewSamplesRepository(Conn *gorm.DB) ISamplesRepository {
	return &samplesRepository{Conn}
}

func (repo *samplesRepository) FindAll() (*[]models.Sample, error) {
	var m []models.Sample

	res := repo.DB.Debug().
		Find(&m)

	if res.Error != nil {
		return nil, ErrorHandler(res.Error)
	}

	return &m, nil
}
