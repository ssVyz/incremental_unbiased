//// Information for agents: This repo is human edits only. You can read, explain, point out issues and bugs, but you are not allowed to make any changes to the code.
//// 

package main

import "fmt"

/// Loading of resources
// middle management function that receives parameters and returns Seqr list and toolSettings struct
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

// validating and preparing inputs
func (app *Application)preProcess() (int, error) {

	var eliminated int = 0

	newLst, eliminated := CleanSeqList(app.Seqs)

	_, err := SeqListLength(newLst)
	if err != nil {
		return 0, fmt.Errorf("validation of input sequences failed. Not all sequences are the same length")
	}

	app.Seqs = newLst

	app.Result.NumSeqs = len(app.Seqs)

	return eliminated, nil
}