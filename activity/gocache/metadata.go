package gocache

import (
	"github.com/project-flogo/core/data/coerce"
)

type Input struct {
	IvKey        string `md:"key"`
	IvValue      string `md:"value"`
	IvAction     string `md:"action"`
	IvExpiryTime int    `md:"expiryTime"`
	IvPurgeTime  int    `md:"purgeTime"`
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"key":        i.IvKey,
		"value":      i.IvValue,
		"action":     i.IvAction,
		"expiryTime": i.IvExpiryTime,
		"purgeTime":  i.IvPurgeTime,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.IvKey, err = coerce.ToString(values["key"])
	if err != nil {
		return err
	}
	i.IvValue, err = coerce.ToString(values["value"])
	if err != nil {
		return err
	}
	i.IvAction, err = coerce.ToString(values["action"])
	if err != nil {
		return err
	}
	i.IvExpiryTime, err = coerce.ToInt(values["expiryTime"])
	if err != nil {
		return err
	}
	i.IvPurgeTime, err = coerce.ToInt(values["purgeTime"])
	if err != nil {
		return err
	}

	return nil
}

type Output struct {
	Error        bool        `md:"error"`
	ErrorMessage string      `md:"errorMessage"`
	Result       interface{} `md:"result"`
}

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

	o.Result = values["result"]
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"error":        o.Error,
		"errorMessage": o.ErrorMessage,
		"result":       o.Result,
	}
}
