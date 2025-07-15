package main

import (
	"UAS-SD/model"
	"UAS-SD/node"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Inisialisasi list kosong
	model.DaftarSepatu = node.Listsepatu{}
	model.DaftarPembeli = node.ListPembeli{}

	// Contoh data awal
	model.CreateSepatu(node.Sepatu{ID: 1, Nama: "Nike Air Max", Harga: 1200000})
	model.CreateSepatu(node.Sepatu{ID: 2, Nama: "Adidas Ultraboost", Harga: 1500000})

	// Create Pembeli with full details
	pembeli1 := node.Pembeli{
		ID:     1,
		Nama:   "Budi",
		Sepatu: 1, // ID Sepatu Nike Air Max
		Email:  "budi@example.com",
		NoTelp: "08123456789",
		Alamat: node.Pembeli{}.Alamat, // Initialize Alamat struct
	}
	pembeli1.Alamat.Kota = "Jakarta"
	pembeli1.Alamat.Jalan = "Jl. Sudirman"
	pembeli1.Alamat.Nomer = 10
	model.CreatePembeli(pembeli1)

	pembeli2 := node.Pembeli{
		ID:     2,
		Nama:   "Ani",
		Sepatu: 2, // ID Sepatu Adidas Ultraboost
		Email:  "ani@example.com",
		NoTelp: "0876543210",
		Alamat: node.Pembeli{}.Alamat, // Initialize Alamat struct
	}
	pembeli2.Alamat.Kota = "Bandung"
	pembeli2.Alamat.Jalan = "Jl. Asia Afrika"
	pembeli2.Alamat.Nomer = 5
	model.CreatePembeli(pembeli2)

	// Register HTTP handlers
	http.HandleFunc("/", RenderIndexPage)
	http.HandleFunc("/tambah", RenderFormPage)
	http.HandleFunc("/insert", InsertPembeliHandler)
	http.HandleFunc("/edit", RenderEditPage)
	http.HandleFunc("/update", UpdatePembeliHandler)
	http.HandleFunc("/delete", DeletePembeliHandler)

	// Start the HTTP server
	port := ":8080"
	fmt.Printf("Server started on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
