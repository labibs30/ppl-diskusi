package controller

import (
	"dataekspor-be/common"
	"dataekspor-be/dto"
	"dataekspor-be/entity"
	"dataekspor-be/service"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type UserController interface {
	GetAllEksportir(ctx *gin.Context)
	GetUserByID(ctx *gin.Context)
	GetUserProfile(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
	authService service.AuthService
}

func NewUserController(userService service.UserService, jwtService service.JWTService, authService service.AuthService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
		authService: authService,
	}
}

func (c *userController) GetAllEksportir(ctx *gin.Context) {
	var eksportir []entity.User = c.userService.GetAllEksportir()
	res := common.BuildResponse(true, "OK", eksportir)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		res := common.BuildErrorResponse("No ID", "No Parameter ID was found", common.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var user entity.User = c.userService.GetUserByID(id)
	if user.IsZero() {
		res := common.BuildErrorResponse("Data not found", "No data with given ID", common.EmptyObj{})
		ctx.JSON(http.StatusOK, res)
	} else {
		res := common.BuildResponse(true, "OK", user)
		ctx.JSON(http.StatusOK, res)
	}
}

func (c *userController) GetUserProfile(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	authHeader = strings.Replace(authHeader, "Bearer ", "", -1)
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		response := common.BuildErrorResponse("Invalid Token", errToken.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	var user entity.User = c.userService.GetUserByID(id)
	if user.IsZero() {
		res := common.BuildErrorResponse("Data not found", "No corresponding user", common.EmptyObj{})
		ctx.JSON(http.StatusOK, res)
	} else {
		res := common.BuildResponse(true, "OK", user)
		ctx.JSON(http.StatusOK, res)
	}
}

func (c *userController) UpdateUser(ctx *gin.Context) {
	var userDTO dto.UserUpdateDTO
	errDTO := ctx.ShouldBind(&userDTO)
	if errDTO != nil {
		res := common.BuildErrorResponse("Failed to process request", errDTO.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	authHeader := ctx.GetHeader("Authorization")
	authHeader = strings.Replace(authHeader, "Bearer ", "", -1)
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		response := common.BuildErrorResponse("Invalid Token", errToken.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])

	userDTO.ID = id

	var checkUserEmail = c.userService.GetUserByID(id)
	if checkUserEmail.Email != userDTO.Email && !c.authService.IsDuplicateEmail(userDTO.Email) {
		response := common.BuildErrorResponse("Failed to process request", "Duplicate Email", common.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
		return
	}
	user := c.userService.UpdateUser(userDTO)
	res := common.BuildResponse(true, "OK", user)
	ctx.JSON(http.StatusOK, res)
}
