package handlers

import (
	"image/png"
	"net/http"
	"github.com/gin-gonic/gin"
	"QRGo/utils"
)

func QRHandler(c *gin.Context) {
	data := c.Query("data")
	if data == "" {
		c.String(http.StatusBadRequest, "Missing data")
		return
	}
	if len(data) > 53 || !utils.IsAllowed(data) {
		c.String(http.StatusBadRequest, "Bad data input")
		return
	}

	m := utils.GenerateQRCode(data)
	img := utils.GenerateImage(m)
	c.Header("Content-Type", "image/png")
	png.Encode(c.Writer, img)
}
