package services

import (
	"stars/core"
	"stars/datasource"
	"stars/models"
)

type StarsService interface {
	GetAll() []models.StarInfo
	GetById(id int) *models.StarInfo
	DeleteById(id int) error
	UpdateByUserInfo(user *models.StarInfo, columns []string) error
	Create(user *models.StarInfo) error

	Search(country string) []models.StarInfo
}

type starsService struct {
	core *core.StarsCore
}

func NewStarsService() StarsService {
	return &starsService{
		core: core.NewStarsCore(datasource.InstanceMaster()),
	}
}

func (s *starsService) GetAll() []models.StarInfo {
	return s.core.GetAll()
}

func (s *starsService) GetById(id int) *models.StarInfo {

	return s.core.GetById(id)
}

func (s *starsService) DeleteById(id int) error {
	return s.core.DeleteById(id)
}

func (s *starsService) UpdateByUserInfo(user *models.StarInfo, columns []string) error {
	return s.core.UpdateByUserInfo(user, columns)
}

func (s *starsService) Create(user *models.StarInfo) error {
	return s.core.Create(user)
}

func (s *starsService) Search(country string) []models.StarInfo {
	return s.core.Search(country)
}
