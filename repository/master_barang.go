package repository

import (
	masterbarang "apotek-roxy/domain/requests/master_barang"
	"apotek-roxy/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type masterBarangRepository struct {
	db *gorm.DB
}

func NewMasterBarangRepository(db *gorm.DB) *masterBarangRepository {
	return &masterBarangRepository{db: db}
}

func (r *masterBarangRepository) CreateMasterBarang(request masterbarang.CreateMasterBarangRequest) (*models.MasterBarang, error) {
	var masterBarangModel models.MasterBarang

	db := r.db.Model(&masterBarangModel)

	// auto increment idbarang "BR-0001, BR-0002, BR-0003, ..."
	var count int64
	if err := db.Model(&models.MasterBarang{}).Count(&count).Error; err != nil {
		return nil, err
	}

	nextID := fmt.Sprintf("BR-%04d", count+1)
	masterBarangModel.IDBarang = nextID
	masterBarangModel.NmBarang = request.NmBarang
	masterBarangModel.Qty = request.Qty
	masterBarangModel.Harga = request.Harga

	result := db.Debug().Create(&masterBarangModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &masterBarangModel, nil
}

func (r *masterBarangRepository) GetMasterBarangAll() (*[]models.MasterBarang, error) {
	var masterBarangModel []models.MasterBarang

	db := r.db.Model(&masterBarangModel)

	result := db.Debug().Find(&masterBarangModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &masterBarangModel, nil
}

func (r *masterBarangRepository) GetMasterBarangById(idBarang string) (*models.MasterBarang, error) {
	var masterBarangModel models.MasterBarang

	db := r.db.Model(&masterBarangModel)

	result := db.Debug().Where("id_barang = ?", idBarang).First(&masterBarangModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &masterBarangModel, nil
}

func (r *masterBarangRepository) GetMasterBarangByNmBarang(nmBarang string) (*[]models.MasterBarang, error) {
	var masterBarangModel []models.MasterBarang

	db := r.db.Model(&masterBarangModel)

	result := db.Debug().Where("nm_barang = ?", nmBarang).First(&masterBarangModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &masterBarangModel, nil
}

func (r *masterBarangRepository) UpdateMasterBarang(idBarang string, request masterbarang.UpdateMasterBarangRequest) (*models.MasterBarang, error) {
	var masterBarangModel models.MasterBarang

	db := r.db.Model(&masterBarangModel)

	checkMasterBarangById := db.Debug().Where("id_barang = ?", idBarang).First(&masterBarangModel)
	if checkMasterBarangById.RowsAffected < 1 {
		return &masterBarangModel, errors.New("master barang not found")
	}

	masterBarangModel.NmBarang = request.NmBarang
	masterBarangModel.Qty = request.Qty
	masterBarangModel.Harga = request.Harga

	result := db.Debug().Save(&masterBarangModel)
	if result.Error != nil {
		return &masterBarangModel, result.Error
	}

	return &masterBarangModel, nil
}

func (r *masterBarangRepository) DeleteMasterBarang(idBarang string) (*models.MasterBarang, error) {
	var masterBarangModel models.MasterBarang

	db := r.db.Model(&masterBarangModel)

	checkMasterBarangById := db.Debug().Where("id_barang = ?", idBarang).First(&masterBarangModel)
	if checkMasterBarangById.RowsAffected < 1 {
		return &masterBarangModel, errors.New("master barang not found")
	}

	deleteMasterBarang := db.Debug().Delete(&masterBarangModel)
	if deleteMasterBarang.Error != nil {
		return &masterBarangModel, deleteMasterBarang.Error
	}

	return &masterBarangModel, nil
}
