package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/404th/viabucket/fileUploader"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error while loading .env: %s", err.Error())
	}

	e := gin.Default()
	//set a lower memory limit for multipart forms (default is 32 Mib)
	e.MaxMultipartMemory = 8 << 20 //8Mib

	e.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"mssage": err.Error(),
			})
		}
		log.Println(file.Filename) //print filename
		// dst := fmt.Sprintf("./%s", file.Filename) //The location of the build target file
		//save the file to specific dst
		// c.SaveUploadedFile(file, dst)

		fileUploader.UploadFile(file)
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%s uploaded!", file.Filename),
		})
	})
	e.Run(":9090")
}
