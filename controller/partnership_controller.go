package controller

import (
	"dataekspor-be/common"
	"dataekspor-be/dto"
	"dataekspor-be/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type PartnershipController interface {
	InsertPartnership(ctx *gin.Context)
	GetAllPartnership(ctx *gin.Context)
	GetPartnershipByID(ctx *gin.Context)
	UpdatePartnership(ctx *gin.Context)
	DeletePartnership(ctx *gin.Context)
}

type partnershipController struct {
	jwtService         service.JWTService
	partnershipService service.PartnershipService
}

func NewPartnershipController(jwtService service.JWTService, partnershipService service.PartnershipService) PartnershipController {
	return &partnershipController{
		partnershipService: partnershipService,
		jwtService:         jwtService,
	}
}

func (c *partnershipController) InsertPartnership(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	authHeader = strings.Replace(authHeader, "Bearer ", "", -1)
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		response := common.BuildErrorResponse("Invalid Token", errToken.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	userIdFromJWT, ok := claims["user_id"].(string)

	if !ok {
		response := common.BuildErrorResponse("Unknown request", "bad request on user_id in jwt", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	var partnershipDTO dto.AddPartnershipDTO

	if err := ctx.ShouldBind(&partnershipDTO); err != nil {
		res := common.BuildErrorResponse("Failed to bind partnership request", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if partnershipDTO.UserID == uuid.Nil {
		userID, err := uuid.Parse(userIdFromJWT)
		if err != nil {
			res := common.BuildErrorResponse("Failed to bind partnership request", err.Error(), common.EmptyObj{})
			ctx.JSON(http.StatusBadRequest, res)
			return
		}

		partnershipDTO.UserID = userID
	}

	result, err := c.partnershipService.CreatePartnership(ctx.Request.Context(), partnershipDTO)
	if err != nil {
		res := common.BuildErrorResponse("Failed to create partnership", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *partnershipController) GetAllPartnership(ctx *gin.Context) {
	result, err := c.partnershipService.GetAllPartnership(ctx.Request.Context())
	if err != nil {
		res := common.BuildErrorResponse("Failed to get partnership request", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *partnershipController) GetPartnershipByID(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		res := common.BuildErrorResponse("Failed to find partnership", "id is empty", common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.partnershipService.GetPartnershipByID(ctx.Request.Context(), id)
	if err != nil {
		res := common.BuildErrorResponse("Failed to get partnership request", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *partnershipController) UpdatePartnership(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		res := common.BuildErrorResponse("Failed to find partnership", "id is empty", common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var partnership dto.UpdatePartnershipDTO

	if err := ctx.ShouldBind(&partnership); err != nil {
		res := common.BuildErrorResponse("Failed to bind partnership request", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.partnershipService.UpdatePartnership(ctx.Request.Context(), id, partnership)
	if err != nil {
		res := common.BuildErrorResponse("Failed to update partnership request", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *partnershipController) DeletePartnership(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		res := common.BuildErrorResponse("Failed to find partnership", "id is empty", common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err := c.partnershipService.DeletePartnership(ctx.Request.Context(), id)
	if err != nil {
		res := common.BuildErrorResponse("Failed to find partnership", "id is empty", common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "deleted", "data with id "+id+" has been deleted")
	ctx.JSON(http.StatusOK, res)
}
