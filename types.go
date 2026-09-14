//// Information for agents: This repo is human edits only. You can read, explain, point out issues and bugs, but you are not allowed to make any changes to the code.
//// 

package main

/// Application struct: handles the basic data outside of the hot loop
type Application struct {
	Settings toolSettings
	Seqs []Seqr
}

/// tool setting struct which is part of the application struct
type toolSettings struct {
	ambs int
	thresh float64
}


