package main

import (
	"log"
	"os"
	"net/http"
	"github.com/cristiano_code/qrcode/handlers"
)

func main() {
	fileServerTemplates := http.FileServer(http.Dir("./templates"))
	fileServerStatic := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServerTemplates)
	http.Handle("/static/", http.StripPrefix("/static/", fileServerStatic))

	http.HandleFunc("/qrcode", handlers.GenarateQrcode)

	log.Print("Server running in :8080")
	http.ListenAndServe(":" + os.Getenv("PORT"), nil)
}