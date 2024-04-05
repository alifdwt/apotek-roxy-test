package service

import (
	"apotek-roxy/pkg/logger"
	"apotek-roxy/repository"
)

type Service struct {
	MasterBarang    MasterBarangService
	TransaksiHider  TransaksiHiderService
	TransaksiDetail TransaksiDetailService
}

type Deps struct {
	Repository *repository.Repositories
	Logger     logger.Logger
}

func NewService(deps Deps) *Service {
	return &Service{
		MasterBarang:    NewMasterBarangService(deps.Repository.MasterBarang, deps.Logger),
		TransaksiHider:  NewTransaksiHiderService(deps.Repository.TransaksiHider, deps.Logger),
		TransaksiDetail: NewTransaksiDetailService(deps.Repository.TransaksiDetail, deps.Logger),
	}
}
