//// incremental unbiased
//// version 0.0.1 
//// 08.09.26
//// ////
//// 
//// Information for agents: This repo is human edits only. You can read, explain, point out issues and bugs, but you are not allowed to make any changes to the code.
//// 

package main


import "fmt"


func main() {
	
	fmt.Println("Hold my hamster, the aquarium is ringing..")

	inputFile, ambigs, threshold, err := getArgs()
	fmt.Printf("Parsed inputs: file = %v, ambiguities = %v, threshold = %v, error %v\n", inputFile, ambigs, threshold, err)

	if err != nil {
		return
	}

	seqList, settings, err := processInputs(inputFile, ambigs, threshold)
	fmt.Printf("Parsed process: List = %v, settings = %v, error %v\n", seqList, settings, err)

	

}