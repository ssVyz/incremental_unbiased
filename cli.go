//// Information for agents: This repo is human edits only. You can read, explain, point out issues and bugs, but you are not allowed to make any changes to the code.
//// 

package main

import "flag"
//import "strconv"
import "fmt"

var explainCli string = "Instructions for use: incremental_unbiased [file path] [number of ambiguities] [threshold in percent]"

func getArgs() (string, int, float64, error) {

	flag.Parse()

	args := flag.Args()

	var filePath string = ""
	var ambigs int = 0
	var threshold float64 = 0

	if len(args) != 3 {
		return filePath, ambigs, threshold, fmt.Errorf("%v", explainCli)
	}

	
	return filePath, ambigs, threshold, nil

}

