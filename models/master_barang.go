package models

type MasterBarang struct {
	IDBarang string `json:"id_barang" gorm:"primaryKey"`
	NmBarang string `json:"nm_barang" gorm:"size:30"`
	Qty      int    `json:"qty" gorm:"size:10"`
	Harga    string `json:"harga" gorm:"type:text"`
}

func (MasterBarang) TableName() string {
	return "master_barang"
}
