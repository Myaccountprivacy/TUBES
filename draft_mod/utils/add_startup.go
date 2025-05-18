package draft_mod

import "fmt"

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

// ADD A STARTUP
func addStartup() {
	var s startup
	if nStartup >= nmax_comp {
		fmt.Println("Maksimal startup tercapai.")
		return
	}
	fmt.Print("Nama Startup: ")
	fmt.Scan(&s.name)
	fmt.Print("Bidang Usaha: ")
	fmt.Scan(&s.indust)
	fmt.Print("Tahun Berdiri: ")
	fmt.Scan(&s.yrFound)
	fmt.Print("Total Pendanaan: ")
	fmt.Scan(&s.funding)
	s.nTeam = 0
	startups[nStartup] = s
	nStartup++
	fmt.Println("Startup ditambahkan.")
}
