package HMap

import (
	"errors"
	"fmt"
	"sync"
)

type HMap struct {
	Content map[string]map[string]interface{}
	Context map[string]interface{}
	M       *sync.RWMutex
}

// New a HMap without lock
func New() *HMap {
	tmp := make(map[string]map[string]interface{}, 0)
	return &HMap{
		Content: tmp,
		Context: make(map[string]interface{}, 0),
		M:       nil,
	}
}

// New a HMap with lock specific, if 'mutex' is nil,then it works as 'New()', which is not concurrently safe
func Default (mutex *sync.RWMutex) *HMap {
	tmp := make(map[string]map[string]interface{}, 0)
	return &HMap{
		Content: tmp,
		Context: make(map[string]interface{}, 0),
		M:       mutex,
	}
}

// Set  key-value for context
func (hm *HMap) SetContext(key string, value interface{}) {
	if hm.M != nil {
		hm.M.Lock()
		defer hm.M.Unlock()
	}
	hm.Context[key] = value
}

// Get value by key of context
func (hm *HMap) GetContext(key string) (interface{}, bool) {
	if hm.M != nil {
		hm.M.RLock()
		defer hm.M.RUnlock()
	}
	v, ok := hm.Context[key]
	return v, ok
}

// Set by main key and sub key
func (hm *HMap) Set(mainKey string, subKey string, value interface{}) {
	if hm.M != nil {
		hm.M.Lock()
		defer hm.M.Unlock()
	}

	if _, ok := hm.Content[mainKey]; !ok {
		tmp := make(map[string]interface{}, 0)
		tmp[subKey] = value
		hm.Content[mainKey] = tmp
	} else {
		hm.Content[mainKey][subKey] = value
	}
}

// Get a value by main key and sub key
func (hm *HMap) Get(mainKey string, subKey string) (interface{}, error) {
	if hm.M != nil {
		hm.M.RLock()
		defer hm.M.RUnlock()
	}

	if _, ok := hm.Content[mainKey]; ok {
		if _, ok2 := hm.Content[mainKey][subKey]; ok2 {
			return hm.Content[mainKey][subKey], nil
		} else {
			return nil, errors.New("no such subKey key " + subKey)
		}
	} else {
		return nil, errors.New("no such mainKey key " + mainKey)
	}
}

// Get values array by main key
func (hm *HMap) GetByMainKey(mainKey string) ([]interface{}, error) {
	if hm.M != nil {
		hm.M.RLock()
		defer hm.M.RUnlock()
	}

	if _, ok := hm.Content[mainKey]; !ok {
		return nil, errors.New("no such mainKey key " + mainKey)
	}
	rs := make([]interface{}, 0, 5)
	tmp := hm.Content[mainKey]
	for _, v := range tmp {
		rs = append(rs, v)
	}
	return rs, nil
}

// Delete a value indexed by main key and sub key
func (hm *HMap) Delete(mainKey string, subKey string) {
	if hm.M != nil {
		hm.M.Lock()
		defer hm.M.Unlock()
	}

	if _, ok := hm.Content[mainKey]; !ok {
		return
	} else {
		if _, ok2 := hm.Content[mainKey][subKey]; !ok2 {
			return
		} else {
			delete(hm.Content[mainKey], subKey)
		}
	}
}

// Delete all datas by mainKey
func (hm *HMap) DeleteByMainKey(mainKey string) {
	if hm.M != nil {
		hm.M.Lock()
		defer hm.M.Unlock()
	}

	if _, ok := hm.Content[mainKey]; !ok {
		return
	} else {
		hm.Content[mainKey] = make(map[string]interface{}, 0)
	}
}

// Delete all datas in hm
func (hm *HMap) DeleteAll() {
	if hm.M != nil {
		hm.M.Lock()
		defer hm.M.Unlock()
	}
	hm.Content = make(map[string]map[string]interface{}, 0)
}

// Clear a hm
func (hm *HMap) Clear() {
	hm.DeleteAll()
}

// Print this hm
func (hm *HMap) Print() {
	if hm.M != nil {
		hm.M.RLock()
		defer hm.M.RUnlock()
	}

	fmt.Println("MAIN_KEY|SUB_KEY|VALUE")
	for k1, v1 := range hm.Content {
		for k2, v2 := range v1 {
			fmt.Println(fmt.Sprintf(" %s | %s | %v ", k1, k2, v2))
		}
	}
}
