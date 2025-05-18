package draft_mod

import "fmt"

// PRINTS STARTUP DATA & TEAM MEMBER(S)
func print() {
	var i, j int
	for i = 0; i < nStartup; i++ {
		var s startup = startups[i]
		fmt.Printf("Nama: %s; Industri: %s; Tahun: %d; Dana: Rp%.2f\n", s.name, s.indust, s.yrFound, s.funding)
		fmt.Println("Tim:")
		for j = 0; j < s.nTeam; j++ {
			var t mem = s.team[j]
			fmt.Printf("  - %s (%s)\n", t.name, t.role)
		}
	}
}
