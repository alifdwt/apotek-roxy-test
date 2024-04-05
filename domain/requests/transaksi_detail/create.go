package transaksidetail

import "github.com/go-playground/validator/v10"

type CreateTransaksiDetailRequest struct {
	IDTrans  string `json:"id_trans" validate:"required"`
	IDBarang string `json:"id_barang" validate:"required"`
	Qty      int    `json:"qty" validate:"required"`
}

func (c *CreateTransaksiDetailRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}
