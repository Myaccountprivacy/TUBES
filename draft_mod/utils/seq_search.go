package draft_mod

import "fmt"

// SEQUENTIAL SEARCH
func seq() {
	var i int
	var name string
	fmt.Print("Nama Startup: ")
	fmt.Scan(&name)
	for i = 0; i < nStartup; i++ {
		if startups[i].name == name {
			fmt.Println("Ditemukan:", startups[i].name)
			return
		}
	}
	fmt.Println("Tidak ditemukan.")
}
