package main

import (
	"UAS-SD/model"
	"UAS-SD/node"
	"html/template"
	"net/http"
	"strconv"
)

func RenderIndexPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	pembeliList := model.ReadPembeli()

	searchQuery := r.URL.Query().Get("search")
	if searchQuery != "" {
		filteredList := []node.Pembeli{}
		for _, p := range pembeliList {
			if searchQuery != "" && (p.Nama == searchQuery || p.Barang == searchQuery) {
				filteredList = append(filteredList, p)
			}
		}
		pembeliList = filteredList
	}

	err := tmpl.Execute(w, pembeliList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func RenderFormPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/form.html"))
	err := tmpl.Execute(w, nil) 
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func InsertPembeliHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()

	id, _ := strconv.Atoi(r.FormValue("id"))
	nama := r.FormValue("nama")
	sepatuID, _ := strconv.Atoi(r.FormValue("sepatu_id")) 
	email := r.FormValue("email")
	notelp := r.FormValue("notelp")
	kota := r.FormValue("kota")
	jalan := r.FormValue("jalan")
	nomer, _ := strconv.Atoi(r.FormValue("nomer"))

	newPembeli := node.Pembeli{
		ID:     id,
		Nama:   nama,
		Sepatu: sepatuID,
		Email:  email,
		NoTelp: notelp,
		Alamat: node.Pembeli{}.Alamat, 
	}
	newPembeli.Alamat.Kota = kota
	newPembeli.Alamat.Jalan = jalan
	newPembeli.Alamat.Nomer = nomer

	if model.CreatePembeli(newPembeli) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.Error(w, "Failed to add Pembeli", http.StatusInternalServerError)
	}
}

func RenderEditPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/edit.html"))

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Pembeli ID", http.StatusBadRequest)
		return
	}

	pembeli, found := model.SearchPembeli(id)
	if !found {
		http.Error(w, "Pembeli not found", http.StatusNotFound)
		return
	}

	err = tmpl.Execute(w, pembeli)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UpdatePembeliHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()

	id, _ := strconv.Atoi(r.FormValue("id"))
	nama := r.FormValue("nama")
	email := r.FormValue("email")
	notelp := r.FormValue("notelp")
	kota := r.FormValue("kota")
	jalan := r.FormValue("jalan")
	nomer, _ := strconv.Atoi(r.FormValue("nomer"))
	sepatuID, _ := strconv.Atoi(r.FormValue("sepatu_id")) 

	updatedPembeli := node.Pembeli{
		ID:     id,
		Nama:   nama,
		Email:  email,
		NoTelp: notelp,
		Sepatu: sepatuID,
		Alamat: node.Pembeli{}.Alamat, 
	}
	updatedPembeli.Alamat.Kota = kota
	updatedPembeli.Alamat.Jalan = jalan
	updatedPembeli.Alamat.Nomer = nomer

	if model.UpdatePembeli(updatedPembeli, id) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.Error(w, "Failed to update Pembeli", http.StatusInternalServerError)
	}
}

func DeletePembeliHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Pembeli ID", http.StatusBadRequest)
		return
	}

	if model.DeletePembeli(id) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.Error(w, "Failed to delete Pembeli", http.StatusInternalServerError)
	}
}
