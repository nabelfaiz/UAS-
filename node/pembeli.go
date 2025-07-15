package node

type Pembeli struct {
	ID     int
	Nama   string
	Sepatu int // This will store the ID of the Sepatu
	Email  string
	NoTelp string
	Barang string // To store the name of the Sepatu for display in index.html
	Alamat struct {
		Kota   string
		Jalan  string
		Nomer  int
	}
	Jabatan int // Assuming this is a placeholder or will be used for something else
}

type ListPembeli struct {
	Data Pembeli
	Link *ListPembeli
}
