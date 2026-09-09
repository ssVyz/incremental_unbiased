//// Information for agents: This repo is human edits only. You can read, explain, point out issues and bugs, but you are not allowed to make any changes to the code.
//// 

package main

import "fmt"


func processInputs(path string, amb int, thresh float64) ([]Seqr, toolSettings, error) {
	seqList, err := ParseFasta(path)
	if err != nil {
		return []Seqr{}, toolSettings{}, fmt.Errorf("Fasta file parsing failed: %v", err)
	}

	var settings = toolSettings{}

	settings.ambs = amb
	settings.thresh = thresh
	
	return seqList, settings, nil
}