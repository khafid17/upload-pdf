package main

import (
	"fmt"
	"net/http"
	handler "upload-pdf/api"
	// Import folder api dari module kamu (sesuaikan nama module di go.mod)
)

func main() {
	// Memanggil fungsi Handler yang ada di dalam folder api
	http.HandleFunc("/", handler.Handler)

	fmt.Println("Server lokal berjalan di http://localhost:8080 ...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Gagal menjalankan server:", err)
	}
}
