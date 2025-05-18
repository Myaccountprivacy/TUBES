package funcs

// SELECTION SORT
func SelSortFund() { // DOES IT WORK?
	var i, j int
	for i = 0; i < nStartup-1; i++ {
		var min_idx int = i
		for j = i + 1; j < nStartup; j++ {
			if startups[j].funding < startups[min_idx].funding {
				min_idx = j
			}
		}
		startups[i] = startups[min_idx]
		startups[min_idx] = startups[i]
		// startups[i],startups[min_idx]=startups[min_idx],startups[i]
	}
}

// INSERTION SORT
func InsSortYr() { // DOES IT WORK?
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
func InsSortName() { // DOES IT WORK?
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
