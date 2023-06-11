package middleware

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func UploadFile(next gin.HandlerFunc) gin.HandlerFunc {
	return func (c *gin.Context)  {
		file, err := c.FormFile("image")
		if err != nil{
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		src, err := file.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		defer src.Close()

		tempFile, err := ioutil.TempFile("uploads", "image-*.png")
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		defer tempFile.Close()

		if _, err := io.Copy(tempFile, src); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		data := tempFile.Name()

		update := strings.Split(data, "\\")[1]
		fmt.Println(update)

		c.Set("dataFile", update)
		next(c)
	}
}