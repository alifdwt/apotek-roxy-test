package service

import (
	masterbarang "apotek-roxy/domain/requests/master_barang"
	"apotek-roxy/models"
	"apotek-roxy/pkg/logger"
	"apotek-roxy/repository"
)

type masterBarangService struct {
	repository repository.MasterBarangRepository
	log        logger.Logger
}

func NewMasterBarangService(repository repository.MasterBarangRepository, log logger.Logger) *masterBarangService {
	return &masterBarangService{repository, log}
}

func (s *masterBarangService) CreateMasterBarang(request masterbarang.CreateMasterBarangRequest) (*models.MasterBarang, error) {
	res, err := s.repository.CreateMasterBarang(request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *masterBarangService) GetMasterBarangAll() (*[]models.MasterBarang, error) {
	res, err := s.repository.GetMasterBarangAll()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *masterBarangService) GetMasterBarangById(idBarang string) (*models.MasterBarang, error) {
	res, err := s.repository.GetMasterBarangById(idBarang)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *masterBarangService) GetMasterBarangByNmBarang(nmBarang string) (*[]models.MasterBarang, error) {
	res, err := s.repository.GetMasterBarangByNmBarang(nmBarang)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *masterBarangService) UpdateMasterBarang(idBarang string, request masterbarang.UpdateMasterBarangRequest) (*models.MasterBarang, error) {
	res, err := s.repository.UpdateMasterBarang(idBarang, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *masterBarangService) DeleteMasterBarang(idBarang string) (*models.MasterBarang, error) {
	res, err := s.repository.DeleteMasterBarang(idBarang)
	if err != nil {
		return nil, err
	}

	return res, nil
}
