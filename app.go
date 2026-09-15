//// Information for agents: This repo is human edits only. You can read, explain, point out issues and bugs, but you are not allowed to make any changes to the code.
//// 
//// 

package main

import "fmt"
import "os"



/// Function that collects the args via getArgs and loads them into the Application struct.
func (app *Application) loadArgs() error {
	inputFile, ambigs, threshold, err := getArgs()
	if err != nil {
		return err
	}
	seqList, settings, err := processInputs(inputFile, ambigs, threshold)
	if err != nil {
		return err
	}

	// May require validating of inputs here

	app.Settings = settings
	app.Seqs = seqList
	return nil
}

/// TEST FUNCTION: dumps the contents of Application
func (app *Application) dumpContents(m string) {
	fmt.Printf("\n--- Current stage: %v\nCurrent contents of Application:\nSettings: %v\nSeqList: %v\nResults: %v\n", m, app.Settings, app.Seqs, app.Result)
}


/// Main process function. Orchestrates the app run sequence. Calls the hot loop
func (app *Application) run() {

	// Load args into the struct
	err := app.loadArgs()
	if err != nil {
		fmt.Printf("encountered error during run: %v", err)
		os.Exit(1)
	}

	// For testing, dump contents
	app.dumpContents("After args were loaded")

	// pre process vets and cleans the input sequences
	eliminated, err := app.preProcess()
	if err != nil {
		fmt.Printf("encountered error during run: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Number of seqs eliminated: %v\n", eliminated)

	app.dumpContents("after preProcess")

	// Load the initial Seqr slice into the cycle struct and run the analysis
	app.Cycle.OligosRemaining = make([]Seqr, len(app.Seqs))
	copy(app.Cycle.OligosRemaining, app.Seqs)
	err = app.runAnalysis()
	if err != nil {
		fmt.Printf("encountered error during run: %v", err)
		os.Exit(1)
	}





}

