package controller

import (
	"dataekspor-be/common"
	"dataekspor-be/dto"
	"dataekspor-be/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	GetAllCategory(ctx *gin.Context)
	InsertCategory(ctx *gin.Context)
	GetCategoryByID(ctx *gin.Context)
	GetCategoryByNameOrDesc(ctx *gin.Context)
	UpdateCategory(ctx *gin.Context)
	DeleteCategory(ctx *gin.Context)
}

type categoryController struct {
	categoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &categoryController{
		categoryService: categoryService,
	}
}

func (c *categoryController) GetAllCategory(ctx *gin.Context) {
	result, err := c.categoryService.GetAllCategory(ctx.Request.Context())
	if err != nil {
		res := common.BuildErrorResponse("Failed to get category request", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *categoryController) InsertCategory(ctx *gin.Context) {
	var categoryDTO dto.AddCategoryDTO

	if err := ctx.ShouldBind(&categoryDTO); err != nil {
		res := common.BuildErrorResponse("Failed to bind category request", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.categoryService.CreateCategory(ctx.Request.Context(), categoryDTO)
	if err != nil {
		res := common.BuildErrorResponse("Failed to create category", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *categoryController) GetCategoryByID(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		res := common.BuildErrorResponse("Failed to find category", "id is empty", common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.categoryService.GetCategoryByID(ctx.Request.Context(), id)
	if err != nil {
		res := common.BuildErrorResponse("Failed to get category request", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *categoryController) GetCategoryByNameOrDesc(ctx *gin.Context) {
	param := ctx.Query("find")

	result, err := c.categoryService.FindByNameOrDesc(ctx.Request.Context(), param)
	if err != nil {
		res := common.BuildErrorResponse("Failed to find category", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *categoryController) UpdateCategory(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		res := common.BuildErrorResponse("Failed to find category", "id is empty", common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var category dto.UpdateCategoryDTO

	if err := ctx.ShouldBind(&category); err != nil {
		res := common.BuildErrorResponse("Failed to bind category request", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.categoryService.UpdateCategory(ctx.Request.Context(), id, category)
	if err != nil {
		res := common.BuildErrorResponse("Failed to update category request", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *categoryController) DeleteCategory(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		res := common.BuildErrorResponse("Failed to find category", "id is empty", common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err := c.categoryService.DeleteCategory(ctx.Request.Context(), id)
	if err != nil {
		res := common.BuildErrorResponse("Failed to find category", "id is empty", common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "deleted", "data with id "+id+" has been deleted")
	ctx.JSON(http.StatusOK, res)
}
