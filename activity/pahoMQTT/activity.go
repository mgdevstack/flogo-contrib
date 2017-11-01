package pahoMQTT

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var log = logger.GetLogger("activity-pahoMQTT")

const (
	ivBroker       = "broker"
	ivTopic        = "topic"
	ivClientID     = "clientId"
	ivCleansession = "cleansession"
	ivQos          = "qos"
	ivPayload      = "payload"
	ivAction       = "action"
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

	// Get client's input config
	mqttBroker := context.GetInput(ivBroker).(string)
	mqttTopic := context.GetInput(ivTopic).(string)
	mqttClientID := context.GetInput(ivClientID).(string)
	mqttCleansession := context.GetInput(ivCleansession).(bool)
	mqttQos := context.GetInput(ivQos).(int)
	mqttPayload := context.GetInput(ivPayload).(string)
	mqttAction := context.GetInput(ivAction).(string)

	pubTopic := mqttTopic + "/" + mqttClientID
	// mqtt_action := context.GetInput("action").(string)

	// log.Infof("MQTT Action: [%s]", mqtt_action)

	opts := MQTT.NewClientOptions()
	opts.AddBroker(mqttBroker)
	opts.SetClientID(mqttClientID)
	opts.SetCleanSession(mqttCleansession)

	if mqttAction == "pub" {
		client := MQTT.NewClient(opts)
		if ctoken := client.Connect(); ctoken.Wait() && ctoken.Error() != nil {
			panic(ctoken.Error())
		}
		ptoken := client.Publish(pubTopic, byte(mqttQos), false, mqttPayload)
		log.Info("Data published with token: ", ptoken)
		ptoken.Wait()
		client.Disconnect(200)
		log.Info("MQTT client disconnected")
	}
	return true, nil
}
