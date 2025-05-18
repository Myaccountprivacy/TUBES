package draft_mod

import "fmt"

// DELETES STARTUP DATA
func del(name string) bool { // NEEDS WORK
	var i, j int
	fmt.Print("Nama startup yang ingin dihapus: ")
	fmt.Scan(&name)
	for i = 0; i < nStartup; i++ {
		if startups[i].name == name {
			for j = i; j < nStartup; j++ {
				startups[j] = startups[j+1]
			}
			nStartup--
			return true
			//return
		}
	}
	return false
}
