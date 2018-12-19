## HMap
[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/fwhezfwhez/HMap)

HMap is suited for a map which has two key.

#### start

`go get github.com/fwhezfwhez/HMap`

#### Example:
```go
package main

import (
	"fmt"
	"sync"

	//"github.com/fwhezfwhez/HMap"
	"HMap"
)
func main() {
	//init a not-concurrently safe hMap
	//hm := HMap.New()

	// init a concurrently safe hMap
	hm := HMap.Default(&sync.RWMutex{})

	//set a value with two keys
	hm.Set("app_list", "cdd", "100.0.9.0")
	hm.Set("app_list", "cdd2", "100.0.9.1")
	hm.Set("app_list", "cdd3", "100.0.9.2")
	hm.Set("app_list2", "cdd4", "100.0.9.2")
	hm.Set("app_list3", "cdd5", "100.0.9.2")
	hm.Print()

	//get the value with two keys
	v,er:=hm.Get("app_list","cdd")
	if er!=nil{
		panic(er.Error())
	}
	fmt.Println(v)

	//get values array by main key
	vArr,er:=hm.GetByMainKey("app_list")
	if er!=nil{
		panic(er.Error())
	}
	fmt.Println(vArr)

	//delete values whose main key is "key"
	hm.DeleteByMainKey("app_list")
	hm.Print()
	//delete value whose main key is "key" and sub key is "sub_key"
	hm.Delete("app_list2","cdd4")
	hm.Print()
	//delete all datas
	hm.DeleteAll()
	hm.Print()
	//the same
 	hm.Clear()
	hm.Print()
}


```