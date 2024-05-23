package sleep

import (
	"github.com/project-flogo/core/data/coerce"
)

type Input struct {
	IvInterval     int    `md:"interval"`
	IvIntervalType string `md:"intervalType"`
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"interval":     i.IvInterval,
		"intervalType": i.IvIntervalType,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.IvInterval, err = coerce.ToInt(values["interval"])
	if err != nil {
		return err
	}
	i.IvIntervalType, err = coerce.ToString(values["intervalType"])
	if err != nil {
		return err
	}

	return nil
}
