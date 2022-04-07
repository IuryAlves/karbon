package main

import (
	"fmt"
	"os"
	"time"

	mcs "github.com/iuryalves/karbon/multiclusterservice"
	"github.com/iuryalves/karbon/watttime"

)

var (
	token watttime.Token
	selectedRegion string
	belgiumRTEI watttime.RealTimeEmissionsIndex
	kansasRTEI watttime.RealTimeEmissionsIndex
)

func init() {
	token = watttime.Login(os.Getenv("WATTTIME_USER"), os.Getenv("WATTTIME_PASSWORD"))
}

func main() {
	for {
		belgiumRTEI = watttime.Index(token, "BE")
		kansasRTEI = watttime.Index(token, "SPP_KANSAS")

		if belgiumRTEI.Percent < kansasRTEI.Percent {
			selectedRegion = "europe-west1-b/gke-eu"
		} else {
			selectedRegion = "us-central1-b/gke-us"
		}
		mcs.SelectRegion(selectedRegion)
		fmt.Println("Running ....")
		time.Sleep(15 * time.Second)
	}
}
