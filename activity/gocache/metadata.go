package gocache

import (
	"github.com/project-flogo/core/data/coerce"
)

type Input struct {
	ivKey        string `md:"key"`
	ivValue      string `md:"value"`
	ivAction     string `md:"action"`
	ivExpiryTime int    `md:"expiryTime"`
	ivPurgeTime  int    `md:"purgeTime"`
	ivLoadset    string `md:"loadset"`
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"key":        i.ivKey,
		"value":      i.ivValue,
		"action":     i.ivAction,
		"expiryTime": i.ivExpiryTime,
		"purgeTime":  i.ivPurgeTime,
		"loadset":    i.ivLoadset,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.ivKey, err = coerce.ToString(values["key"])
	if err != nil {
		return err
	}
	i.ivValue, err = coerce.ToString(values["value"])
	if err != nil {
		return err
	}
	i.ivAction, err = coerce.ToString(values["action"])
	if err != nil {
		return err
	}
	i.ivExpiryTime, err = coerce.ToInt(values["expiryTime"])
	if err != nil {
		return err
	}
	i.ivPurgeTime, err = coerce.ToInt(values["purgeTime"])
	if err != nil {
		return err
	}
	i.ivLoadset, err = coerce.ToString(values["loadset"])
	if err != nil {
		return err
	}

	return nil
}

// Output is the output from the javascript engine
type Output struct {
	Error        bool                   `md:"error"`
	ErrorMessage string                 `md:"errorMessage"`
	Result       map[string]interface{} `md:"result"`
}

// FromMap converts the values from a map into the struct Output
func (o *Output) FromMap(values map[string]interface{}) error {
	errorValue, err := coerce.ToBool(values["error"])
	if err != nil {
		return err
	}
	o.Error = errorValue
	errorMessage, err := coerce.ToString(values["errorMessage"])
	if err != nil {
		return err
	}
	o.ErrorMessage = errorMessage
	result, err := coerce.ToObject(values["result"])
	if err != nil {
		return err
	}
	o.Result = result
	return nil
}

// ToMap converts the struct Output into a map
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"error":        o.Error,
		"errorMessage": o.ErrorMessage,
		"result":       o.Result,
	}
}
