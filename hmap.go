package HMap

import (
	"errors"
	"fmt"
)

type HMap struct {
	Content map[string]map[string]interface{}
}

//New a HMap
func New() *HMap {
	tmp := make(map[string]map[string]interface{}, 0)
	return &HMap{Content: tmp}
}

//Set by main key and sub key
func (hm *HMap) Set(mainKey string, subKey string, value interface{}) {
	if _, ok := hm.Content[mainKey]; !ok {
		tmp := make(map[string]interface{}, 0)
		tmp[subKey] = value
		hm.Content[mainKey] = tmp
	} else {
		hm.Content[mainKey][subKey] = value
	}
}

//get a value by main key and sub key
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

//get values array by main key
func (hm *HMap) GetByMainKey(mainKey string) ([]interface{}, error) {
	if _, ok := hm.Content[mainKey]; !ok {
		return nil, errors.New("no such mainKey key " + mainKey)
	}
	rs := make([]interface{},0)
	tmp := hm.Content[mainKey]
	for _,v:=range tmp {
		rs = append(rs,v)
	}
	return rs,nil
}

//delete a value indexed by main key and sub key
func (hm *HMap) Delete(mainKey string, subKey string){
	if _, ok := hm.Content[mainKey]; !ok {
		return
	}else{
		if _, ok2 := hm.Content[mainKey][subKey]; !ok2 {
			return
		}else{
			delete(hm.Content[mainKey], subKey)
		}
	}
}

//delete all datas by mainKey
func (hm *HMap) DeleteByMainKey(mainKey string) {
	if _, ok := hm.Content[mainKey]; !ok {
		return
	}else{
		hm.Content[mainKey] = make(map[string]interface{},0)
	}
}

//delete all datas in hm
func (hm *HMap) DeleteAll(){
	hm.Content = make(map[string]map[string]interface{},0)
}

//clear a hm
func (hm *HMap) Clear(){
	hm.DeleteAll()
}


//print this hm
func (hm *HMap) Print(){
	fmt.Println("MAIN_KEY|SUB_KEY|VALUE")
	for k1,v1:=range hm.Content{
		for k2,v2:=range v1{
			fmt.Println(fmt.Sprintf(" %s | %s | %v ",k1,k2,v2))
		}
	}
}
