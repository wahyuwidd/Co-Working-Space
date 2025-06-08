package main

import (
	"fmt"
	"strings"
)

type Space struct {
	Nama      string
	Lokasi    string
	Fasilitas []string
	Harga     int
	Ulasan    []Review
}

type Review struct {
	Rating   float64
	Komentar string
}

// var spacs []Space

var spaces = []Space{
	{
		Nama:   "CozyWork",
		Lokasi: "Jakarta",
		Fasilitas: []string{
			"WiFi",
			"AC",
		},
		Harga: 250000,
		Ulasan: []Review{
			{Rating: 4.5, Komentar: "Tempat nyaman untuk bekerja"},
			{Rating: 5.0, Komentar: "Pelayanan sangat baik"},
		},
	},
	{
		Nama:   "CenterHub",
		Lokasi: "Bandung",
		Fasilitas: []string{
			"WiFi",
			"Parkir",
		},
		Harga: 180000,
		Ulasan: []Review{
			{Rating: 4.0, Komentar: "Suasana kreatif"},
			{Rating: 4.8, Komentar: "Fasilitas lengkap"},
		},
	},
	{
		Nama:   "EksWork",
		Lokasi: "Sudirman",
		Fasilitas: []string{
			"Lounge",
			"Coffee",
		},
		Harga: 500000,
		Ulasan: []Review{
			{Rating: 4.9, Komentar: "Sangat profesional"},
			{Rating: 4.7, Komentar: "Harga sesuai kualitas"},
		},
	},
}

func TambahSpace() {
	var space Space
	var jumlahFasilitas int

	fmt.Println("\n--- Menambahkan Space Baru ---")
	fmt.Print("Masukkan nama Space        : ")
	fmt.Scan(&space.Nama)
	fmt.Print("Masukkan lokasi            : ")
	fmt.Scan(&space.Lokasi)
	fmt.Print("Masukkan jumlah fasilitas  : ")
	fmt.Scan(&jumlahFasilitas)
	space.Fasilitas = make([]string, jumlahFasilitas)
	for i := 0; i < jumlahFasilitas; i++ {
		fmt.Printf("Masukkan fasilitas ke-%d    : ", i+1)
		fmt.Scan(&space.Fasilitas[i])
	}
	fmt.Print("Masukkan harga             : ")
	fmt.Scan(&space.Harga)

	spaces = append(spaces, space)
	fmt.Println("----------------------------------------")
	fmt.Println("Space berhasil ditambahkan!")
	fmt.Println("----------------------------------------")
}

func TampilSpace() {
	fmt.Println()
	fmt.Println("DAFTAR CO-WORKING SPACE")
	fmt.Println("-----------------------")

	if len(spaces) == 0 {
		fmt.Println("Tidak ada data co-working space")
		return
	}

	fmt.Printf("%-4s %-10s %-10s %-20s %-15s\n", "No", "Nama", "Lokasi", "Fasilitas", "Harga")

	for i := 0; i < len(spaces); i++ {
		space := spaces[i]
		fmt.Printf("%-4d %-10s %-10s %-20s %-15d\n",
			i+1,
			space.Nama,
			space.Lokasi,
			strings.Join(space.Fasilitas, ","),
			space.Harga)
	}
}

func EditSpace() {
	var noSpace int
	var pilihan int

	TampilSpace()

	if len(spaces) == 0 {
		return
	}

	fmt.Println()
	fmt.Print("Masukkan nomor Space yang ingin diedit: ")
	fmt.Scan(&noSpace)

	if noSpace < 1 || noSpace > len(spaces) {
		fmt.Println("Nomor Space tidak valid.")
		return
	}

	space := spaces[noSpace-1]
	fmt.Println()
	fmt.Println("--- Mengedit Space", space.Nama, "---")

	fmt.Println("1. Edit Nama")
	fmt.Println("2. Edit Lokasi")
	fmt.Println("3. Edit Fasilitas")
	fmt.Println("4. Edit Harga")
	fmt.Println("0. Kembali")
	fmt.Print("Pilih menu:> ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 0:
		return
	case 1:
		fmt.Println("Nama saat ini		:", space.Nama)
		fmt.Print("Masukkan nama baru	: ")
		fmt.Scan(&space.Nama)
		fmt.Println("----------------------------------------")
		fmt.Println("Nama berhasil diubah!")
		fmt.Println("----------------------------------------")
		spaces[noSpace-1] = space
	case 2:
		fmt.Println("Lokasi saat ini		:", space.Lokasi)
		fmt.Print("Masukkan lokasi baru	: ")
		fmt.Scan(&space.Lokasi)
		fmt.Println("----------------------------------------")
		fmt.Println("Lokasi berhasil diubah!")
		fmt.Println("----------------------------------------")
		spaces[noSpace-1] = space
	case 3:
		fmt.Println("Fasilitas saat ini		:", space.Fasilitas)
		fmt.Print("Masukkan fasilitas baru	: ")
		fmt.Scan(&space.Fasilitas)
		fmt.Println("----------------------------------------")
		fmt.Println("Fasilitas berhasil diubah!")
		fmt.Println("----------------------------------------")
		spaces[noSpace-1] = space
	case 4:
		fmt.Println("Harga saat ini		:", space.Harga)
		fmt.Print("Masukkan harga baru	: ")
		fmt.Scan(&space.Harga)
		fmt.Println("----------------------------------------")
		fmt.Println("Harga berhasil diubah!")
		fmt.Println("----------------------------------------")
		spaces[noSpace-1] = space
	}

}

func HapusSpace() {
	var noSpace int

	TampilSpace()

	if len(spaces) == 0 {
		return
	}

	fmt.Println()
	fmt.Print("Masukkan nomor Space yang ingin dihapus: ")
	fmt.Scan(&noSpace)

	if noSpace < 1 || noSpace > len(spaces) {
		fmt.Println("Nomor Space tidak valid.")
		return
	}

	spaces = append(spaces[:noSpace-1], spaces[noSpace:]...)
	fmt.Println("----------------------------------------")
	fmt.Println("Space berhasil dihapus!")
	fmt.Println("----------------------------------------")
}

func ManajemenSpace() {
	for {
		var pilihan int
		fmt.Println()
		fmt.Println("=== Selamat datang di menu Manajemen Space ===")
		fmt.Println("[1] Tambah Space")
		fmt.Println("[2] Lihat daftar Space")
		fmt.Println("[3] Edit Space")
		fmt.Println("[4] Hapus Space")
		fmt.Println("[0] Kembali")
		fmt.Print("Pilih menu:> ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 0:
			main()
		case 1:
			TambahSpace()
		case 2:
			TampilSpace()
		case 3:
			EditSpace()
		case 4:
			HapusSpace()
		}
	}
}

func main() {
	var role int
	fmt.Println("========================================")
	fmt.Println("  Selamat datang di Co Working Space    ")
	fmt.Println("========================================")
	fmt.Println("Silahkan pilih role anda")
	fmt.Println("[1] Pemilik Co Working Space")
	fmt.Println("[2] Pekerja remote atau freelancer")
	fmt.Println("[0] Keluar")
	fmt.Print("Pilih role:> ")
	fmt.Scan(&role)

	for {
		switch role {
		case 0:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini.")
			return
		case 1:
			ManajemenSpace()
		case 2:
			ManajemenPengguna()
		default:
			fmt.Println("----------------------------------------")
			fmt.Println("Pilihan tidak valid!")
			fmt.Println("----------------------------------------")
		}
	}
}

func Searching() {

}

func Sorting() {

}

func ManajemenPengguna() {
	for {
		var pilihan int
		fmt.Println()
		fmt.Println("=== Selamat datang di menu pengguna ===")

		TampilSpace()
		fmt.Println()

		fmt.Println("Silahkan pilih menu:")
		fmt.Println("[1] Cari Tempat Co Working Space")
		fmt.Println("[0] Kembali")
		fmt.Print("Pilih menu:> ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 0:
			main()
		case 1:
			// CariSpace()
		}
	}
}
