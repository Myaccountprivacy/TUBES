package main

import (
	"draft_mod/utils"
	"fmt"
)

// CONSTANTS FOR ARRAYS
const (
	nmax_team int = 10
	nmax_comp int = 100
)

// STRUCTS
type mem struct {
	name string
	role string
}
type startup struct {
	yrFound, nTeam                   int
	name, indust, investorName, role string
	funding                          float64
	team                             [nmax_team]mem
}

// GLOBAL VARIABLES
var startups [nmax_comp]startup
var nStartup int = 0

// MAIN SHALL ALWAYS BE LAST!!!
func main() {
	var data startup
	for {
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
		var choice int
		fmt.Scan(&choice)

		// DEPENDING ON CHOICE...
		switch choice {
		case 1:
			utils.addStartup()
		case 2:
			utils.addTeamMem()
		case 3:
			var updated bool = utils.update(data.name)
			if updated {
				fmt.Println("Startup diubah.")
			} else {
				fmt.Println("Startup tidak ditemukan.")
			}
		case 4:
			var deleted bool = utils.del(data.name)
			if deleted {
				fmt.Println("Startup dihapus.")
			} else {
				fmt.Println("Startup tidak ditemukan.")
			}
		case 5:
			utils.print()
		case 6:
			utils.seq()
		case 7:
			utils.bin()
		case 8:
			utils.selSortFund()
			fmt.Println("Diurutkan berdasarkan pendanaan.")
		case 9:
			utils.insSortYr()
			fmt.Println("Diurutkan berdasarkan tahun pendirian.")
		case 10:
			utils.rep()
		case 0:
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak ada. Pilih lagi.")
		}
	}
}
