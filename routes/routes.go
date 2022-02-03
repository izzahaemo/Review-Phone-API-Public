package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"dataphone/controllers"
	"dataphone/middlewares"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	//loginregister
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	//role
	roleMiddlewareRoute := r.Group("/role")
	roleMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	roleMiddlewareRoute.GET("/role", controllers.GetAllRole)
	r.POST("/role", controllers.CreateRole)
	r.GET("/role/:id", controllers.GetRoleById)
	roleMiddlewareRoute.PATCH("/:id", controllers.UpdateRole)
	roleMiddlewareRoute.DELETE("/:id", controllers.DeleteRole)

	//Menu
	menuMiddlewareRoute := r.Group("/menu")
	menuMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	r.GET("/menu/role/:id", controllers.GetMenuByRoleId)
	menuMiddlewareRoute.POST("/", controllers.CreateMenu)
	menuMiddlewareRoute.PATCH("/:id", controllers.UpdateMenu)
	menuMiddlewareRoute.DELETE("/:id", controllers.DeleteMenu)

	//SubMenu
	submenuMiddlewareRoute := r.Group("/submenu")
	submenuMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	r.GET("/submenu", controllers.GetAllSubmenu)
	r.GET("/submenu/:id", controllers.GetSubmenuById)
	submenuMiddlewareRoute.POST("/", controllers.CreateSubmenu)
	submenuMiddlewareRoute.PATCH("/:id", controllers.UpdateSubmenu)
	submenuMiddlewareRoute.DELETE("/:id", controllers.DeleteSubmenu)

	//accessmenu
	accessmenuMiddlewareRoute := r.Group("/accessmenu")
	accessmenuMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	accessmenuMiddlewareRoute.POST("/", controllers.CreateAccessmenu)
	accessmenuMiddlewareRoute.PATCH("/:id", controllers.UpdateAccessmenu)
	accessmenuMiddlewareRoute.DELETE("/:id", controllers.DeleteAccessmenu)

	//brand
	brandMiddlewareRoute := r.Group("/brand")
	brandMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	r.GET("/brand", controllers.GetAllBrand)
	r.GET("/brand/:id", controllers.GetBrandById)
	r.GET("/brand/picture/:id", controllers.GetBrandPicture)
	brandMiddlewareRoute.POST("/", controllers.CreateBrand)
	brandMiddlewareRoute.POST("/upload/:id", controllers.UploadBrand)
	brandMiddlewareRoute.PATCH("/:id", controllers.UpdateBrand)
	brandMiddlewareRoute.DELETE("/:id", controllers.DeleteBrand)

	//phone
	phoneMiddlewareRoute := r.Group("/phone")
	r.GET("/phone", controllers.GetAllPhone)
	r.GET("/phone/:id", controllers.GetPhoneById)
	r.GET("/phone/brand/:id", controllers.GetPhoneByBrand)
	r.GET("/phone/picture/:id", controllers.GetPhonePicture)
	phoneMiddlewareRoute.POST("/", controllers.CreatePhone)
	phoneMiddlewareRoute.POST("/upload/:id", controllers.UploadPhone)
	phoneMiddlewareRoute.PATCH("/:id", controllers.UpdatePhone)
	phoneMiddlewareRoute.DELETE("/:id", controllers.DeletePhone)

	//review
	reviewMiddlewareRoute := r.Group("/review")
	r.GET("/review/:id", controllers.GetReviewById)
	r.GET("/review/phone/:id", controllers.GetReviewByPhone)
	reviewMiddlewareRoute.POST("/", controllers.CreateReview)
	reviewMiddlewareRoute.PATCH("/:id", controllers.UpdateReview)
	reviewMiddlewareRoute.DELETE("/:id", controllers.DeleteReview)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
