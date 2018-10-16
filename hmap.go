package HMap

import (
	"errors"
	"fmt"
	"sync"
)

type HMap struct {
	Content map[string]map[string]interface{}
	Context map[string]interface{}
	M       *sync.Mutex
}

// New a HMap
func New() *HMap {
	tmp := make(map[string]map[string]interface{}, 0)
	return &HMap{
		Content: tmp,
		M:       &sync.Mutex{},
		Context: make(map[string]interface{}, 0),
	}
}

// Lock of the map
func (hm *HMap) GetLock() *sync.Mutex {
	return hm.M
}
func (hm *HMap) Lock() *sync.Mutex {
	hm.M.Lock()
	return hm.M
}
func (hm *HMap) UnLock() *sync.Mutex {
	hm.M.Unlock()
	return hm.M
}

// Set  key-value for context
func (hm *HMap) SetContext(key string, value interface{}) {
	hm.Context[key] = value
}

// Get value by key of context
func (hm *HMap) GetContext(key string) (interface{},bool){
	v,ok:= hm.Context[key]
	return v,ok
}


// Set by main key and sub key
func (hm *HMap) Set(mainKey string, subKey string, value interface{}) {
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
	if _, ok := hm.Content[mainKey]; !ok {
		return nil, errors.New("no such mainKey key " + mainKey)
	}
	rs := make([]interface{}, 0)
	tmp := hm.Content[mainKey]
	for _, v := range tmp {
		rs = append(rs, v)
	}
	return rs, nil
}

// Delete a value indexed by main key and sub key
func (hm *HMap) Delete(mainKey string, subKey string) {
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
	if _, ok := hm.Content[mainKey]; !ok {
		return
	} else {
		hm.Content[mainKey] = make(map[string]interface{}, 0)
	}
}

// Delete all datas in hm
func (hm *HMap) DeleteAll() {
	hm.Content = make(map[string]map[string]interface{}, 0)
}

// Clear a hm
func (hm *HMap) Clear() {
	hm.DeleteAll()
}

// Print this hm
func (hm *HMap) Print() {
	fmt.Println("MAIN_KEY|SUB_KEY|VALUE")
	for k1, v1 := range hm.Content {
		for k2, v2 := range v1 {
			fmt.Println(fmt.Sprintf(" %s | %s | %v ", k1, k2, v2))
		}
	}
}
