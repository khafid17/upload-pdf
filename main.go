package handler

import (
	"fmt"
	"net/http"
)

// Handler utama yang diexport untuk Vercel Serverless
func Handler(w http.ResponseWriter, r *http.Request) {
	// Pastikan route-nya sesuai atau handle root
	fmt.Fprint(w, "Hello, World! Server Golang kamu BERHASIL berjalan di Vercel Serverless!")
}