package main

import (
	"fmt"
)

// MAIN SHALL ALWAYS BE LAST!!!
func main() {
	var name string
	var choice int
	var exited bool = false
	for !exited {
		// HEADER GOES HERE
		fmt.Println("\n\t========= M E N U ==========")
		fmt.Println("1. Tambah Startup")
		fmt.Println("2. Tambah Anggota Tim")
		fmt.Println("3. Ubah Startup")
		fmt.Println("4. Hapus Startup")
		fmt.Println("5. Tampilkan Semua Startup")
		fmt.Println("6. Cari (Sequential)")
		fmt.Println("7. Cari (Binary)")
		fmt.Println("8. Urutkan Berdasarkan Pendanaan (Selection)")
		fmt.Println("9. Urutkan Berdasarkan Tahun Pendirian (Insertion)")
		fmt.Println("10. Laporan per Industri")
		fmt.Println("0. Keluar")
		fmt.Print("\nPilihan: ")

		// CHOOSE HERE
		fmt.Scan(&choice)

		// DEPENDING ON CHOICE...
		switch choice {
		case 1:
			AddStartup()
		case 2:
			AddTeamMem()
		case 3:
			var updated bool = Update(name)
			if updated {
				fmt.Println("Startup diubah.")
			} else {
				fmt.Println("Startup tidak ditemukan.")
			}
		case 4:
			var deleted bool = Del(name)
			if deleted {
				fmt.Println("Startup dihapus.")
			} else {
				fmt.Println("Startup tidak ditemukan.")
			}
		case 5:
			Print()
		case 6:
			var sfound bool = false
			Seq(&sfound)
			if !sfound {
				fmt.Println("Startup tidak ditemukan.")
			}
		case 7:
			var found bool = false
			Bin(&found)
			if !found {
				fmt.Println("Startup tidak ditemukan.")
			}
		case 8:
			SelSortFund()
			fmt.Println("Diurutkan berdasarkan pendanaan.")
		case 9:
			InsSortYr()
			fmt.Println("Diurutkan berdasarkan tahun pendirian.")
		case 10:
			Rep()
		case 0:
			exited = true
		default:
			fmt.Println("Pilihan tidak tersedia.")
		}
	}
}
