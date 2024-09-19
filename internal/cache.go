package internal

import (
	"sync"
	"time"
)

const cacheLifeSpan = 30

type CacheInterface interface {
	RetrieveWeatherData(zipcode string) (WeatherData, bool)
	AddWeatherData(zipcode string, data WeatherData)
}

// MapCache implements the cache interface as a hashmap
type MapCache struct {
	cache map[string]WeatherData
	mux   sync.Mutex
}

// NewMapCache returns a fresh NewMapCache with everything initialized
func NewMapCache() *MapCache {
	c := &MapCache{
		mux:   sync.Mutex{},
		cache: map[string]WeatherData{},
	}
	return c
}

// RetrieveWeatherData fetchs the data from cache.
// The bool is false if the data does not exist in the cache or if the data is expired.
func (c *MapCache) RetrieveWeatherData(zipcode string) (WeatherData, bool) {
	c.mux.Lock()
	val, ok := c.cache[zipcode]
	c.mux.Unlock()
	if time.Now().After(val.RequestTime.Add(time.Minute * cacheLifeSpan)) {
		return WeatherData{}, false
	}
	return val, ok
}

// AddWeatherData stores the data in the hashmap with the zipcode being the key
func (c *MapCache) AddWeatherData(zipcode string, data WeatherData) {
	data.RequestTime = time.Now()
	data.IsCachedResponse = true
	c.mux.Lock()
	c.cache[zipcode] = data
	c.mux.Unlock()
}
