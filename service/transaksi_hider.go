package service

import (
	transaksihider "apotek-roxy/domain/requests/transaksi_hider"
	"apotek-roxy/models"
	"apotek-roxy/pkg/logger"
	"apotek-roxy/repository"
)

type transaksiHiderService struct {
	repository repository.TransaksiHiderRepository
	log        logger.Logger
}

func NewTransaksiHiderService(repository repository.TransaksiHiderRepository, log logger.Logger) *transaksiHiderService {
	return &transaksiHiderService{repository, log}
}

func (s *transaksiHiderService) CreateTransaksiHider(request transaksihider.CreateTransaksiHiderRequest) (*models.TransaksiHider, error) {
	res, err := s.repository.CreateTransaksiHider(request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *transaksiHiderService) GetTransaksiHiderAll() (*[]models.TransaksiHider, error) {
	res, err := s.repository.GetTransaksiHiderAll()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *transaksiHiderService) GetTransaksiHiderById(idTrans string) (*models.TransaksiHider, error) {
	res, err := s.repository.GetTransaksiHiderById(idTrans)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *transaksiHiderService) UpdateTransaksiHider(idTrans string, request transaksihider.UpdateTransaksiHiderRequest) (*models.TransaksiHider, error) {
	res, err := s.repository.UpdateTransaksiHider(idTrans, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *transaksiHiderService) DeleteTransaksiHider(idTrans string) (*models.TransaksiHider, error) {
	res, err := s.repository.DeleteTransaksiHider(idTrans)
	if err != nil {
		return nil, err
	}

	return res, nil
}
