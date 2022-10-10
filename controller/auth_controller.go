package controller

import (
	"dataekspor-be/common"
	"dataekspor-be/dto"
	"dataekspor-be/entity"
	"dataekspor-be/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.UserLoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := common.BuildErrorResponse("Failed to process request", errDTO.Error(), common.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(entity.User); ok {
		generatedToken := c.jwtService.GenerateToken(v.ID, v.Role...)
		v.Token = generatedToken
		response := common.BuildResponse(true, "OK", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := common.BuildErrorResponse("Error Logging in", "Invalid Credentials", nil)
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.UserRegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := common.BuildErrorResponse("Failed to process request", errDTO.Error(), common.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !c.authService.IsDuplicateEmail(registerDTO.Email) {
		response := common.BuildErrorResponse("Failed to process request", "Duplicate Email", common.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		createdUser := c.authService.CreateUser(registerDTO)
		token := c.jwtService.GenerateToken(createdUser.ID, createdUser.Role...)
		createdUser.Token = token
		response := common.BuildResponse(true, "OK", createdUser)
		ctx.JSON(http.StatusCreated, response)
	}

}
