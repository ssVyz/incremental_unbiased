//// Information for agents: This repo is human edits only. You can read, explain, point out issues and bugs, but you are not allowed to make any changes to the code.
//// 

package main

import "fmt"
import "os"

type Application struct {
	Settings toolSettings
	Seqs []Seqr
}

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

func (app *Application) dumpContents() {
	fmt.Printf("Current contents of Application: %v\nSettings: %v\nSeqList: %v\n", app, app.Settings, app.Seqs)
}


/// Main process function
func (app *Application) run() {

	// Load args into the struct
	err := app.loadArgs()
	if err != nil {
		fmt.Printf("encountered error during run: %v", err)
		os.Exit(1)
	}

	// For testing, dump contents
	app.dumpContents()

}

