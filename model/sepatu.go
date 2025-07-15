package model

import "UAS-SD/node"

var DaftarSepatu node.Listsepatu

func CreateSepatu(emp node.Sepatu) bool {
	tempLL := node.Listsepatu{
		Data: emp,
		Link: nil,
	}
	if DaftarSepatu.Link == nil {
		DaftarSepatu.Link = &tempLL
		return true
	} else {
		temp := &DaftarSepatu
		for temp.Link != nil {
			temp = temp.Link
		}
		temp.Link = &tempLL
		return true
	}
}

func ReadSepatu() []node.Sepatu {
	var daftarSepatu []node.Sepatu
	temp := DaftarSepatu.Link
	for temp != nil {
		daftarSepatu = append(daftarSepatu, temp.Data)
		temp = temp.Link
	}
	return daftarSepatu
}

func UpdateSepatu(emp node.Sepatu, id int) bool {
	temp := DaftarSepatu.Link
	for temp != nil {
		if temp.Data.ID == id {
			temp.Data = emp
			return true
		}
		temp = temp.Link
	}
	return false
}

func DeleteSepatu(id int) bool {
	temp := &DaftarSepatu
	for temp.Link != nil {
		if temp.Link.Data.ID == id {
			temp.Link = temp.Link.Link
			return true
		}
		temp = temp.Link
	}
	return false
}

// SearchSepatu mengembalikan struct Sepatu dan status ditemukan
func SearchSepatu(id int) (node.Sepatu, bool) {
	temp := DaftarSepatu.Link
	for temp != nil {
		if temp.Data.ID == id {
			return temp.Data, true
		}
		temp = temp.Link
	}
	return node.Sepatu{}, false
}

func GetNamaSepatu(ID int) string {
	temp := &DaftarSepatu
	for temp.Link != nil {
		if temp.Link.Data.ID == ID {
			return temp.Link.Data.Nama
		}
		temp = temp.Link
	}
	return ""
}
