package services

import (
	"github.com/syoimin/go-api-template/sample/models"
	"github.com/syoimin/go-api-template/sample/repositories"
)

type ISamplesService interface {
	ListSamples() (*[]models.Sample, error)
	// GetSamples(id uint) (*models.Sample, error)
	// CreateSamples(m *models.Sample) (*models.Sample, error)
	// UpdateSamples(id uint, m *models.Sample) (*models.Sample, error)
	// DeleteSamples(id uint) (*models.Sample, error)
}

type samplesService struct {
	repositories.ISamplesRepository
}

// コンストラクタ
func NewSamplesService(repo repositories.ISamplesRepository) ISamplesService {
	return &samplesService{repo}
}

func (srv *samplesService) ListSamples() (*[]models.Sample, error) {
	// 一覧取得（テーブル結合）
	r, err := srv.ISamplesRepository.FindAll()

	if err != nil {
		return nil, err
	}

	return r, nil
}
