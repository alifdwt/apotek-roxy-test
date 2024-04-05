package handler

import (
	transaksidetail "apotek-roxy/domain/requests/transaksi_detail"
	transaksihider "apotek-roxy/domain/requests/transaksi_hider"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initTransaksiGroup(api *gin.Engine) {
	transaksi := api.Group("/transaksi")

	transaksi.POST("/", h.handlerCreateTransaksi)
}

func (h *Handler) handlerCreateTransaksi(c *gin.Context) {
	var body transaksidetail.CreateTransaksiDetailRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := body.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	barang, err := h.services.MasterBarang.GetMasterBarangById(body.IDBarang)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("failed to get master barang: %w", err)))
		return
	}

	harga, _ := strconv.Atoi(barang.Harga)
	total := harga * body.Qty

	transaksiHider := transaksihider.CreateTransaksiHiderRequest{
		Total: total,
	}

	header, err := h.services.TransaksiHider.CreateTransaksiHider(transaksiHider)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("failed to create transaksi: %w", err)))
		return
	}

	detailArg := transaksidetail.CreateTransaksiDetailRequest{
		IDTrans:  header.IDTrans,
		IDBarang: body.IDBarang,
		Qty:      body.Qty,
	}

	_, err = h.services.TransaksiDetail.CreateTransaksiDetail(detailArg)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("failed to create transaksi detail: %w", err)))
		return
	}

	resp, err := h.services.TransaksiHider.GetTransaksiHiderById(header.IDTrans)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("failed to get transaksi: %w", err)))
		return
	}

	c.JSON(http.StatusCreated, resp)
}
