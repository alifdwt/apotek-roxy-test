package masterbarang

import "github.com/go-playground/validator/v10"

type CreateMasterBarangRequest struct {
	// IDBarang string `json:"id_barang" validate:"required"`
	NmBarang string `json:"nm_barang" validate:"required"`
	Qty      int    `json:"qty" validate:"required"`
	Harga    string `json:"harga" validate:"required"`
}

func (c *CreateMasterBarangRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}
