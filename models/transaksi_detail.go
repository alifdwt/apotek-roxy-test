package models

type TransaksiDetail struct {
	ID             string         `json:"id" gorm:"primaryKey"`
	IDTrans        string         `json:"id_trans" gorm:"size:15"`
	IDBarang       string         `json:"id_barang" gorm:"size:15"`
	Qty            int            `json:"qty" gorm:"size:10"`
	TransaksiHider TransaksiHider `json:"transaksi_hider" gorm:"foreignKey:ID;references:IDTrans"`
}

func (TransaksiDetail) TableName() string {
	return "transaksi_detail"
}
