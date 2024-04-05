package repository

import "gorm.io/gorm"

type Repositories struct {
	MasterBarang    MasterBarangRepository
	TransaksiHider  TransaksiHiderRepository
	TransaksiDetail TransaksiDetailRepository
}

func NewRepository(db *gorm.DB) *Repositories {
	return &Repositories{
		MasterBarang:    NewMasterBarangRepository(db),
		TransaksiHider:  NewTransaksiHiderRepository(db),
		TransaksiDetail: NewTransaksiDetailRepository(db),
	}
}
