package gocache

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/patrickmn/go-cache"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/log"
)

const (
	CacheName = "Gocache"
)

func init() {
	_ = activity.Register(&Activity{}, New)
}

var activityMetadata = activity.ToMetadata(&Input{}, &Output{})

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

	output := Output{}
	var result interface{}

	output.Error = false
	output.ErrorMessage = ""
	output.Result = result

	switch input.IvAction {

	case "OBJ_SET":
		var c *cache.Cache
		val, ok := data.GetGlobalScope().GetAttr(CacheName)
		if ok {
			c = val.Value().(*cache.Cache)
		} else {
			log.RootLogger().Infof("cache doesn't exist yet, it will be created first with default settings of expiryTime [%d] minutes and purgeTime [%d] minutes", input.IvExpiryTime, input.IvPurgeTime)
			c = initializeCache(input.IvExpiryTime, input.IvPurgeTime)
		}
		if len(input.IvValue) > 0 {
			var cacheData map[string]interface{}
			if err := json.Unmarshal([]byte(input.IvValue), &cacheData); err != nil {
				log.RootLogger().Errorf("couldn't load cache with data: %s", err.Error())
				output.Error = true
				output.ErrorMessage = err.Error()
				output.Result = result
			} else {
				for key, value := range cacheData {
					set(c, key, value)
				}
			}
		}

	case "SET":
		key := input.IvKey
		value := input.IvValue
		var c *cache.Cache
		val, ok := data.GetGlobalScope().GetAttr(CacheName)
		if ok {
			c = val.Value().(*cache.Cache)
		} else {
			log.RootLogger().Infof("cache doesn't exist yet, it will be created first with default settings of expiryTime [%d] minutes and purgeTime [%d] minutes", input.IvExpiryTime, input.IvPurgeTime)
			c = initializeCache(input.IvExpiryTime, input.IvPurgeTime)
		}
		set(c, key, value)

	case "GET":
		key := input.IvKey

		var c *cache.Cache
		val, ok := data.GetGlobalScope().GetAttr(CacheName)
		if ok {
			c = val.Value().(*cache.Cache)
		} else {
			log.RootLogger().Error("cache doesn't exist")
			return false, fmt.Errorf("cache doesn't exist")
		}

		cacheVal, found := get(c, key)

		if found {
			output.Result = cacheVal
		} else {
			log.RootLogger().Infof("No cache entry was found for [%s]", key)
		}

	case "BATCH_GET":
		key := input.IvKey

		var c *cache.Cache
		val, ok := data.GetGlobalScope().GetAttr(CacheName)
		if ok {
			c = val.Value().(*cache.Cache)
		} else {
			log.RootLogger().Error("cache doesn't exist")
			return false, fmt.Errorf("cache doesn't exist")
		}

		result := make(map[string]interface{})

		var keys []string

		if err := json.Unmarshal([]byte(key), &keys); err != nil {
			log.RootLogger().Error("Error during unmarshaling: %v", err)
		}

		for _, item := range keys {
			cacheVal, found := get(c, item)
			if found {
				result[item] = cacheVal
			}
		}

		output.Result = result

	case "DELETE":
		key := input.IvKey
		var c *cache.Cache
		val, ok := data.GetGlobalScope().GetAttr(CacheName)
		if ok {
			c = val.Value().(*cache.Cache)
		} else {
			log.RootLogger().Error("cache doesn't exist")
			return false, fmt.Errorf("cache doesn't exist")
		}
		delete(c, key)

	default:
		log.RootLogger().Errorf("action [%s] does not exist in Gocache", input.IvAction)
	}

	err = ctx.SetOutputObject(&output)
	if err != nil {
		return false, err
	}

	return true, nil
}

func initializeCache(expiryTime int, purgeTime int) *cache.Cache {
	newCache := cache.New(time.Duration(expiryTime)*time.Minute, time.Duration(purgeTime)*time.Minute)
	data.GetGlobalScope().AddAttr(CacheName, data.TypeAny, newCache)
	log.RootLogger().Infof("Created cache with expiryTime [%d] minutes and purgeTime [%d] minutes", expiryTime, purgeTime)
	return newCache
}

func set(c *cache.Cache, key string, value interface{}) {
	c.Set(key, value, cache.DefaultExpiration)
}

func get(c *cache.Cache, key string) (interface{}, bool) {
	return c.Get(key)
}

func delete(c *cache.Cache, key string) {
	c.Delete(key)
}
