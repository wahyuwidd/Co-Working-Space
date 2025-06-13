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
var reviews []Review

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

	fmt.Printf("%-4s %-15s %-15s %-25s %-10s %-8s %-40s\n", "No", "Nama", "Lokasi", "Fasilitas", "Harga", "Rating", "Komentar")

	for i := 0; i < len(spaces); i++ {
		space := spaces[i]
		var total float64
		var jumlah int
		var rataRating float64
		var komentar string

		for j := 0; j < len(space.Ulasan); j++ {
			total += space.Ulasan[j].Rating
			jumlah++
			if space.Ulasan[j].Komentar != "" {
				if komentar != "" {
					komentar += ", "
				}
				komentar += space.Ulasan[j].Komentar
			}
		}

		if jumlah > 0 {
			rataRating = total / float64(jumlah)
		} else {
			rataRating = 0.0
		}

		fmt.Printf("%-4d %-15s %-15s %-25s %-10d %-8.1f %-40s\n",
			i+1,
			space.Nama,
			space.Lokasi,
			strings.Join(space.Fasilitas, ","),
			space.Harga,
			rataRating,
			komentar)
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
			main()
		}
	}
}

func ManajemenPengguna() {
	for {
		var pilihan int
		fmt.Println()
		fmt.Println("=== Selamat datang di menu pengguna ===")

		TampilSpace()
		fmt.Println()

		fmt.Println("Silahkan pilih menu:")
		fmt.Println("[1] Cari Co Working Space")
		fmt.Println("[2] Urutkan Co Working Space")
		fmt.Println("[3] Review Co Working Space")
		fmt.Println("[4] Filter berdasarkan Fasilitas")
		fmt.Println("[0] Kembali")
		fmt.Print("Pilih menu:> ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 0:
			main()
		case 1:
			CariSpace()
		case 2:
			UrutkanSpace()
		case 3:
			RatingSpace()
		case 4:
			FilterFasilitas()
		}

	}
}

func CariSpace() {
	var pilihan int
	fmt.Println("Silahkan pilih pencarian:")
	fmt.Println("1. Cari berdasarkan Nama (Sequential Search)")
	fmt.Println("2. Cari berdasarkan Lokasi (Binary Search)")
	fmt.Println("0. Kembali")
	fmt.Print("Pilih menu:> ")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 0:
		ManajemenPengguna()
	case 1:
		CariNamaSpace()
	case 2:
		CariLokasiSpace()

	default:
		fmt.Println("----------------------------------------")
		fmt.Println("Pilihan tidak valid!")
		fmt.Println("----------------------------------------")
	}
}

func CariNamaSpace() {
	var space Space
	var cari string
	fmt.Print("Masukkan nama co-working space yang ingin dicari: ")
	fmt.Scan(&cari)

	for i := 0; i < len(spaces); i++ {
		if spaces[i].Nama == cari {
			space = spaces[i]
			fmt.Println()
			fmt.Println("Co-working space ditemukan")
			fmt.Println("-----------------------")
			fmt.Println("Nama:", space.Nama)
			fmt.Println("Lokasi:", space.Lokasi)
			fmt.Println("Fasilitas:", strings.Join(space.Fasilitas, ","))
			fmt.Println("Harga:", space.Harga)
		}
	}

	if space.Nama == "" {
		fmt.Println("Co-working space tidak ditemukan.")
	}
}

func CariLokasiSpace() {
	var space Space
	var cari string
	fmt.Print("Masukkan lokasi co-working space yang ingin dicari: ")
	fmt.Scan(&cari)

	left := 0
	right := len(spaces) - 1
	for left <= right {
		mid := (left + right) / 2
		if spaces[mid].Lokasi == cari {
			space = spaces[mid]
			fmt.Println()
			fmt.Println("Co-working space ditemukan: ")
			fmt.Println("-----------------------")
			fmt.Println("Nama:", space.Nama)
			fmt.Println("Lokasi:", space.Lokasi)
			fmt.Println("Fasilitas:", space.Fasilitas)
			fmt.Println("Harga:", space.Harga)
			break
		} else if spaces[mid].Lokasi < cari {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if space.Lokasi == "" {
		fmt.Println("Co-working space tidak ditemukan.")
	}
}

func UrutkanSpace() {
	var pilihan int
	fmt.Println("Pilih pengurutan berdasarkan:")
	fmt.Println("1. Berdasarkan Harga Sewa ")
	fmt.Println("2. Berdasarkan Rating Tertinggi")
	fmt.Println("0. Kembali")
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 0:
		ManajemenPengguna()
	case 1:
		selectionSort()
	case 2:
		insertionSort()
	default:
		fmt.Println("----------------------------------------")
		fmt.Println("Pilihan tidak valid!")
		fmt.Println("----------------------------------------")
	}
}

func selectionSort() {
	var min int
	for i := 0; i < len(spaces); i++ {
		min = i
		for j := i + 1; j < len(spaces); j++ {
			if spaces[j].Harga < spaces[min].Harga {
				min = j
			}
		}
		spaces[i], spaces[min] = spaces[min], spaces[i]
	}

	fmt.Println("--------------------------------------------------------")
	fmt.Println("Space berhasil diurutkan berdasarkan harga sewa!")
	fmt.Println("--------------------------------------------------------")
}

func insertionSort() {
	for i := 1; i < len(spaces); i++ {
		j := i
		for j > 0 && rataRataRating(spaces[j]) > rataRataRating(spaces[j-1]) {
			spaces[j], spaces[j-1] = spaces[j-1], spaces[j]
			j--
		}
	}

	fmt.Println("----------------------------------------------------")
	fmt.Println("Spaces berhasil diurutkan berdasarkan rating tertinggi!")
	fmt.Println("----------------------------------------------------")
}

func rataRataRating(s Space) float64 {
	if len(s.Ulasan) == 0 {
		return 0
	}
	var total float64
	for i := 0; i < len(s.Ulasan); i++ {
		total += s.Ulasan[i].Rating
	}
	return total / float64(len(s.Ulasan))
}

func RatingSpace() {
	var menu int
	fmt.Println("=============================================")
	fmt.Println("  Selamat datang di Review Co Working Space  ")
	fmt.Println("=============================================")
	fmt.Println("Silahkan pilih menu")
	fmt.Println("[1] Memberikan Review")
	fmt.Println("[2] Mengubah Review")
	fmt.Println("[3] Menghapus Review")
	fmt.Println("[0] Kembali")
	fmt.Print("Pilih menu: ")
	fmt.Scan(&menu)

	switch menu {
	case 0:
		ManajemenPengguna()
	case 1:
		TambahReview()
	case 2:
		UbahReview()
	case 3:
		HapusReview()
	case 4:
		main()
	default:
		fmt.Println("----------------------------------------")
		fmt.Println("Pilihan tidak valid!")
		fmt.Println("----------------------------------------")
	}
}

func TambahReview() {
	var pilihan int
	var r Review

	if len(spaces) == 0 {
		fmt.Println("Daftar co-working space kosong, silahkan tambahkan co-working space terlebih dahulu!")
		return
	}

	TampilSpace()
	fmt.Println()
	fmt.Print("Masukkan nomor co-working space : ")
	fmt.Scan(&pilihan)

	fmt.Print("Masukkan rating (1-5)		: ")
	fmt.Scan(&r.Rating)
	if r.Rating < 1 || r.Rating > 5 {
		fmt.Println("Rating harus diantara 1-5")
		return
	}

	fmt.Print("Masukkan komentar		: ")
	fmt.Scan(&r.Komentar)
	if r.Komentar == "" {
		fmt.Println("Komentar tidak boleh kosong")
		return
	}

	spaces[pilihan-1].Ulasan = append(spaces[pilihan-1].Ulasan, r)
	fmt.Println("----------------------------------------")
	fmt.Println("Review berhasil ditambahkan!")
	fmt.Println("----------------------------------------")

}

func UbahReview() {
	var noSpace, noReview int
	var RatingBaru float64
	var KomentarBaru string

	if len(spaces) == 0 {
		fmt.Println("Daftar co-working space kosong, silakan tambahkan terlebih dahulu!")
		return
	}

	for i := 0; i < len(spaces); i++ {
		fmt.Printf("%d. %s\n", i+1, spaces[i].Nama)
	}

	fmt.Print("Pilih nomor co-working space: ")
	fmt.Scan(&noSpace)

	if noSpace < 1 || noSpace > len(spaces) {
		fmt.Println("Nomor tidak valid.")
		return
	}

	space := &spaces[noSpace-1]

	if len(space.Ulasan) == 0 {
		fmt.Println("Belum ada review untuk co-working space ini.")
		return
	}

	fmt.Println("Daftar review:")
	for i := 0; i < len(space.Ulasan); i++ {
		fmt.Printf("%d. Rating: %.1f | Komentar: %s\n", i+1, space.Ulasan[i].Rating, space.Ulasan[i].Komentar)
	}

	fmt.Print("Pilih nomor review yang ingin diubah: ")
	fmt.Scan(&noReview)

	if noReview < 1 || noReview > len(space.Ulasan) {
		fmt.Println("Nomor review tidak valid.")
		return
	}

	fmt.Print("Masukkan rating baru (1-5): ")
	fmt.Scan(&RatingBaru)
	if RatingBaru < 1 || RatingBaru > 5 {
		fmt.Println("Rating harus di antara 1-5.")
		return
	}

	fmt.Print("Masukkan komentar baru: ")
	fmt.Scan(&KomentarBaru)
	if KomentarBaru == "" {
		fmt.Println("Komentar tidak boleh kosong.")
		return
	}

	space.Ulasan[noReview-1].Rating = RatingBaru
	space.Ulasan[noReview-1].Komentar = KomentarBaru

	fmt.Println("----------------------------------------")
	fmt.Println("Review berhasil diubah!")
	fmt.Println("----------------------------------------")
}

func HapusReview() {
	var noSpace, noReview int
	space := &spaces[noSpace-1]

	if len(spaces) == 0 {
		fmt.Println("Daftar co-working space kosong, silakan tambahkan terlebih dahulu!")
		return
	}
	fmt.Println("Daftar Co-Working Space:")
	for i := 0; i < len(spaces); i++ {
		fmt.Printf("%d. %s\n", i+1, spaces[i].Nama)
	}

	fmt.Print("Masukkan nomor co-working space: ")
	fmt.Scan(&noSpace)
	if noSpace < 1 || noSpace > len(spaces) {
		fmt.Println("Nomor space tidak valid.")
		return
	}

	if len(space.Ulasan) == 0 {
		fmt.Println("Belum ada review untuk co-working space ini.")
		return
	}
	fmt.Println("Daftar review:")
	for i := 0; i < len(space.Ulasan); i++ {
		fmt.Printf("%d. Rating: %.1f | Komentar: %s\n", i+1, space.Ulasan[i].Rating, space.Ulasan[i].Komentar)
	}

	fmt.Print("Masukkan nomor review yang ingin dihapus: ")
	fmt.Scan(&noReview)
	if noReview < 1 || noReview > len(space.Ulasan) {
		fmt.Println("Nomor review tidak valid.")
		return
	}

	index := noReview - 1
	space.Ulasan = append(space.Ulasan[:index], space.Ulasan[index+1:]...)
	fmt.Println("----------------------------------------")
	fmt.Println("Review berhasil dihapus!")
	fmt.Println("----------------------------------------")
}

func FilterFasilitas() {
	var fasilitas string
	fmt.Print("Masukkan fasilitas yang ingin dicari: ")
	fmt.Scan(&fasilitas)

	sesuai := filterByFasilitas(fasilitas)

	if len(sesuai) == 0 {
		fmt.Println("Tidak ada co-working space dengan fasilitas tersebut.")
		return
	}

	fmt.Println("\nCo-working space dengan fasilitas", fasilitas+":")
	for i := 0; i < len(sesuai); i++ {
		space := sesuai[i]
		fmt.Println("Nama:", space.Nama)
		fmt.Println("Lokasi:", space.Lokasi)
		fmt.Println("Fasilitas:", strings.Join(space.Fasilitas, ", "))
		fmt.Println("Harga:", space.Harga)
		fmt.Println("----------------------------------------")
	}
}

func filterByFasilitas(fasilitas string) []Space {
	var hasil []Space

	for i := 0; i < len(spaces); i++ {
		for j := 0; j < len(spaces[i].Fasilitas); j++ {
			if strings.ToLower(spaces[i].Fasilitas[j]) == strings.ToLower(fasilitas) {
				hasil = append(hasil, spaces[i])
				break
			}
		}
	}
	return hasil
}
