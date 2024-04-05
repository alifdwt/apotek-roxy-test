package transaksidetail

import "github.com/go-playground/validator/v10"

type UpdateTransaksiDetailRequest struct {
	IDBarang string `json:"id_barang" validate:"required"`
	Qty      int    `json:"qty" validate:"required"`
}

func (c *UpdateTransaksiDetailRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}
