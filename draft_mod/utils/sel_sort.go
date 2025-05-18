package draft_mod

// SELECTION SORT
func selSortFund() { // DOES IT WORK?
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
