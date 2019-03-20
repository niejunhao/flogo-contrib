package xml2json

import (
	"encoding/json"
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	xj "github.com/basgys/goxml2json"
)

// activityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("tibco-xml2json-activity")

const (
	ivXMLData    = "xmlData"
	ovJSONObject = "jsonObject"
)

// XML2JSONActivity is an Activity that can stop flow execution for given time duration.
// inputs : {interval, intervalType}
// outputs: none
type XML2JSONActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new AppActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &XML2JSONActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *XML2JSONActivity) Metadata() *activity.Metadata {
	return a.metadata
}

func (a *XML2JSONActivity) Eval(context activity.Context) (done bool, err error) {

	activityLog.Info("Executing XML2JSON activity")
	xmlData, err := data.CoerceToString(context.GetInput(ivXMLData))

	if err != nil {
		return false, activity.NewError("Invalid XML string", "", nil)
	}

	var raw map[string]interface{}


	xml := strings.NewReader(xmlData)
	jsonData, err := xj.Convert(xml, xj.WithTypeConverter(xj.Float), xj.WithTypeConverter(xj.Int), xj.WithTypeConverter(xj.Bool))
	if err != nil {
		activityLog.Error(err)
		return false, activity.NewError("Failed to convert XML data", "", nil)
	}

	err = json.Unmarshal(jsonData.Bytes(), &raw)
	if err != nil {
		activityLog.Error(err)
		return false, activity.NewError("Failed to parse JSON data", "", nil)
	}

	output := &data.ComplexObject{Value: raw}
	context.SetOutput(ovJSONObject, output)

	activityLog.Info("XML2JSON activity completed")
	return true, nil
}
