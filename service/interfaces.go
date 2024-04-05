package service

import (
	masterbarang "apotek-roxy/domain/requests/master_barang"
	transaksidetail "apotek-roxy/domain/requests/transaksi_detail"
	transaksihider "apotek-roxy/domain/requests/transaksi_hider"
	"apotek-roxy/models"
)

type MasterBarangService interface {
	CreateMasterBarang(request masterbarang.CreateMasterBarangRequest) (*models.MasterBarang, error)
	GetMasterBarangAll() (*[]models.MasterBarang, error)
	GetMasterBarangById(idBarang string) (*models.MasterBarang, error)
	GetMasterBarangByNmBarang(nmBarang string) (*[]models.MasterBarang, error)
	UpdateMasterBarang(idBarang string, request masterbarang.UpdateMasterBarangRequest) (*models.MasterBarang, error)
	DeleteMasterBarang(idBarang string) (*models.MasterBarang, error)
}

type TransaksiHiderService interface {
	CreateTransaksiHider(request transaksihider.CreateTransaksiHiderRequest) (*models.TransaksiHider, error)
	GetTransaksiHiderAll() (*[]models.TransaksiHider, error)
	GetTransaksiHiderById(idTrans string) (*models.TransaksiHider, error)
	UpdateTransaksiHider(idTrans string, request transaksihider.UpdateTransaksiHiderRequest) (*models.TransaksiHider, error)
	DeleteTransaksiHider(idTrans string) (*models.TransaksiHider, error)
}

type TransaksiDetailService interface {
	CreateTransaksiDetail(request transaksidetail.CreateTransaksiDetailRequest) (*models.TransaksiDetail, error)
	GetTransaksiDetailAll() (*[]models.TransaksiDetail, error)
	GetTransaksiDetailById(idTrans string) (*models.TransaksiDetail, error)
	UpdateTransaksiDetail(idTrans string, request transaksidetail.UpdateTransaksiDetailRequest) (*models.TransaksiDetail, error)
	DeleteTransaksiDetail(idTrans string) (*models.TransaksiDetail, error)
}
