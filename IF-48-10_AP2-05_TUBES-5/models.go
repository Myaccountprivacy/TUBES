package main

// STRUCTS
type mem struct {
	name string
	role string
}
type startup struct {
	yrFound, nTeam int
	name, indust   string
	funding        float64
	team           [nmax_team]mem
}
