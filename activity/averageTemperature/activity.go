package averageTemperature

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("activity-averageTemperature")

const (
	ivTemperature         = "temperature"
	ivPreviousTemperature = "previousTemperature"
	ivTotalCount          = "totalCount"

	ovAverage = "average"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

func init() {
	log.SetLogLevel(logger.InfoLevel)
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	var avg float64
	// do eval
	temp := context.GetInput(ivTemperature).(float64)
	prevTemp := context.GetInput(ivPreviousTemperature).(float64)
	counter := context.GetInput(ivTotalCount).(int)

	avg = getAverage(temp+prevTemp, float64(counter))

	log.Info("Average ", avg)

	context.SetOutput(ovAverage, avg)

	return true, nil
}

func getAverage(s, c float64) float64 {
	return s / float64(c)
}
