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
	// CacheName represents the name of the cache as it exists in the Flogo engine
	CacheName = "Gocache"
	// DefaultExpiryTime is the default expiry time if no cache was created
	DefaultExpiryTime = 5
	// DefaultPurgeTime is the default purge time if no cache was created
	DefaultPurgeTime = 10
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

// Metadata implements activity.Activity.Metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMetadata
}

// Eval implements activity.Activity.Eval
func (a *Activity) Eval(context activity.Context) (done bool, err error) {

	input := &Input{}
	err = context.GetInputObject(input)
	if err != nil {
		return false, err
	}

	output := Output{}
	result := make(map[string]interface{})
	if err != nil {
		output.Error = true
		output.ErrorMessage = err.Error()
		return false, err
	}

	switch input.ivAction {
	case "INITIALIZE":
		// Execute the initialize action
		var c *cache.Cache
		c = initializeCache(input.ivExpiryTime, input.ivPurgeTime)

		// Load the cache with data if needed
		if len(input.ivLoadset) > 0 {
			var cacheData map[string]interface{}
			if err := json.Unmarshal([]byte(input.ivLoadset), &cacheData); err != nil {
				log.RootLogger().Errorf("couldn't load cache with data: %s", err.Error())
				return false, fmt.Errorf("couldn't load cache with data: %s", err.Error())
			}
			for key, value := range cacheData {
				// Execute the set action
				set(c, key, value.(string))
			}
		}

		// Set the output context
		output.Result = result
		return true, nil
	case "SET":
		// Get the input values
		key := input.ivKey
		value := input.ivValue

		// Initialize if the cache doesn't exist
		var c *cache.Cache
		val, ok := data.GetGlobalScope().GetAttr(CacheName)
		if ok {
			c = val.Value().(*cache.Cache)
		} else {
			log.RootLogger().Infof("cache doesn't exist yet, it will be created first with default settings of expiryTime [%d] minutes and purgeTime [%d] minutes", DefaultExpiryTime, DefaultPurgeTime)
			c = initializeCache(DefaultExpiryTime, DefaultPurgeTime)
		}

		// Execute the set action
		set(c, key, value)

		// Set the output context
		output.Result = result
		return true, nil
	case "GET":
		// Get the input values
		key := input.ivKey

		// Fail if the cache doesn't exist
		var c *cache.Cache
		val, ok := data.GetGlobalScope().GetAttr(CacheName)
		if ok {
			c = val.Value().(*cache.Cache)
		} else {
			log.RootLogger().Error("cache doesn't exist")
			return false, fmt.Errorf("cache doesn't exist")
		}

		// Execute the Get action
		cacheVal, found := get(c, key)

		// Set the output context
		if found {
			var valueJSON json.RawMessage
			var vmObject map[string]interface{}
			valueJSON, err = json.Marshal(cacheVal)
			if err != nil {
				return false, err
			}
			err = json.Unmarshal(valueJSON, &vmObject)
			if err != nil {
				return false, err
			}
			output.Result = vmObject
			return true, nil
		}
		log.RootLogger().Infof("No cache entry was found for [%s]", key)
		output.Result = result
		return true, nil
	case "DELETE":
		// Get the input values
		key := input.ivKey

		// Fail if the cache doesn't exist
		var c *cache.Cache
		val, ok := data.GetGlobalScope().GetAttr(CacheName)
		if ok {
			c = val.Value().(*cache.Cache)
		} else {
			log.RootLogger().Error("cache doesn't exist")
			return false, fmt.Errorf("cache doesn't exist")
		}

		// Execute the delete action
		delete(c, key)

		output.Result = result
		return true, nil
	default:
		log.RootLogger().Errorf("action [%s] does not exist in Gocache", input.ivAction)
		return false, fmt.Errorf("action [%s] does not exist in Gocache", input.ivAction)
	}
}

// initializeCache initializes the cache with expiration time in minutes, and which purges expired items every set amount of minutes
func initializeCache(expiryTime int, purgeTime int) *cache.Cache {
	newCache := cache.New(time.Duration(expiryTime)*time.Minute, time.Duration(purgeTime)*time.Minute)
	data.GetGlobalScope().AddAttr(CacheName, data.TypeAny, newCache)
	log.RootLogger().Infof("Created cache with expiryTime [%d] minutes and purgeTime [%d] minutes", expiryTime, purgeTime)
	return newCache
}

// set adds a new entry to the cache with a default expiration time
func set(c *cache.Cache, key string, value string) {
	c.Set(key, value, cache.DefaultExpiration)
}

// get retrieves an entry from the cache
func get(c *cache.Cache, key string) (interface{}, bool) {
	return c.Get(key)
}

// delete removes an entry from the cache
func delete(c *cache.Cache, key string) {
	c.Delete(key)
}
