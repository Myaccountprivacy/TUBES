package funcs

import "fmt"

// ADD STARTUP DATA
func AddStartup() {
	var s startup
	if nStartup >= nmax_comp {
		fmt.Println("Maksimal startup tercapai.")
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

// FUNCTIONS (UPDATE, DELETE)
func Update(name string) bool { // NEEDS WORK
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
func Del(name string) bool { // NEEDS WORK
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
		}
	}
	return false
}

// ADD TEAM MEMBER
func AddTeamMem() {
	var name string
	var i int
	var found bool
	fmt.Print("Nama Startup: ")
	fmt.Scan(&name)
	for i = 0; i < nStartup && !found; i++ {
		if startups[i].name == name {
			found = true
			if startups[i].nTeam >= nmax_team {
				fmt.Println("Tim sudah penuh.")
				return
			} else {
				var member mem
				fmt.Print("Nama Anggota: ")
				fmt.Scan(&member.name)
				fmt.Print("Peran: ")
				fmt.Scan(&member.role)
				startups[i].team[startups[i].nTeam] = member
				startups[i].nTeam++
				fmt.Println("Anggota ditambahkan.")
			}
		}
	}
	if !found {
		fmt.Println("Startup tidak ditemukan.")
	}
}
