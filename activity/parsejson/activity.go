package parsejson

import (
	"encoding/json"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// activityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("tibco-parsejson-activity")

const (
	ivJSONData   = "jsonData"
	ovJSONObject = "jsonObject"
)

// ParseJSONActivity is an Activity that can stop flow execution for given time duration.
// inputs : {interval, intervalType}
// outputs: none
type ParseJSONActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new AppActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &ParseJSONActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *ParseJSONActivity) Metadata() *activity.Metadata {
	return a.metadata
}

func (a *ParseJSONActivity) Eval(context activity.Context) (done bool, err error) {

	activityLog.Info("Executing ParseJSON activity")
	jsonData, err := data.CoerceToString(context.GetInput(ivJSONData))

	if err != nil {
		return false, activity.NewError("Invalid JSON string data", "", nil)
	}

	var raw map[string]interface{}

	err = json.Unmarshal([]byte(jsonData), &raw)
	if err != nil {
		activityLog.Error(err)
		return false, activity.NewError("Failed to parse JSON data", "", nil)
	}

	output := &data.ComplexObject{Value: raw}
	context.SetOutput(ovJSONObject, output)

	activityLog.Info("ParseJSON activity completed")
	return true, nil
}
