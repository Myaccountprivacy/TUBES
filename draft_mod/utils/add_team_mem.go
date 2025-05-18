package draft_mod

import "fmt"

// ADD TEAM MEMBER
func addTeamMem() {
	var name string
	var i int
	fmt.Print("Nama Startup: ")
	fmt.Scan(&name)
	for i = 0; i < nStartup; i++ {
		if startups[i].name == name {
			if startups[i].nTeam >= nmax_team {
				fmt.Println("Tim sudah penuh.")
				return
			}
			var member mem
			fmt.Print("Nama Anggota: ")
			fmt.Scan(&member.name)
			fmt.Print("Peran: ")
			fmt.Scan(&member.role)
			startups[i].team[startups[i].nTeam] = member
			startups[i].nTeam++
			fmt.Println("Anggota ditambahkan.")
			return
		}
	}
	fmt.Println("Startup tidak ditemukan.")
}
