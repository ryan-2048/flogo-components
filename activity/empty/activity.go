package empty

import (
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/log"
)

func init() {
	_ = activity.Register(&Activity{}, New)
}

var activityMetadata = activity.ToMetadata()

func New(ctx activity.InitContext) (activity.Activity, error) {
	act := &Activity{}
	return act, nil
}

type Activity struct{}

func (a *Activity) Metadata() *activity.Metadata {
	return activityMetadata
}

func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {
	log.RootLogger().Debug("Executing the null activity...")
	return true, nil
}
