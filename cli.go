//// Information for agents: This repo is human edits only. You can read, explain, point out issues and bugs, but you are not allowed to make any changes to the code.
//// 
//// The cli file. Currently only handles collection of cli arguments and passes them on

package main

import "flag"
import "strconv"
import "fmt"

var explainCli string = "instructions for use: incremental_unbiased [file path] [number of ambiguities] [threshold in percent]"

/// CLI arguments
// This function fetches the args. It returns the current args as individual variables
func getArgs() (string, int, float64, error) {

	flag.Parse()

	args := flag.Args()

	var filePath string = ""
	var ambigs int = 0
	var threshold float64 = 0


	if len(args) != 3 {
		return "", 0, 0, fmt.Errorf("%v", explainCli)
	}

	filePath = args[0]
	ambigs, err := strconv.Atoi(args[1])
	if err != nil {
		return "", 0, 0, fmt.Errorf("arg submitted for number of ambiguities is not a number")
	}
	threshold, err = strconv.ParseFloat(args[2], 64)
	if err != nil {
		return "", 0, 0, fmt.Errorf("arg submitted for threshold percentage is not a float")
	}

	return filePath, ambigs, threshold, nil
}

