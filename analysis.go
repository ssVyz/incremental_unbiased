//// Information for agents: This repo is human edits only. You can read, explain, point out issues and bugs, but you are not allowed to make any changes to the code.
//// 
//// 


package main


//import "fmt"



func (app *Application) runAnalysis() error {


	/// OUTER HOT LOOP
	for len(app.Cycle.OligosRemaining) > 0 {
		
		_ = app.discoverOneVariant()
		

	}



	return nil

}

func (app *Application) discoverOneVariant() string {
	cyc := &app.Cycle
	maxDepth := app.Settings.ambs
	reqCover := app.Settings.thresh
	cyc.IncrementDepth = 0
	bestHitcount := 0
	maxHitcount := len(cyc.OligosRemaining)
	reqHitcount := int(float64(maxHitcount) * (reqCover/100))
	resultString := ""

	app.writeInitialVariantBuffer()

	// INNER HOT LOOP
	for bestHitcount < reqHitcount && cyc.IncrementDepth <= maxDepth {
		cyc.writeVariantMapAndCount()

	}

	return resultString

}


func (app *Application) writeInitialVariantBuffer() {
	cyc := &app.Cycle

	// prepare the initial variant buffer and copy over the sequences in oligos remaining
	cyc.VariantBuffer = make([]string, 0, len(cyc.OligosRemaining))
	cyc.VariantMap = make(map[string]int)
	for _, rec := range cyc.OligosRemaining {
		cyc.VariantBuffer = append(cyc.VariantBuffer, rec.Seq)
	}

}



func (cyc *AnalysisCycle) writeVariantBuffer() {

	//consensus, _ := MakeConsensus(cyc.OligosRemaining)


}

func (cyc *AnalysisCycle) writeVariantMapAndCount() {
	// load the variant map and initialize entries with 0
	for _, seq := range cyc.VariantBuffer {
		_, ok := cyc.VariantMap[seq]
		if ok {
			continue
		} else {
			cyc.VariantMap[seq] = 0
		}
	}
	// count occurances
	for key, _ := range cyc.VariantMap {
		for _, sequence := range cyc.OligosRemaining {
			if OligoMatch(key, sequence.Seq) {
				cyc.VariantMap[key]++
			}
		}
	}
}



