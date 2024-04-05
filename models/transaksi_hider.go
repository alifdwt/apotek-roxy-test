package models

import "time"

type TransaksiHider struct {
	IDTrans          string            `json:"id_trans" gorm:"primaryKey;size:15"`
	TglTrans         time.Time         `json:"tgl_trans"`
	Total            int               `json:"total"`
	TransaksiDetails []TransaksiDetail `json:"transaksi_details" gorm:"foreignKey:IDTrans"`
}

func (TransaksiHider) TableName() string {
	return "transaksi_hider"
}
