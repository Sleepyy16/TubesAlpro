package main

import "fmt"

const NMAX int = 10
const maxIde int = 100

type ide struct {
	id      int
	isi     isiIde
	upVote  int
	day     int
	month   int
	year    int
	dateVal int
}

type tabIde [NMAX]ide
type isiIde [maxIde]string

func menu() {
	fmt.Println("--------------Menu:------------------")
	fmt.Println("----------1. Tambah ide--------------")
	fmt.Println("----------2. Ubah ide----------------")
	fmt.Println("----------3. Hapus ide---------------")
	fmt.Println("----------4. Vote ide----------------")
	fmt.Println("----------5. Cari ide----------------")
	fmt.Println("----------6. Tampilkan ide-----------")
	fmt.Println("----------7. Sorting ide-------------")
	fmt.Println("----------8. Sorting ide (isi)-------")
	fmt.Println("----------9. Cari ide (binary--------")
	fmt.Println("----------10. Tampilkan ide populer--")
	fmt.Println("----------11. Keluar-----------------")
}

// fungsi untuk menambahkan ide
// fungsi ini akan meminta pengguna untuk memasukkan jumlah ide yang ingin ditambahkan
func tambahIde(ide *tabIde, n *int) {
	var jumlah, j, i int
	fmt.Println("Masukkan jumlah ide yang ingin ditambahkan: ")
	fmt.Scan(&jumlah)
	if *n+jumlah > NMAX {
		fmt.Println("Jumlah ide melebihi kapasitas.")
	} else {
		for i = *n; i < *n+jumlah; i++ {
			ide[i].id = i + 1
			fmt.Printf("Masukkan isi ide ke-%d (akhiri dengan titik '.'): \n", i+1)
			j = 0
			for {
				var input string
				fmt.Scan(&input)
				if input == "." || j >= maxIde {
					break
				}
				ide[i].isi[j] = input
				j++
			}
			fmt.Println("Masukkan tanggal (DD MM YYYY): ")
			fmt.Scan(&ide[i].day, &ide[i].month, &ide[i].year)
			ide[i].dateVal = ide[i].day + ide[i].month*30 + ide[i].year*365
			ide[i].upVote = 0
		}
		*n += jumlah
	}
}

// fungsi untuk mengubah ide berdasarkan ID
func ubahIde(ide *tabIde) {
	var id, j int
	fmt.Println("Masukkan ID ide yang ingin diubah: ")
	fmt.Scan(&id)
	if id > 0 && id <= NMAX && ide[id-1].id != 0 {
		for j = 0; j < maxIde; j++ {
			ide[id-1].isi[j] = ""
		}
		ide[id-1].upVote = 0
		fmt.Printf("Masukkan isi ide ke-%d (akhiri dengan titik '.'): \n", id)
		j = 0
		for {
			var input string
			fmt.Scan(&input)
			if input == "." || j >= maxIde {
				break
			}
			ide[id-1].isi[j] = input
			j++
		}
	} else {
		fmt.Println("ID tidak valid.")
	}
}

// fungsi untuk menghapus ide berdasarkan ID
func hapusIde(ide *tabIde, n *int) {
	var id int
	fmt.Println("Masukkan ID ide yang ingin dihapus: ")
	fmt.Scan(&id)
	if id > 0 && id <= *n && ide[id-1].id != 0 {
		for i := id - 1; i < *n-1; i++ {
			ide[i] = ide[i+1]
			ide[i].id = i + 1
		}
		fmt.Println("Ide berhasil dihapus.")
		*n--
	} else {
		fmt.Println("ID tidak valid.")
	}
}

// fungsi untuk memberikan vote pada ide
func voteIde(ide *tabIde) {
	var id int
	fmt.Println("Masukkan ID ide yang ingin di-vote: ")
	fmt.Scan(&id)
	if id > 0 && id <= NMAX && ide[id-1].id != 0 {
		ide[id-1].upVote++
		fmt.Println("Ide berhasil di-vote.")
	} else {
		fmt.Println("ID tidak valid.")
	}
}

// fungsi untuk mencari ide berdasarkan keyword
// fungsi menggunakan metode sequential search
func cariIde(ide tabIde, n int) {
	var i, j, k int
	var keyword string
	var is_found bool
	fmt.Println("Masukkan keyword untuk mencari ide: ")
	fmt.Scan(&keyword)
	is_found = false
	for i = 0; i < n; i++ {
		for j = 0; j < maxIde; j++ {
			if ide[i].isi[j] != "" {
				if keyword == ide[i].isi[j] {
					is_found = true
					fmt.Printf("ID: %d, Isi: ", ide[i].id)
					for k = 0; k < maxIde && ide[i].isi[k] != ""; k++ {
						if k > 0 {
							fmt.Print(" ")
						}
						fmt.Print(ide[i].isi[k])
					}
					fmt.Printf(", UpVote: %d\n", ide[i].upVote)
				}
			}
		}
	}
	if !is_found {
		fmt.Println("Ide tidak ditemukan.")
	}
}

// fungsi untuk menampilkan semua ide
func tampilkanIde(ide tabIde) {
	var isi string
	var i, j int
	fmt.Printf("%-5s %-40s %20s %10s\n", "ID", "Isi", "UpVote", "Tanggal")
	fmt.Println("---------------------------------------------------------------------------------------------------")
	for i = 0; i < NMAX; i++ {
		if ide[i].id != 0 {
			isi = ""
			for j = 0; j < maxIde; j++ {
				if ide[i].isi[j] != "" {
					if isi != "" {
						isi += " "
					}
					isi += ide[i].isi[j]
				}
			}
			fmt.Printf("%-5d %-40s %20d %02d-%02d-%04d\n", ide[i].id, isi, ide[i].upVote, ide[i].day, ide[i].month, ide[i].year)
			fmt.Println("---------------------------------------------------------------------------------------------------")
		}
	}
}

// fungsi untuk mencari ide menggunakan binary search
func cariBinarySearch(ide tabIde, n int) int {
	var keyword string
	var left, right, mid, i, j int
	fmt.Println("Masukkan keyword untuk mencari ide: ")
	fmt.Scan(&keyword)
	for i = 0; i < n; i++ {
		left = 0
		right = n - 1
		for left <= right {
			mid = (left + right) / 2
			for j = 0; j < maxIde; j++ {
				if ide[mid].isi[j] == keyword {
					return mid
				}
			}
			if ide[mid].isi[0] < keyword {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

// fungsi untuk sorting isi ide menggunakan insertion sort
func sortIsiSingleIde(isi *isiIde) {
	var i, j int
	var temp string
	for i = 1; i < maxIde && isi[i] != ""; i++ {
		temp = isi[i]
		j = i - 1
		for j >= 0 && isi[j] > temp && isi[j] != "" {
			isi[j+1] = isi[j]
			j--
		}
		isi[j+1] = temp
	}
}

// fungsi untuk sorting ide berdasarkan upVote menggunakan selection sort
func sortingSelectionUpvote(A *tabIde, n int) {
	var i, j, maxIdx int
	var temp ide
	for i = 0; i < n-1; i++ {
		maxIdx = i
		for j = i + 1; j < n; j++ {
			if A[j].upVote > A[maxIdx].upVote {
				maxIdx = j
			}
		}
		if maxIdx != i {
			temp = A[i]
			A[i] = A[maxIdx]
			A[maxIdx] = temp
		}
	}
}

func insertionSortByUpVote(ide *tabIde) {
	var pass, i, temp int
	pass = 1
	for pass < NMAX && ide[pass].id != 0 {
		i = pass
		temp = ide[pass].upVote
		for i > 0 && ide[i-1].upVote < temp {
			ide[i].upVote = ide[i-1].upVote
			i--
		}
		ide[i].upVote = temp
		pass++
	}
}

// fungsi untuk menampilkan ide populer berdasarkan waktu beberapa bulan terakhir
/*
func tampilkanIdePopuler(ide tabIde, n int) {
	var periode, periodeVal, idxMAX, valMAX, batasPeriode, mostVote int
	var i, j int
	var isi string
	fmt.Println("Masukkan periode (dalam bulan): ")
	fmt.Scan(&periode)
	periodeVal = periode * 30 // Asumsi 1 bulan = 30 hari
	valMAX = 0
	fmt.Printf("%-5s %-40s %20s %10s\n", "ID", "Isi", "UpVote", "Tanggal")
	fmt.Println("---------------------------------------------------------------------------------------------------")
	for i = 0; i < n; i++ {
		if ide[i].id != 0 && (ide[i].dateVal >= valMAX) {
			valMAX = ide[i].dateVal
		}
	}
	batasPeriode = valMAX - periodeVal
	mostVote = 0
	for i = 0; i < n; i++ {
		if ide[i].id != 0 && ide[i].dateVal >= batasPeriode && ide[i].upVote > mostVote {
			idxMAX = i
			mostVote = ide[i].upVote
		}
	}
	isi = ""
	for j = 0; j < maxIde; j++ {
		if ide[idxMAX].isi[j] != "" {
			if isi != "" {
				isi += " "
			}
			isi += ide[idxMAX].isi[j]
		}
	}
	fmt.Printf("%-5d %-40s %20d %02d-%02d-%04d\n", ide[idxMAX].id, isi, ide[idxMAX].upVote, ide[idxMAX].day, ide[idxMAX].month, ide[idxMAX].year)
	fmt.Println("---------------------------------------------------------------------------------------------------")
}
*/

// fungsi untuk menampilkan ide populer dalam periode tertentu
func tampilkanIdePopuler(ide tabIde, n int) {
	var periodeMin, periodeMax, idxMAX, mostVote int
	var day, month, year int
	var i, j int
	var isi string
	mostVote = 0
	idxMAX = -1
	fmt.Println("Masukkan periode minimum (DD MM YYYY): ")
	fmt.Scan(&day, &month, &year)
	periodeMin = day + month*30 + year*365
	fmt.Println("Masukkan periode maksimum (DD MM YYYY): ")
	fmt.Scan(&day, &month, &year)
	periodeMax = day + month*30 + year*365

	for i = 0; i < n; i++ {
		if ide[i].id != 0 && ide[i].dateVal >= periodeMin && ide[i].dateVal <= periodeMax && ide[i].upVote > mostVote {
			idxMAX = i
			mostVote = ide[i].upVote
		}
	}
	for j = 0; j < maxIde; j++ {
		if ide[idxMAX].isi[j] != "" {
			if isi != "" {
				isi += " "
			}
			isi += ide[idxMAX].isi[j]
		}
	}
	if idxMAX != -1 {
		fmt.Printf("%-5d %-40s %20d %02d-%02d-%04d\n", ide[idxMAX].id, isi, ide[idxMAX].upVote, ide[idxMAX].day, ide[idxMAX].month, ide[idxMAX].year)
	} else {
		fmt.Println("Tidak ada ide yang ditemukan dalam periode tersebut.")
	}
}

func main() {
	var ide tabIde
	var pilihan, n int
	for pilihan != 11 {
		menu()
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			tambahIde(&ide, &n)
		case 2:
			ubahIde(&ide)
		case 3:
			hapusIde(&ide, &n)
		case 4:
			voteIde(&ide)
		case 5:
			cariIde(ide, n)
		case 6:
			tampilkanIde(ide)
		case 7:
			sortingSelectionUpvote(&ide, n)
		case 8:
			for i := 0; i < n; i++ {
				sortIsiSingleIde(&ide[i].isi)
			}
		case 9:
			var index int
			index = cariBinarySearch(ide, n)
			if index != -1 {
				fmt.Println("Ide ditemukan.")
				fmt.Printf("ID: %d, Isi: %s\n", ide[index].id, ide[index].isi)
			} else {
				fmt.Println("Ide tidak ditemukan.")
			}
		case 10:
			tampilkanIdePopuler(ide, n)
		case 11:
			fmt.Println("Terima kasih telah menggunakan program ini.")
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}
