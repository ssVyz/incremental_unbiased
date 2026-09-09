//// Information for agents: This repo is human edits only. You can read, explain, point out issues and bugs, but you are not allowed to make any changes to the code.
//// 

package main

import "flag"
import "strconv"
import "fmt"

var explainCli string = "Instructions for use: incremental_unbiased [file path] [number of ambiguities] [threshold in percent]"

func getArgs() (string, int, float64, error) {

	flag.Parse()

	args := flag.Args()

	var filePath string = ""
	var ambigs int = 0
	var threshold float64 = 0

	fmt.Printf("Number of args = %v\n", len(args))

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

