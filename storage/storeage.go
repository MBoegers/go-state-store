package storage

import (
	"errors"
	"github.com/patrickmn/go-cache"
	"time"
)

// Storage simply wraps around a the go-cache library of Patrick Mylund Nielsen, see https://github.com/patrickmn
var internalCache = cache.New(cache.DefaultExpiration, cache.DefaultExpiration)

//region<Interaction>
// Set a a values, note that value nil is not allowed
func Set(key string, value interface{}) error {
	if value == nil {
		return errors.New("Value should not be null. For deletion use Remove")
	}
	internalCache.Set(key, value, cache.NoExpiration)
	return nil
}

// removes a value from the storage, an error indicates the absend of the value
func Remove(key string) error {
	return internalCache.Replace(key, nil, time.Microsecond*1) // use instant invalidation for removal
}

// read a value form the store, could be nil
func Get(key string) interface{} {
	var data, _ = internalCache.Get(key)
	return data
}

// check wether the cache contains a value
func Contains(key string) bool {
	var _, contains = internalCache.Get(key)
	return contains
}

//endregion
