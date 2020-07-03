package route

import (
	"context"
	"fmt"
	"log"
	"mini-ecommerce/handler"
	"mini-ecommerce/middleware"
	"net/http"
	"path/filepath"

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
	userProtectedRoutes := apiRoutes.Group("/user", middleware.AuthorizeJWT())

	{
		userRoutes.POST("/", userHandler.AddUser)
		userRoutes.POST("signin", userHandler.SignInUser)
	}

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

	// r.Static("/file", "saved")

	r.POST("/one", func(c *gin.Context) {

		// ctx := context.Background()

		fmt.Println("file Upload")
		// single file
		file, err := c.FormFile("profile")
		if err != nil {
			log.Fatal(err)
		}
		log.Println(file.Filename)
		filepath.Base(file.Filename)

		err = c.SaveUploadedFile(file, "saved/"+file.Filename)
		if err != nil {
			log.Fatal(err)
		}
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	r.POST("/many", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		fmt.Println(form.Value["name"])
		files := form.File["file"]

		for _, file := range files {
			log.Println(file.Filename)
			err := c.SaveUploadedFile(file, "saved/"+file.Filename)
			if err != nil {
				log.Fatal(err)
			}
		}

		fmt.Println(c.PostForm("key"))
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})

	return r.Run(address)

}
