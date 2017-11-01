package pahoMQTT

import (
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput(ivBroker, "tcp://iot.eclipse.org:1883")
	tc.SetInput(ivTopic, "flogo/sensor")
	tc.SetInput(ivClientID, "sensor1")
	tc.SetInput(ivPayload, "33")
	tc.SetInput(ivQos, 0)
	tc.SetInput(ivCleansession, false)
	tc.SetInput(ivAction, "pub")

	act.Eval(tc)

	//check result attr
}

/* // Test input dump data
"inputs":[
    {
      "name": "broker",
      "type": "string",
      "required": true,
      "value":"tcp://iot.eclipse.org:1883"
    },
    {
      "name": "topic",
      "type": "string",
      "required":true,
      "value":"flogo/sensor"
    },
    {
      "name": "clientId",
      "type": "string",
      "value":"sensor1"
    },
    {
      "name": "cleansession",
      "type": "boolean",
      "value":false
    },
    {
      "name": "qos",
      "type": "integer",
      "value":0
    },
    {
      "name": "payload",
      "type": "string",
      "value":"temp"
    },
    {
      "name": "action",
      "type": "string",
      "value":"pub"
    }
  ],
  "outputs": [
    {
      "name": "output",
      "type": "any"
    }
  ]
*/
