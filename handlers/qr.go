package handlers

import (
	"image/png"
	"net/http"

	"QRGo/utils"
)

func QRHandler(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query().Get("data")
	if data == "" {
		http.Error(w, "Missing data", http.StatusBadRequest)
		return
	}
	if len(data) > 53 || !utils.IsAllowed(data) {
		http.Error(w, "Bad data input", http.StatusBadRequest)
		return
	}
	
	m := utils.GenerateQRCode(data)
	img := utils.GenerateImage(m)
	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, img)
}
