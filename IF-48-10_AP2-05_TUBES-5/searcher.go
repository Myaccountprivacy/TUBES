package main

import "fmt"

// SEQUENTIAL SEARCH
func Seq(sfound *bool) {
	var i int
	var name string
	fmt.Print("Nama Startup: ")
	fmt.Scan(&name)
	for i = 0; i < nStartup; i++ {
		if startups[i].name == name {
			fmt.Println("Ditemukan:", startups[i].name)
			*sfound = true
		}
	}
}

// BINARY SEARCH
func Bin(found *bool) {
	var name string
	fmt.Print("Nama Startup: ")
	fmt.Scan(&name)
	InsSortName()
	var lo int = 0
	var hi int = nStartup - 1
	for lo <= hi {
		var mid int = (lo + hi) / 2
		if startups[mid].name == name {
			fmt.Println("Ditemukan:", startups[mid].name)
			*found = true
			lo = hi + 1
		} else if startups[mid].name < name {
			lo = mid + 1
		} else if startups[mid].name > name {
			hi = mid - 1
		}
	}
}
