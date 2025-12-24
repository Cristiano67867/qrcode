package handlers

import (
	"log"
	"net/http"
	"encoding/json"
	"github.com/skip2/go-qrcode"
	"github.com/Cristiano67867/qrcode/models"
)

func GenarateQrcode(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	var qrcodeData models.Qrcode
	
	err := json.NewDecoder(r.Body).Decode(&qrcodeData)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		log.Fatal(err)
		return
	}
	defer r.Body.Close()

	qr, qr_err := qrcode.Encode(qrcodeData.Text, qrcode.Medium, 256)
	if qr_err != nil {
		http.Error(w, "Erro to genarate qrcode", http.StatusInternalServerError)
		log.Fatal(qr_err)
		return
	}

	w.Write(qr)
}
