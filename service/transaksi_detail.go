package service

import (
	transaksidetail "apotek-roxy/domain/requests/transaksi_detail"
	"apotek-roxy/models"
	"apotek-roxy/pkg/logger"
	"apotek-roxy/repository"
)

type transaksiDetailService struct {
	repository repository.TransaksiDetailRepository
	log        logger.Logger
}

func NewTransaksiDetailService(repository repository.TransaksiDetailRepository, log logger.Logger) *transaksiDetailService {
	return &transaksiDetailService{repository, log}
}

func (s *transaksiDetailService) CreateTransaksiDetail(request transaksidetail.CreateTransaksiDetailRequest) (*models.TransaksiDetail, error) {
	res, err := s.repository.CreateTransaksiDetail(request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *transaksiDetailService) GetTransaksiDetailAll() (*[]models.TransaksiDetail, error) {
	res, err := s.repository.GetTransaksiDetailAll()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *transaksiDetailService) GetTransaksiDetailById(idTrans string) (*models.TransaksiDetail, error) {
	res, err := s.repository.GetTransaksiDetailById(idTrans)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *transaksiDetailService) UpdateTransaksiDetail(idTrans string, request transaksidetail.UpdateTransaksiDetailRequest) (*models.TransaksiDetail, error) {
	res, err := s.repository.UpdateTransaksiDetail(idTrans, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *transaksiDetailService) DeleteTransaksiDetail(idTrans string) (*models.TransaksiDetail, error) {
	res, err := s.repository.DeleteTransaksiDetail(idTrans)
	if err != nil {
		return nil, err
	}

	return res, nil
}
