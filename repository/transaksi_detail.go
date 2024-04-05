package repository

import (
	transaksidetail "apotek-roxy/domain/requests/transaksi_detail"
	"apotek-roxy/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type transaksiDetailRepository struct {
	db *gorm.DB
}

func NewTransaksiDetailRepository(db *gorm.DB) *transaksiDetailRepository {
	return &transaksiDetailRepository{db: db}
}

func (r *transaksiDetailRepository) CreateTransaksiDetail(request transaksidetail.CreateTransaksiDetailRequest) (*models.TransaksiDetail, error) {
	var transaksiDetailModel models.TransaksiDetail

	db := r.db.Model(&transaksiDetailModel)

	var count int64
	if err := db.Model(&models.TransaksiDetail{}).Count(&count).Error; err != nil {
		return nil, err
	}

	transaksiDetailModel.IDTrans = request.IDTrans
	transaksiDetailModel.IDBarang = request.IDBarang
	transaksiDetailModel.Qty = request.Qty
	transaksiDetailModel.ID = fmt.Sprintf("TR-%04d", count+1)

	result := db.Debug().Create(&transaksiDetailModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &transaksiDetailModel, nil
}

func (r *transaksiDetailRepository) GetTransaksiDetailAll() (*[]models.TransaksiDetail, error) {
	var transaksiDetailModel []models.TransaksiDetail

	db := r.db.Model(&transaksiDetailModel)

	result := db.Debug().Find(&transaksiDetailModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &transaksiDetailModel, nil
}

func (r *transaksiDetailRepository) GetTransaksiDetailById(idTrans string) (*models.TransaksiDetail, error) {
	var transaksiDetailModel models.TransaksiDetail

	db := r.db.Model(&transaksiDetailModel)

	result := db.Debug().Where("id_trans = ?", idTrans).Find(&transaksiDetailModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &transaksiDetailModel, nil
}

func (r *transaksiDetailRepository) UpdateTransaksiDetail(idTrans string, request transaksidetail.UpdateTransaksiDetailRequest) (*models.TransaksiDetail, error) {
	var transaksiDetailModel models.TransaksiDetail

	db := r.db.Model(&transaksiDetailModel)

	checkTransaksiDetailById := db.Debug().Where("id_trans = ?", idTrans).First(&transaksiDetailModel)
	if checkTransaksiDetailById.RowsAffected < 1 {
		return &transaksiDetailModel, errors.New("transaksi detail not found")
	}

	transaksiDetailModel.IDBarang = request.IDBarang
	transaksiDetailModel.Qty = request.Qty

	result := db.Debug().Save(&transaksiDetailModel)
	if result.Error != nil {
		return &transaksiDetailModel, result.Error
	}

	return &transaksiDetailModel, nil
}

func (r *transaksiDetailRepository) DeleteTransaksiDetail(idTrans string) (*models.TransaksiDetail, error) {
	var transaksiDetailModel models.TransaksiDetail

	db := r.db.Model(&transaksiDetailModel)

	checkTransaksiDetailById := db.Debug().Where("id_trans = ?", idTrans).First(&transaksiDetailModel)
	if checkTransaksiDetailById.RowsAffected < 1 {
		return &transaksiDetailModel, errors.New("transaksi detail not found")
	}

	deleteTransaksiDetail := db.Debug().Delete(&transaksiDetailModel)
	if deleteTransaksiDetail.Error != nil {
		return &transaksiDetailModel, deleteTransaksiDetail.Error
	}

	return &transaksiDetailModel, nil
}
