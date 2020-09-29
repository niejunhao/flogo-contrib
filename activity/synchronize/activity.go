/*
 * Copyright © 2020. TIBCO Software Inc.
 * This file is subject to the license terms contained
 * in the license file that is distributed with this file.
 */
package synchronize

import (
	"sync"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
)

var activityLog = log.ChildLogger(log.RootLogger(), "tibco-activity-synchronize")

func init() {
	_ = activity.Register(&SyncActivity{}, New)
}

var activityMd = activity.ToMetadata(&Settings{}, &Input{})
var syncMap sync.Map

func New(ctx activity.InitContext) (activity.Activity, error) {
	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}
	return &SyncActivity{settings: s}, nil
}

type SyncActivity struct {
	settings *Settings
}

func (a *SyncActivity) Metadata() *activity.Metadata {
	return activityMd
}
func (a *SyncActivity) Eval(context activity.Context) (done bool, err error) {
	activityLog.Debug("Executing Synchronize activity")
	input := &Input{}
	_ = context.GetInputObject(input)
	//Read Inputs
	if len(input.Key) <= 0 {
		return false, activity.NewError("Key is not configured", "LOCK-4001", nil)
	}

	switch a.settings.Operation {
	case "Lock":
		lock, found := syncMap.Load(input.Key)
		if !found {
			lock = &sync.Mutex{}
			syncMap.Store(input.Key, lock)
		}
		activityLog.Infof("Waiting for lock. Key:%s, Instance:%s, Flow:%s ", input.Key, context.ActivityHost().ID(), context.ActivityHost().Name())
		(lock.(*sync.Mutex)).Lock()
		activityLog.Infof("Lock acquired. Key:%s, Instance:%s, Flow:%s ", input.Key, context.ActivityHost().ID(), context.ActivityHost().Name())
	case "Unlock":
		lock, found := syncMap.Load(input.Key)
		if found {
			(lock.(*sync.Mutex)).Unlock()
			activityLog.Infof("Lock is released. Key:%s, Instance:%s, Flow:%s ", input.Key, context.ActivityHost().ID(), context.ActivityHost().Name())
		} else {
			activityLog.Warnf("No lock found for the key:%s", input.Key)
		}
	}
	activityLog.Debug("Synchronize activity successfully executed")
	return true, nil
}
