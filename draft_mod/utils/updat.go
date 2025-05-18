package draft_mod

import "fmt"

// UPDATES STARTUP DATA
func update(name string) bool { // NEEDS WORK
	var i int
	fmt.Print("Nama startup yang ingin diubah: ")
	fmt.Scan(&name)
	for i = 0; i < nStartup; i++ {
		if startups[i].name == name {
			var s startup
			fmt.Print("Nama Baru: ")
			fmt.Scan(&s.name)
			fmt.Print("Bidang Usaha Baru: ")
			fmt.Scan(&s.indust)
			fmt.Print("Tahun Berdiri Baru: ")
			fmt.Scan(&s.yrFound)
			fmt.Print("Total Pendanaan Baru: ")
			fmt.Scan(&s.funding)
			s.nTeam = startups[i].nTeam
			startups[i] = s
			return true
		}
	}
	return false
}
