package sleep

import (
	"time"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// activityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-flogo-sleep")

const (
	ivInterval     = "interval"
	ivIntervalType = "intervalType"
)

// SleepActivity is an Activity that can stop flow execution for given time duration.
// inputs : {interval, intervalType}
// outputs: none
type SleepActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new AppActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &SleepActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *SleepActivity) Metadata() *activity.Metadata {
	return a.metadata
}

func (a *SleepActivity) Eval(context activity.Context) (done bool, err error) {

	//mv := context.GetInput(ivMessage)
	activityLog.Info("Executing Sleep activity")
	interval, _ := context.GetInput(ivInterval).(int)

	intervalType, _ := context.GetInput(ivIntervalType).(string)

	switch intervalType {
	case "Millisecond":
		time.Sleep(time.Duration(interval) * time.Millisecond)
	case "Second":
		time.Sleep(time.Duration(interval) * time.Second)
	case "Minute":
		time.Sleep(time.Duration(interval) * time.Minute)
	default:
		return false, activity.NewError("Unsupported Interval Type. Supported Types- [Millisecond, Second, Minute]", "", nil)
	}

	activityLog.Info("Sleep activity completed")
	return true, nil
}
