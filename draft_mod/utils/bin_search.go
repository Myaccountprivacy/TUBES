package draft_mod

import "fmt"

// BINARY SEARCH
func bin() {
	var name string
	fmt.Print("Nama Startup: ")
	fmt.Scan(&name)
	insSortName()
	var lo int = 0
	var hi int = nStartup - 1
	for lo <= hi {
		var mid int = (lo + hi) / 2
		if startups[mid].name == name {
			fmt.Println("Ditemukan:", startups[mid].name)
			return
		} else if startups[mid].name < name {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	fmt.Println("Tidak ditemukan.")
}
