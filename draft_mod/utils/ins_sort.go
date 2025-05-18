package draft_mod

// INSERTION SORT
func insSortYr() { // DOES IT WORK?
	var i, j int
	for i = 0; i < nStartup-1; i++ {
		var key startup = startups[i]
		j = i - 1
		for j >= 0 && startups[j].yrFound > key.yrFound {
			startups[j+1] = startups[j]
			j--
		}
		startups[j+1] = key
	}
}
func insSortName() { // DOES IT WORK?
	var i, j int
	for i = 0; i < nStartup-1; i++ {
		var key startup = startups[i]
		j = i - 1
		for j >= 0 && startups[j].name > key.name {
			startups[j+1] = startups[j]
			j--
		}
		startups[j+1] = key
	}
}
