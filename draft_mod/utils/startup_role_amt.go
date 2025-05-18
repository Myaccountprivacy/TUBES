package draft_mod

import "fmt"

// AMOUNT OF STARTUPS PER ROLE
func rep() {
	var industry [nmax_comp]string
	var count [nmax_comp]int
	var unique int = 0
	var i, j int
	for i = 0; i < nStartup; i++ {
		var found bool = false
		for j = 0; j < unique; j++ {
			if industry[j] == startups[i].indust {
				count[j]++
				found = true
				break
			}
		}
		if !found {
			industry[unique] = startups[i].indust
			count[unique] = 1
			unique++
		}
	}
	fmt.Println("Laporan Jumlah Startup per Bidang Usaha:")
	for i = 0; i < unique; i++ {
		fmt.Printf("%s: %d startup\n", industry[i], count[i])
	}
}
