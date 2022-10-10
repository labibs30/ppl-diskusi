package controller

import (
	"dataekspor-be/common"
	"dataekspor-be/dto"
	"dataekspor-be/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CityController interface {
	InsertCity(ctx *gin.Context)
	GetAllCity(ctx *gin.Context)
	GetCityByID(ctx *gin.Context)
	UpdateCity(ctx *gin.Context)
	DeleteCity(ctx *gin.Context)
}

type cityController struct {
	cityService service.CityService
}

func NewCityController(cityService service.CityService) CityController {
	return &cityController{
		cityService: cityService,
	}
}

func (c *cityController) InsertCity(ctx *gin.Context) {
	var cityDTO dto.AddCityDTO

	if err := ctx.ShouldBind(&cityDTO); err != nil {
		res := common.BuildErrorResponse("Failed to bind city request", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.cityService.CreateCity(ctx.Request.Context(), cityDTO)
	if err != nil {
		res := common.BuildErrorResponse("Failed to create city", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *cityController) GetAllCity(ctx *gin.Context) {
	result, err := c.cityService.GetAllCity(ctx.Request.Context())
	if err != nil {
		res := common.BuildErrorResponse("Failed to get city request", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *cityController) GetCityByID(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		res := common.BuildErrorResponse("Failed to find city", "id is empty", common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.cityService.GetCityByID(ctx.Request.Context(), id)
	if err != nil {
		res := common.BuildErrorResponse("Failed to get city request", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *cityController) UpdateCity(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		res := common.BuildErrorResponse("Failed to find city", "id is empty", common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var city dto.UpdateCityDTO

	if err := ctx.ShouldBind(&city); err != nil {
		res := common.BuildErrorResponse("Failed to bind city request", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.cityService.UpdateCity(ctx.Request.Context(), id, city)
	if err != nil {
		res := common.BuildErrorResponse("Failed to update city request", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *cityController) DeleteCity(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		res := common.BuildErrorResponse("Failed to find city", "id is empty", common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err := c.cityService.DeleteCity(ctx.Request.Context(), id)
	if err != nil {
		res := common.BuildErrorResponse("Failed to find city", "id is empty", common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "deleted", "data with id "+id+" has been deleted")
	ctx.JSON(http.StatusOK, res)
}
