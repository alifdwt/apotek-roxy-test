package transaksi

import (
	transaksidetail "apotek-roxy/domain/requests/transaksi_detail"

	"github.com/go-playground/validator/v10"
)

type CreateTransaksiRequest struct {
	IDTrans string                                         `json:"id_trans" validate:"required"`
	Detail  []transaksidetail.CreateTransaksiDetailRequest `json:"detail" validate:"required"`
}

func (c *CreateTransaksiRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}
