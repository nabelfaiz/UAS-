package model

import "UAS-SD/node"

var DaftarPembeli node.ListPembeli

// CREATE
func CreatePembeli(jab node.Pembeli) bool {
	tempLL := node.ListPembeli{
		Data: jab,
		Link: nil,
	}
	if DaftarPembeli.Link == nil {
		DaftarPembeli.Link = &tempLL
		return true
	} else {
		temp := &DaftarPembeli
		for temp.Link != nil {
			temp = temp.Link
		}
		temp.Link = &tempLL
		return true
	}
}

// READ
func ReadPembeli() []node.Pembeli {
	daftarPembeli := []node.Pembeli{}
	temp := &DaftarPembeli
	for temp.Link != nil {
		pembeliData := temp.Link.Data
		sepatu, found := SearchSepatu(pembeliData.Sepatu) // Use the Sepatu ID stored in Pembeli
		if found {
			pembeliData.Barang = sepatu.Nama // Set Barang to the Sepatu's Name
		} else {
			pembeliData.Barang = "N/A" // Or handle as appropriate if Sepatu not found
		}
		daftarPembeli = append(daftarPembeli, pembeliData)
		temp = temp.Link
	}
	return daftarPembeli
}

// UPDATE
func UpdatePembeli(jab node.Pembeli, ID int) bool {
	temp := DaftarPembeli.Link
	for temp != nil {
		if temp.Data.ID == ID {
			temp.Data = jab
			return true
		}
		temp = temp.Link
	}
	return false
}

// DELETE
func DeletePembeli(ID int) bool {
	temp := &DaftarPembeli
	for temp.Link != nil {
		if temp.Link.Data.ID == ID {
			temp.Link = temp.Link.Link
			return true
		}
		temp = temp.Link
	}
	return false
}

// SearchPembeli mengembalikan struct Pembeli dan status ditemukan
func SearchPembeli(ID int) (node.Pembeli, bool) {
	temp := DaftarPembeli.Link
	for temp != nil {
		if temp.Data.ID == ID {
			return temp.Data, true
		}
		temp = temp.Link
	}
	return node.Pembeli{}, false
}

func GetNama(ID int) string {
	temp := &DaftarPembeli
	for temp.Link != nil {
		if temp.Link.Data.ID == ID {
			return temp.Link.Data.Nama
		}
		temp = temp.Link
	}
	return ""
}
