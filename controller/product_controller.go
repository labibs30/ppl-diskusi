package controller

import (
	"dataekspor-be/common"
	"dataekspor-be/dto"
	"dataekspor-be/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController interface {
	InsertProduct(ctx *gin.Context)
	GetAllProducts(ctx *gin.Context)
	GetProductByID(ctx *gin.Context)
	GetProductByNameOrDesc(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
	DeleteProduct(ctx *gin.Context)
}

type productController struct {
	productService service.ProductService
}

func NewProductController(service service.ProductService) ProductController {
	return &productController{
		productService: service,
	}
}

func (c *productController) InsertProduct(ctx *gin.Context) {
	var productDTO dto.AddProductDTO

	if err := ctx.ShouldBind(&productDTO); err != nil {
		res := common.BuildErrorResponse("Failed to bind product request", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.productService.CreateProduct(ctx.Request.Context(), productDTO)
	if err != nil {
		res := common.BuildErrorResponse("Failed to create product", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *productController) GetAllProducts(ctx *gin.Context) {
	result, err := c.productService.GetAllProduct(ctx.Request.Context())
	if err != nil {
		res := common.BuildErrorResponse("Failed to find product", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *productController) GetProductByID(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		res := common.BuildErrorResponse("Failed to find category", "id is empty", common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.productService.GetProductByID(ctx.Request.Context(), id)
	if err != nil {
		res := common.BuildErrorResponse("Failed to find product", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *productController) GetProductByNameOrDesc(ctx *gin.Context) {
	param := ctx.Query("search_query")

	result, err := c.productService.GetProductByNameOrDesc(ctx.Request.Context(), param)
	if err != nil {
		res := common.BuildErrorResponse("Failed to find product", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *productController) UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		res := common.BuildErrorResponse("Failed to find product", "id is empty", common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var product dto.UpdateProductDTO

	if err := ctx.ShouldBind(&product); err != nil {
		res := common.BuildErrorResponse("Failed to bind product request", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.productService.UpdateProduct(ctx.Request.Context(), id, product)
	if err != nil {
		res := common.BuildErrorResponse("Failed to update product request", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *productController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		res := common.BuildErrorResponse("Failed to find product", "id is empty", common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err := c.productService.DeleteProduct(ctx.Request.Context(), id)
	if err != nil {
		res := common.BuildErrorResponse("Failed to find product", "id is empty", common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "deleted", "data with id "+id+" has been deleted")
	ctx.JSON(http.StatusOK, res)
}
