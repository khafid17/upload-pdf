package main

import (
	"fmt"
	"net/http"
)

// Handler fungsi untuk endpoint hello
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Menulis respon teks ke browser/client
	fmt.Fprint(w, "Hello, World! Server Golang kamu berhasil berjalan dengan lancar.")
}

func main() {
	// Daftarkan endpoint "/" (root) ke fungsi helloHandler
	http.HandleFunc("/", helloHandler)

	fmt.Println("Server berjalan di http://localhost:8080 ...")
	// Jalankan server di port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Gagal menjalankan server:", err)
	}
}