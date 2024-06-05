package main

import (
	"fmt"
)

type ingfo struct {
	nama                                    string
	win, lose, seri, goal, kebobolan, score int
	anggota                                 tabpem
}
type pemain struct {
	nama             string
	no, umur, tinggi int
}

const NMAX int = 100

type tabTim [NMAX]ingfo
type tabpem [NMAX]pemain

func main() {
	var data tabTim
	var nData, inputan, batas int
	var nama string
	var tim string

	//fmt.Println("masukkan banyak tim")

	score(&data, nData) // memberi nilai ke dalam struct score

	for inputan != 9 {
		fmt.Println("==============")
		fmt.Println("pilihan menu")
		fmt.Println("--------------")
		fmt.Println("1. input data tim")
		fmt.Println("2. highscore")
		fmt.Println("3. delete")
		fmt.Println("4. tampilan data")
		fmt.Println("5. tambah anggota")
		fmt.Println("6. tampilakan anggota")
		fmt.Println("7. edit data tim")
		fmt.Println("8. cari dengan nama anggota")
		fmt.Println("9. exit")
		fmt.Println("==============")

		sort(&data, nData)

		fmt.Println("masukan pilihan:")
		fmt.Scan(&inputan) // masukkan untuk memilih pilihan menu
		if inputan == 9 {  // mencetak pesan terakhir program
			fmt.Println("goodbye")
		}
		fmt.Println()
		if inputan == 1 {
			// untuk memasukkan data tim
			tambahdata(&data, &nData)
			score(&data, nData)
			sort(&data, nData)
		} else if inputan == 2 {
			// untuk mencari tim dengan score tertinggi
			var pos int
			pos = scoreT(data, nData)
			if pos != -1 {
				fmt.Println("tim dengan score tertinggi:")
				fmt.Printf("%-10s\t\t\t menang:%v kalah:%v seri:%v goal:%v kebobolan:%v score:%v\n", data[pos].nama, data[pos].win, data[pos].lose, data[pos].seri, data[pos].goal, data[pos].kebobolan, data[pos].score)
			} else {
				fmt.Println("Tidak ada tim yang ditemukan")
			}
			fmt.Println()
		} else if inputan == 3 {
			// melakukan hapus data tim
			if nData == 0 {
				fmt.Println("tidak ada data")
				fmt.Println()
			} else {
				fmt.Println("inputkan nama tim")
				fmt.Scan(&nama)
				deleteTeam(&data, &nData, nama)
			}
		} else if inputan == 4 {
			// untuk menampilkan data tim
			tampil(data, nData)
		} else if inputan == 5 {
			// untuk memasukkan data anggota tim
			fmt.Println("masukkan nama tim")
			fmt.Scan(&tim)

			fmt.Println("masukkan banyak anggota")
			fmt.Scan(&batas)
			anggota(&data, nData, batas, tim)
		} else if inputan == 6 {
			// melakukan cetak data anggota tim
			cetakAnggota(data, nData, batas)
		} else if inputan == 7 {
			// melakukan edit data tim
			if nData == 0 {
				fmt.Println("tidak ada data")
				fmt.Println()
			} else {
				fmt.Println("inputkan nama tim")
				fmt.Scan(&nama)
				edit(&data, nama, nData)
			}
		} else if inputan == 8 {
			// melakukan pencarian tim menggunakan data anggota tim
			cariPanggota(data, nData, batas)
		} else if inputan > 9 {
			// mencetak pesan jika pilihan tidak sesuai dengan yang ada di menu
			fmt.Println("tidak ada pilihan yang sesuai")
		}
	}

}

func sort(a *tabTim, n int) {
	// untuk melakukan sorting sesuai score tertinggi
	var i, pas int
	var temp ingfo

	pas = 1
	for pas <= n-1 {
		i = pas
		temp = a[pas]
		for i > 0 && temp.score > a[i-1].score {
			a[i] = a[i-1]
			i--
		}
		a[i] = temp
		pas++
	}
}

func scoreT(a tabTim, n int) int {
	//mencari score tertinggi
	var i, tertinggi, hasil int
	for i = 0; i < n; i++ {
		if a[i].score < a[i+1].score {
			tertinggi = i + 1
		} else {
			tertinggi = i
		}
	}
	if a[tertinggi].score == 0 {
		hasil = -1
	}
	return hasil

}

func binarySearch(a tabTim, n int, nama string) int {
	// melakukan search binary
	var left, right, mid, hasil int
	left = 0
	right = n - 1

	// memastikan nilai bahwa ada tim dengan nilai tertinggi
	for left <= right {
		mid = (left + right) / 2
		if a[mid].nama == nama {
			hasil = mid
		}

		if a[mid].nama != nama {
			left = mid + 1
			right = mid - 1
		}

	}
	return hasil

}

func tampil(a tabTim, n int) {
	// mencetak data dari tim
	var i int
	if n == 0 {
		fmt.Println("tidak ada data")
		fmt.Println()
	} else {
		fmt.Println("Data terurut dari score terbanyak:")
		for i = 0; i < n; i++ {
			fmt.Printf("%-10s\t\t\t menang:%v kalah:%v seri:%v goal:%v kebobolan:%v score:%v\n", a[i].nama, a[i].win, a[i].lose, a[i].seri, a[i].goal, a[i].kebobolan, a[i].score)
		}
		fmt.Println()
	}
}

func edit(a *tabTim, x string, n int) {
	// procedure mengedit data tim
	var pilihan, win, lose, seri, goal, kebobolan, hasil int
	var nama string
	hasil = binarySearch(*a, n, x)
	if a[hasil].nama == x {
		fmt.Println("pilih data yang ingin diganti:")
		fmt.Println("1. ganti nama tim")
		fmt.Println("2. ganti win tim")
		fmt.Println("3. ganti lose tim")
		fmt.Println("4. ganti seri tim")
		fmt.Println("5. ganti goal tim")
		fmt.Println("6. ganti kebobolan tim")
		fmt.Println("7. exit")
		fmt.Println()
		fmt.Println("masukkan pilihan:")
		fmt.Scan(&pilihan)
		if pilihan == 1 { // pilihan untuk mengganti nama tim
			fmt.Println("masukkan nama baru:")
			fmt.Scan(&nama)
			a[hasil].nama = nama
		} else if pilihan == 2 { // pilihan untuk mengganti nilai menang tim
			fmt.Println("masukkan nilai menang:")
			fmt.Scan(&win)
			var test int
			test = 0
			a[hasil].score = test
			a[hasil].win = win
			score(a, n)
		} else if pilihan == 3 { // pilihan untuk mengganti nilai kalah tim
			fmt.Println("masukkan nilai kalah:")
			fmt.Scan(&lose)
			a[hasil].lose = lose
		} else if pilihan == 4 { // pilihan untuk menggati nilai seri tim
			fmt.Println("masukkan nilai seri:")
			fmt.Scan(&seri)
			var test int
			test = 0
			a[hasil].score = test
			a[hasil].seri = seri
			score(a, n)
		} else if pilihan == 5 { // pilihan untuk mengganti nilai goal tim
			fmt.Println("masukkan nilai goal:")
			fmt.Scan(&goal)
			a[hasil].goal = goal
		} else if pilihan == 6 { // pilihan untuk mengganti nilai kebobolan tim
			fmt.Println("masukkan nilai kebobolan:")
			fmt.Scan(&kebobolan)
			a[hasil].kebobolan = kebobolan
		} else if pilihan == 7 {

		}
	} else {
		fmt.Println() // jika data tim yang dimasukan tidak ada
		fmt.Println("data tidak ditemukan")
		fmt.Println()

	}

}

func deleteTeam(a *tabTim, teamCount *int, name string) {
	// procedure menghapus data tim
	var pilihan, i, j int

	for i = 0; i < *teamCount; i++ {
		if a[i].nama == name {
			fmt.Println("apakah kamu yakin ingin menghapus data?")
			fmt.Println("1. ya")
			fmt.Println("2. tidak")
			fmt.Println("------------------")
			fmt.Println("masukkan pilihan")
			fmt.Scan(&pilihan)
			if pilihan == 1 {
				// Geser elemen setelah tim yang dihapus ke kiri
				for j = i; j < *teamCount-1; j++ {
					a[j] = a[j+1]
				}
				// Kosongkan elemen terakhir
				a[*teamCount-1] = ingfo{}
				*teamCount--
				fmt.Println()
				fmt.Println("Tim berhasil dihapus.")
				fmt.Println()

			} else if pilihan == 2 {
				fmt.Println("tim tidak jadi dihapus")

			}
		} else if a[i].nama != name {
			fmt.Println()
			fmt.Println("Tim tidak ditemukan.")
		}
	}

}

func score(a *tabTim, n int) {
	// procedure menghitung score dari tim
	var i int
	for i = 0; i < n; i++ {
		a[i].score = (a[i].win * 2) + (a[i].seri * 1)
	}

}

func anggota(a *tabTim, n int, batas int, tim string) {
	// procedure menambah anggota
	var i, j int

	for i = 0; i < n; i++ {
		if a[i].nama == tim {
			fmt.Println("masukkan data anggota(nama, no punggung, umur, dan tinggi)")
			for j = 0; j < batas; j++ {
				fmt.Scan(&a[i].anggota[j].nama, &a[i].anggota[j].no, &a[i].anggota[j].umur, &a[i].anggota[j].tinggi)
			}
			fmt.Println()
		} else if i > n {
			fmt.Println()
			fmt.Println("tim tidak ditemukan")
			fmt.Println()

		}
	}

}

func cetakAnggota(a tabTim, n int, batas int) {
	// procedure menampilkan data anggota
	var tim string
	var i, j int

	if n == 0 {
		fmt.Println("tidak ada data")
	} else {
		fmt.Println("masukkan nama tim yang ingin dicetak anggotanya")
		fmt.Scan(&tim)
		for i = 0; i < n; i++ {
			if a[i].nama == tim {
				fmt.Println("data anggota:")
				for j = 0; j < batas; j++ {
					fmt.Printf("%-10s nomor punggung:%v umur:%v tinggi:%v\n", a[i].anggota[j].nama, a[i].anggota[j].no, a[i].anggota[j].umur, a[i].anggota[j].tinggi)
				}
				fmt.Println()

			} else if i > n {
				fmt.Println("tim tidak ditemukan")
			}
		}
	}

}

func tambahdata(a *tabTim, n *int) {
	// procedure untuk menambahkan data tim
	var batas, i int

	fmt.Println("masukkan banyak tim:")
	fmt.Scan(&batas)
	for i = *n; i < *n+batas; i++ {
		fmt.Println("nama tim, menang, kalah, seri, goal, kebobolan tim(note: data jangan lebih dari 5 data):")
		fmt.Scan(&a[i].nama, &a[i].win, &a[i].lose, &a[i].seri, &a[i].goal, &a[i].kebobolan)
	}
	*n = *n + batas
}

func cariPanggota(a tabTim, n int, batas int) {
	// cari tim menggunakan nama anggota
	var j int
	var nama string
	var i int

	// memasukkan nama anggota dari tim yang ingin dicari
	fmt.Println("masukkan nama anggota yang ingin dicari timnya:")
	fmt.Scan(&nama)

	fmt.Println("data tim dari anggota pemain")
	for i = 0; i < n; i++ { // untuk mencari array tabtim
		for j = 0; j < batas; j++ { // untuk mencari array tabpem yang ada didalam tabtim
			if a[i].anggota[j].nama == nama {
				fmt.Printf("%-10s\tmenang:%v kalah:%v seri:%v goal:%v kebobolan:%v score:%v\n", a[i].nama, a[i].win, a[i].lose, a[i].seri, a[i].goal, a[i].kebobolan, a[i].score)
				for j = 0; j < batas; j++ { // untuk mencetak semua anggota dari tim
					fmt.Printf("%-10s\t nomor punggung:%v umur:%v tinggi:%v\n", a[i].anggota[j].nama, a[i].anggota[j].no, a[i].anggota[j].umur, a[i].anggota[j].tinggi)

				}

			}
		}
	}
}
