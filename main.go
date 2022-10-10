package main

import (
	"dataekspor-be/config"
	"dataekspor-be/controller"
	"dataekspor-be/middleware"
	"dataekspor-be/repository"
	"dataekspor-be/service"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	authService    service.AuthService       = service.NewAuthService(userRepository)
	userService    service.UserService       = service.NewUserService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	userController controller.UserController = controller.NewUserController(userService, jwtService, authService)

	categoryRepository repository.CategoryRepository = repository.NewCategoryRepository(db)
	categoryService    service.CategoryService       = service.NewCategoryService(categoryRepository)
	categoryController controller.CategoryController = controller.NewCategoryController(categoryService)

	cityRepository repository.CityRepository = repository.NewCityRepository(db)
	cityService    service.CityService       = service.NewCityService(cityRepository)
	cityController controller.CityController = controller.NewCityController(cityService)

	productRepository repository.ProductRepository = repository.NewProductRepository(db)
	productService    service.ProductService       = service.NewProductService(productRepository)
	productController controller.ProductController = controller.NewProductController(productService)

	partnershipRepository repository.PartnershipRepository = repository.NewPartnershipRepository(db)
	partnershipService    service.PartnershipService       = service.NewPartnershipService(partnershipRepository)
	partnershipController controller.PartnershipController = controller.NewPartnershipController(jwtService, partnershipService)

	supplierRepository repository.SupplierRepository = repository.NewSupplierRepository(db)
	supplierService    service.SupplierService       = service.NewSupplierService(supplierRepository)
	supplierController controller.SupplierController = controller.NewSupplierController(supplierService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	server := gin.Default()

	server.Use(
		middleware.CORSMiddleware(),
		gin.Recovery(),
	)

	authRoutes := server.Group("/api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	adminRoutes := server.Group("/api/admin", middleware.AuthorizeJWT(jwtService, "admin"))
	{
		adminRoutes.GET("/eksportir/all", userController.GetAllEksportir)
		adminRoutes.GET("/user/:id", userController.GetUserByID)
	}

	eksportirRoutes := server.Group("/api/eksportir", middleware.AuthorizeJWT(jwtService, "eksportir", "supplier"))
	{
		eksportirRoutes.GET("/profile", userController.GetUserProfile)
		eksportirRoutes.PUT("/profile", userController.UpdateUser)
	}

	categoryPrivateRoutes := server.Group("/api/category", middleware.AuthorizeJWT(jwtService, "admin"))
	{
		categoryPrivateRoutes.POST("", categoryController.InsertCategory)
		categoryPrivateRoutes.PUT("/:id", categoryController.UpdateCategory)
		categoryPrivateRoutes.DELETE("/:id", categoryController.DeleteCategory)
	}

	categoryRoutes := server.Group("/api/category")
	{
		categoryRoutes.GET("", categoryController.GetAllCategory)
		categoryRoutes.GET(":id", categoryController.GetCategoryByID)
	}

	cityPrivateRoutes := server.Group("/api/city", middleware.AuthorizeJWT(jwtService, "admin"))
	{
		cityPrivateRoutes.POST("", cityController.InsertCity)
		cityPrivateRoutes.PUT("/:id", cityController.UpdateCity)
		cityPrivateRoutes.DELETE(":id", cityController.DeleteCity)
	}

	cityRoutes := server.Group("/api/city")
	{
		cityRoutes.GET("", cityController.GetAllCity)
		cityRoutes.GET("/:id", cityController.GetCityByID)

	}

	productPrivateRoutes := server.Group("/api/product", middleware.AuthorizeJWT(jwtService, "supplier"))
	{
		productPrivateRoutes.POST("", productController.InsertProduct)
		productPrivateRoutes.PUT("/:id", productController.UpdateProduct)
		productPrivateRoutes.DELETE(":id", productController.DeleteProduct)
	}

	productRoutes := server.Group("/api/product")
	{
		productRoutes.GET("", productController.GetAllProducts)
		productRoutes.GET("/:id", productController.GetProductByID)
		productRoutes.GET("/find", productController.GetProductByNameOrDesc)
	}

	supplierPrivateRoutes := server.Group("/api/supplier", middleware.AuthorizeJWT(jwtService, "supplier"))
	{
		supplierPrivateRoutes.POST("", supplierController.InsertSupplier)
		supplierPrivateRoutes.PUT("/:id", supplierController.UpdateSupplier)
		supplierPrivateRoutes.DELETE("/:id", supplierController.DeleteSupplier)
	}

	supplierRoutes := server.Group("/api/supplier")
	{
		supplierRoutes.GET("", supplierController.GetAllSupplier)
		supplierRoutes.GET("/:id", supplierController.GetSupplierByID)
	}

	partnershipPrivateRoutes := server.Group("/api/partnership", middleware.AuthorizeJWT(jwtService, "eksportir"))
	{
		partnershipPrivateRoutes.POST("", partnershipController.InsertPartnership)
		partnershipPrivateRoutes.PUT("/:id", partnershipController.UpdatePartnership)
		partnershipPrivateRoutes.DELETE(":id", partnershipController.DeletePartnership)
	}

	partnershipRoutes := server.Group("/api/partnership")
	{
		partnershipRoutes.GET("", partnershipController.GetAllPartnership)
		partnershipRoutes.GET("/:id", partnershipController.GetPartnershipByID)
	}

	testRolesRoutes := server.Group("/api/test", middleware.AuthorizeJWT(jwtService, "admin", "eksportir"))
	{
		testRolesRoutes.GET("", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"status": "authorized",
			})
		})
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)
}
