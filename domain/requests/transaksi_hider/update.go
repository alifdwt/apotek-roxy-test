package transaksihider

import (
	"github.com/go-playground/validator/v10"
)

type UpdateTransaksiHiderRequest struct {
	Total int `json:"total" validate:"required"`
}

func (c *UpdateTransaksiHiderRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}
