package randomNumber

import (
	"math/rand"
	"time"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("activity-randomNumber")

const (
	ovRand = "rand"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

func init() {
	logger.SetLogLevel(logger.InfoLevel)
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

	// do eval
	rand := getRandomNumber()
	context.SetOutput(ovRand, rand)

	return true, nil
}

// getSensorTemperature sends random
func getRandomNumber() float64 {
	rand.Seed(time.Now().Unix())
	return ((rand.Float64() * 10) + 30)
}
