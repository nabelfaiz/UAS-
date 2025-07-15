package node

type Sepatu struct {
	ID    int
	Nama  string
	Harga int
}

type Listsepatu struct {
	Data Sepatu
	Link *Listsepatu
}
