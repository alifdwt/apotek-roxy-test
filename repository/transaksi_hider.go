package repository

import (
	transaksihider "apotek-roxy/domain/requests/transaksi_hider"
	"apotek-roxy/models"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type transaksiHiderRepository struct {
	db *gorm.DB
}

func NewTransaksiHiderRepository(db *gorm.DB) *transaksiHiderRepository {
	return &transaksiHiderRepository{db: db}
}

func (r *transaksiHiderRepository) CreateTransaksiHider(request transaksihider.CreateTransaksiHiderRequest) (*models.TransaksiHider, error) {
	var transaksiHiderModel models.TransaksiHider

	db := r.db.Model(&transaksiHiderModel)

	// auto increment idtrans "TR-0001, TR-0002, TR-0003, ..."
	var count int64
	if err := db.Model(&models.TransaksiHider{}).Count(&count).Error; err != nil {
		return nil, err
	}

	nextID := fmt.Sprintf("TR-%04d", count+1)
	transaksiHiderModel.IDTrans = nextID
	transaksiHiderModel.Total = request.Total

	transaksiHiderModel.TglTrans = time.Now()

	result := db.Debug().Create(&transaksiHiderModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &transaksiHiderModel, nil
}

func (r *transaksiHiderRepository) GetTransaksiHiderAll() (*[]models.TransaksiHider, error) {
	var transaksiHiderModel []models.TransaksiHider

	db := r.db.Model(&transaksiHiderModel)

	result := db.Debug().Order("tgl_trans DESC").Find(&transaksiHiderModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &transaksiHiderModel, nil
}

func (r *transaksiHiderRepository) GetTransaksiHiderById(idTrans string) (*models.TransaksiHider, error) {
	var transaksiHiderModel models.TransaksiHider

	db := r.db.Model(&transaksiHiderModel)

	result := db.Debug().Where("id_trans = ?", idTrans).First(&transaksiHiderModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &transaksiHiderModel, nil
}

func (r *transaksiHiderRepository) UpdateTransaksiHider(idTrans string, request transaksihider.UpdateTransaksiHiderRequest) (*models.TransaksiHider, error) {
	var transaksiHiderModel models.TransaksiHider

	db := r.db.Model(&transaksiHiderModel)

	checkTransaksiHiderById := db.Debug().Where("id_trans = ?", idTrans).First(&transaksiHiderModel)
	if checkTransaksiHiderById.RowsAffected < 1 {
		return &transaksiHiderModel, errors.New("transaksi hider not found")
	}

	transaksiHiderModel.Total = request.Total
	transaksiHiderModel.TglTrans = time.Now()

	result := db.Debug().Save(&transaksiHiderModel)
	if result.Error != nil {
		return &transaksiHiderModel, result.Error
	}

	return &transaksiHiderModel, nil
}

func (r *transaksiHiderRepository) DeleteTransaksiHider(idTrans string) (*models.TransaksiHider, error) {
	var transaksiHiderModel models.TransaksiHider

	db := r.db.Model(&transaksiHiderModel)

	checkTransaksiHiderById := db.Debug().Where("id_trans = ?", idTrans).First(&transaksiHiderModel)
	if checkTransaksiHiderById.RowsAffected < 1 {
		return &transaksiHiderModel, errors.New("transaksi hider not found")
	}

	deleteTransaksiHider := db.Debug().Delete(&transaksiHiderModel)
	if deleteTransaksiHider.Error != nil {
		return &transaksiHiderModel, deleteTransaksiHider.Error
	}

	return &transaksiHiderModel, nil
}
