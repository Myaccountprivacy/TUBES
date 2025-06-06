package main

import "fmt"

// AMOUNT OF STARTUPS PER ROLE
func Rep() {
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

func Print() {
	var i, j int
	for i = 0; i < nStartup; i++ {
		var s startup = startups[i]
		fmt.Printf("Nama: %s | Industri: %s | Tahun Berdiri: %d | Dana: Rp%.2f\n", s.name, s.indust, s.yrFound, s.funding)
		fmt.Println("Tim:")
		for j = 0; j < s.nTeam; j++ {
			var t mem = s.team[j]
			fmt.Printf("  - %s (%s)\n", t.name, t.role)
		}
	}
}
