package route

import (
	"mini-ecommerce/handler"
	"mini-ecommerce/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

//RunAPI ->route setup
func RunAPI(address string) error {

	userHandler := handler.NewUserHandler()
	productHandler := handler.NewProductHandler()

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Main Root Application")
	})

	apiRoutes := r.Group("/api")
	userRoutes := apiRoutes.Group("/user")

	{
		userRoutes.POST("/", userHandler.AddUser)
		userRoutes.POST("signin", userHandler.SignInUser)
	}

	userProtectedRoutes := apiRoutes.Group("/user", middleware.AuthorizeJWT())
	{
		userProtectedRoutes.GET("/", userHandler.GetAllUser)
		userProtectedRoutes.GET("/:id", userHandler.GetUser)
		userProtectedRoutes.PUT("/:id", userHandler.UpdateUser)
		userProtectedRoutes.DELETE("/:id", userHandler.DeleteUser)
	}

	productRoutes := apiRoutes.Group("/product", middleware.AuthorizeJWT())
	{
		productRoutes.POST("/", productHandler.AddProduct)
		productRoutes.PUT("/:id", productHandler.UpdateProduct)
		productRoutes.DELETE("/:id", productHandler.DeleteProduct)
	}

	fileRoutes := r.Group("/file")
	{
		fileRoutes.POST("/single", handler.SingleFile)
		fileRoutes.POST("/multi", handler.MultipleFile)
	}

	return r.Run(address)

}
