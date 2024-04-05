package transaksihider

import (
	"github.com/go-playground/validator/v10"
)

type CreateTransaksiHiderRequest struct {
	Total int `json:"total" validate:"required"`
}

func (c *CreateTransaksiHiderRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}
