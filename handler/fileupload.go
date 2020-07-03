package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//FileUploadHandler -> Interface to File Upload
type FileUploadHandler interface {
	SingleFile(*gin.Context)
	MultipleFile(*gin.Context)
}

//SingleFile --> handle uploading of single file
func SingleFile(c *gin.Context) {

	file, err := c.FormFile("profile")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(file.Filename)

	err = c.SaveUploadedFile(file, "files/"+file.Filename)
	if err != nil {
		log.Fatal(err)
	}
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

//MultipleFile -->> handle uploading of multiple files along with other form data
func MultipleFile(c *gin.Context) {

	form, _ := c.MultipartForm()

	files := form.File["profile"]

	for _, file := range files {
		log.Println(file.Filename)
		err := c.SaveUploadedFile(file, "files/"+file.Filename)
		if err != nil {
			log.Fatal(err)
		}
	}

	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}
