package sleep

import (
	"time"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/log"
)

func init() {
	_ = activity.Register(&Activity{}, New)
}

var activityMetadata = activity.ToMetadata(&Input{})

func New(ctx activity.InitContext) (activity.Activity, error) {
	act := &Activity{}
	return act, nil
}

type Activity struct{}

func (a *Activity) Metadata() *activity.Metadata {
	return activityMetadata
}

func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return false, err
	}

	switch input.IvIntervalType {
	case "Millisecond":
		time.Sleep(time.Duration(input.IvInterval) * time.Millisecond)
	case "Second":
		time.Sleep(time.Duration(input.IvInterval) * time.Second)
	case "Minute":
		time.Sleep(time.Duration(input.IvInterval) * time.Minute)
	default:
		return false, activity.NewError("Unsupported Interval Type. Supported Types- [Millisecond, Second, Minute]", "", nil)
	}

	log.RootLogger().Info("Sleep activity completed")
	return true, nil

}
