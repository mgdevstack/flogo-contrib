package sensorProducerMQTT

import (
	"math/rand"
	"time"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("activity-sensorProducerMQTT")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// Setup log level
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

	// do eval
	brokerURL := context.GetInput("mqttBrokerURL").(string)
	sensorID := context.GetInput("sensorID").(string)

	// logging data to debug
	log.Debugf("Broker Address: [%s]", brokerURL)
	log.Debugf("Sensor ID: [%s]", sensorID)

	//connect with MQTT Broker

	ticker := time.NewTicker(time.Second)

	go func() {
		for t := range ticker.C {
			log.Infof("Temperature: [%0.2f] at time [%s]", getSensorTemperature(), t.Format(time.RFC3339))
		}
	}()

	time.Sleep(time.Second * 2)
	ticker.Stop()
	log.Info("Data Sent")

	//check existing topics against sensor ID (crete topic if missing)

	//store temperature data

	// Set Activities Output
	context.SetOutput("isTopicExists", "true")
	context.SetOutput("result", "sent")
	return true, nil
}

// getSensorTemperature sends random
func getSensorTemperature() float64 {
	rand.Seed(time.Now().Unix())
	return ((rand.Float64() * 10) + 30)
}
