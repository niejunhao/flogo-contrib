/*
 * Copyright © 2020. TIBCO Software Inc.
 * This file is subject to the license terms contained
 * in the license file that is distributed with this file.
 */
package synchronize

import (
	"os"
	"testing"

	"github.com/project-flogo/core/data/mapper"
	"github.com/project-flogo/core/data/resolve"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

func TestSimpleLockUnlock(t *testing.T) {
	_ = os.Setenv("FLOGO_LOG_LEVEL", "DEBUG")

	settings := &Settings{Operation: "Lock"}

	mf := mapper.NewFactory(resolve.GetBasicResolver())
	iCtx := test.NewActivityInitContext(settings, mf)
	lockAct, err := New(iCtx)
	assert.Nil(t, err)

	tc := test.NewActivityContext(lockAct.Metadata())

	//eval
	_, _ = lockAct.Eval(tc)

	settings = &Settings{Operation: "Unlock"}

	iCtx = test.NewActivityInitContext(settings, mf)
	unlockAct, err1 := New(iCtx)
	assert.Nil(t, err1)
	tc = test.NewActivityContext(unlockAct.Metadata())
	assert.Nil(t, err1)
	_, _ = unlockAct.Eval(tc)
}
