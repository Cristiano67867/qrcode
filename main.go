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

    // Corrigir nome da função
    http.HandleFunc("/qrcode", handlers.GenerateQrcode)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Server running on port %s", port)
    http.ListenAndServe(":"+port, nil)
}
