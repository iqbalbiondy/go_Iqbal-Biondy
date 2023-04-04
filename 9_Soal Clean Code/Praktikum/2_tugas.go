//Berikut adalah kode yang sudah saya sesuaikan dengan kebiasaan programmer sepemahaman saya

package main

type Kendaraan struct {
	totalRoda       int
	kecepatanPerJam int
}

type mobil struct {
	Kendaraan
}

func (mobilVar *mobil) berjalan() {
	mobilVar.tambahkecepatan(10)
}

func (mobilVar *mobil) tambahkecepatan(kecepatanbaru int) {
	mobilVar.kecepatanPerJam = mobilVar.kecepatanPerJam + kecepatanbaru
}

func main() {
	mobilCepat := mobil{}
	mobilCepat.berjalan()
	mobilCepat.berjalan()
	mobilCepat.berjalan()

	mobilLamban := mobil{}
	mobilLamban.berjalan()
}