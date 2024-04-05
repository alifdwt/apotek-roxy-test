package handler

import (
	transaksihider "apotek-roxy/domain/requests/transaksi_hider"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initTransaksiHiderGroup(api *gin.Engine) {
	transaksiHider := api.Group("/transaksi-hider")

	transaksiHider.GET("/", h.handlerGetTransaksiHiderAll)
	transaksiHider.GET("/:idTrans", h.handlerGetTransaksiHiderById)

	transaksiHider.POST("/", h.handlerCreateTransaksiHider)
	transaksiHider.PUT("/:idTrans", h.handlerUpdateTransaksiHider)
	transaksiHider.DELETE("/:idTrans", h.handlerDeleteTransaksiHider)
}

func (h *Handler) handlerGetTransaksiHiderAll(c *gin.Context) {
	res, err := h.services.TransaksiHider.GetTransaksiHiderAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) handlerGetTransaksiHiderById(c *gin.Context) {
	idTrans := c.Param("idTrans")
	res, err := h.services.TransaksiHider.GetTransaksiHiderById(idTrans)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) handlerCreateTransaksiHider(c *gin.Context) {
	var body transaksihider.CreateTransaksiHiderRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// tglTransString := body.TglTrans
	// tglTrans, err := time.Parse("2006-01-02", tglTransString)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, errorResponse(err))
	// 	return
	// }

	if err := body.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res, err := h.services.TransaksiHider.CreateTransaksiHider(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h *Handler) handlerUpdateTransaksiHider(c *gin.Context) {
	var body transaksihider.UpdateTransaksiHiderRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := body.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	idTrans := c.Param("idTrans")
	res, err := h.services.TransaksiHider.UpdateTransaksiHider(idTrans, body)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) handlerDeleteTransaksiHider(c *gin.Context) {
	idTrans := c.Param("idTrans")
	res, err := h.services.TransaksiHider.DeleteTransaksiHider(idTrans)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}
