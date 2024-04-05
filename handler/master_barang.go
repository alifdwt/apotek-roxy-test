package handler

import (
	masterbarang "apotek-roxy/domain/requests/master_barang"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initMasterBarangGroup(api *gin.Engine) {
	masterBarang := api.Group("/master-barang")

	masterBarang.GET("/", h.handlerGetMasterBarangAll)
	masterBarang.GET("/:idBarang", h.handlerGetMasterBarangById)
	masterBarang.GET("/search/:nmBarang", h.handlerGetMasterBarangByNmBarang)

	masterBarang.POST("/", h.handlerCreateMasterBarang)
	masterBarang.PUT("/:idBarang", h.handlerUpdateMasterBarang)
	masterBarang.DELETE("/:idBarang", h.handlerDeleteMasterBarang)
}

func (h *Handler) handlerGetMasterBarangAll(c *gin.Context) {
	res, err := h.services.MasterBarang.GetMasterBarangAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) handlerGetMasterBarangById(c *gin.Context) {
	idBarang := c.Param("idBarang")
	res, err := h.services.MasterBarang.GetMasterBarangById(idBarang)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) handlerGetMasterBarangByNmBarang(c *gin.Context) {
	nmBarang := c.Param("nmBarang")
	res, err := h.services.MasterBarang.GetMasterBarangByNmBarang(nmBarang)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) handlerCreateMasterBarang(c *gin.Context) {
	var body masterbarang.CreateMasterBarangRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := body.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res, err := h.services.MasterBarang.CreateMasterBarang(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) handlerUpdateMasterBarang(c *gin.Context) {
	var body masterbarang.UpdateMasterBarangRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	idBarang := c.Param("idBarang")
	res, err := h.services.MasterBarang.UpdateMasterBarang(idBarang, body)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) handlerDeleteMasterBarang(c *gin.Context) {
	idBarang := c.Param("idBarang")
	res, err := h.services.MasterBarang.DeleteMasterBarang(idBarang)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}
