package controller

import (
	"dataekspor-be/common"
	"dataekspor-be/dto"
	"dataekspor-be/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SupplierController interface {
	InsertSupplier(ctx *gin.Context)
	GetAllSupplier(ctx *gin.Context)
	GetSupplierByID(ctx *gin.Context)
	UpdateSupplier(ctx *gin.Context)
	DeleteSupplier(ctx *gin.Context)
}

type supplierController struct {
	supplierService service.SupplierService
}

func NewSupplierController(supplierService service.SupplierService) SupplierController {
	return &supplierController{
		supplierService: supplierService,
	}
}

func (c *supplierController) InsertSupplier(ctx *gin.Context) {
	var supplierDTO dto.AddSupplierDTO

	if err := ctx.ShouldBindJSON(&supplierDTO); err != nil {
		res := common.BuildErrorResponse("Failed to bind supplier request", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	fmt.Println("setelah di bind:", supplierDTO)

	result, err := c.supplierService.CreateSupplier(ctx.Request.Context(), supplierDTO)
	if err != nil {
		res := common.BuildErrorResponse("Failed to create supplier", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *supplierController) GetAllSupplier(ctx *gin.Context) {
	result, err := c.supplierService.GetAllSupplier(ctx.Request.Context())
	if err != nil {
		res := common.BuildErrorResponse("Failed to get supplier request", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *supplierController) GetSupplierByID(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		res := common.BuildErrorResponse("Failed to find supplier", "id is empty", common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.supplierService.GetSupplierByID(ctx.Request.Context(), id)
	if err != nil {
		res := common.BuildErrorResponse("Failed to get supplier request", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *supplierController) UpdateSupplier(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		res := common.BuildErrorResponse("Failed to find supplier", "id is empty", common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var supplier dto.UpdateSupplierDTO

	if err := ctx.ShouldBind(&supplier); err != nil {
		res := common.BuildErrorResponse("Failed to bind supplier request", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.supplierService.UpdateSupplier(ctx.Request.Context(), id, supplier)
	if err != nil {
		res := common.BuildErrorResponse("Failed to update supplier request", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *supplierController) DeleteSupplier(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		res := common.BuildErrorResponse("Failed to find supplier", "id is empty", common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err := c.supplierService.DeleteSupplier(ctx.Request.Context(), id)
	if err != nil {
		res := common.BuildErrorResponse("Failed to find supplier", "id is empty", common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "deleted", "data with id "+id+" has been deleted")
	ctx.JSON(http.StatusOK, res)
}
