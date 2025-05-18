package main

import (
	"draft_mod_2/funcs"
	"fmt"
)

// MAIN SHALL ALWAYS BE LAST!!!
func main() {
	var name string
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
			funcs.AddStartup()
		case 2:
			funcs.AddTeamMem()
		case 3:
			var updated bool = funcs.Update(name)
			if updated {
				fmt.Println("Startup diubah.")
			} else {
				fmt.Println("Startup tidak ditemukan.")
			}
		case 4:
			var deleted bool = funcs.Del(name)
			if deleted {
				fmt.Println("Startup dihapus.")
			} else {
				fmt.Println("Startup tidak ditemukan.")
			}
		case 5:
			funcs.Print()
		case 6:
			funcs.Seq()
		case 7:
			funcs.Bin()
		case 8:
			funcs.SelSortFund()
			fmt.Println("Diurutkan berdasarkan pendanaan.")
		case 9:
			funcs.InsSortYr()
			fmt.Println("Diurutkan berdasarkan tahun pendirian.")
		case 10:
			funcs.Rep()
		case 0:
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak ada. Pilih lagi.")
		}
	}
}
